package queue

import (
	"fmt"
	"log"
	"tgo/api/internal/config"

	"github.com/streadway/amqp"
)

var amqpURL = config.GetEnv("RABBITMQ_URL", "amqp://admin:admin@localhost:5672/")

type RabbitMQAdapter struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQAdapter() (*RabbitMQAdapter, error) {

	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		log.Printf("Erro ao conectar ao RabbitMQ: %v", err)
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Printf("Erro ao abrir canal no RabbitMQ: %v", err)
		return nil, err
	}
	channel.QueueDeclare(
		"test-queue", // Nome da fila
		true,         // Durável (se verdadeiro, a fila sobrevive a reinicializações)
		false,        // Auto-deleta
		false,        // Exclusiva (somente uma conexão pode usar)
		false,        // No-wait
		nil,          // Argumentos adicionais
	)

	return &RabbitMQAdapter{
		conn:    conn,
		channel: channel,
	}, nil
}

func (r *RabbitMQAdapter) Publish(queueName string, message []byte) error {

	// Declarar a fila (garantir que a fila exista antes de publicar)
	_, err := r.channel.QueueDeclare(
		queueName, // Nome da fila
		true,      // Durável
		false,     // Auto-deleta
		false,     // Exclusiva
		false,     // No-wait
		nil,       // Argumentos adicionais
	)
	if err != nil {
		log.Printf("Erro ao declarar a fila: %v", err)
		return err
	}

	err = r.channel.Publish(
		"",        // Troca (exchange) vazia usa o exchange padrão
		queueName, // Nome da fila como chave de roteamento (routing key)
		false,     // Obrigatório
		false,     // Imediato
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	if err != nil {
		log.Printf("Erro ao publicar mensagem: %v", err)
		return err
	}

	log.Printf("Mensagem publicada na fila %s", queueName)
	return nil
}

func (r *RabbitMQAdapter) Consume(queueName string, handler func(message []byte)) error {
	msgs, err := r.channel.Consume(
		queueName, // Nome da fila
		"",        // Nome do consumidor (consumer)
		false,     // Auto-ack
		false,     // Exclusivo
		false,     // No-local
		false,     // No-wait
		nil,       // Argumentos adicionais
	)
	if err != nil {
		log.Printf("Erro ao consumir mensagens da fila: %v", err)
		return err
	}

	go func() {
		for msg := range msgs {
			log.Printf("Mensagem recebida: %s", msg.Body)
			handler(msg.Body)
		}
	}()

	return nil
}

func (r *RabbitMQAdapter) reconnect() error {
	var err error
	r.conn, err = amqp.Dial(amqpURL)
	if err != nil {
		return fmt.Errorf("falha ao reconectar ao RabbitMQ: %v", err)
	}

	r.channel, err = r.conn.Channel()
	if err != nil {
		return fmt.Errorf("falha ao abrir canal no RabbitMQ após reconexão: %v", err)
	}

	log.Println("Reconectado ao RabbitMQ com sucesso")
	return nil
}

func (r *RabbitMQAdapter) isConnectionOpen() bool {
	if r.conn.IsClosed() || r.channel == nil {
		return false
	}
	return true
}

func (r *RabbitMQAdapter) ConsumeWithReconnect(queueName string, handler func(message []byte)) error {
	if !r.isConnectionOpen() {
		log.Println("Conexão perdida, tentando reconectar...")
		err := r.reconnect()
		if err != nil {
			return err
		}
	}

	msgs, err := r.channel.Consume(
		queueName, // Nome da fila
		"",        // Nome do consumidor (consumer)
		true,      // Auto-ack
		false,     // Exclusivo
		false,     // No-local
		false,     // No-wait
		nil,       // Argumentos adicionais
	)
	if err != nil {
		log.Printf("Erro ao consumir mensagens da fila: %v", err)
		return err
	}

	// Processar mensagens em uma goroutine
	go func() {
		for msg := range msgs {
			log.Printf("Mensagem recebida: %s", msg.Body)
			handler(msg.Body) // Passa a mensagem recebida para o handler
		}
	}()

	return nil
}

func (r *RabbitMQAdapter) Close() {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}
