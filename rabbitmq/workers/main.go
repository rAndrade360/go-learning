package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rabbitmq/amqp091-go"
)

func init() {
	godotenv.Load()
}

func main() {
	rabbit_uql := os.Getenv("RABBIT_URL")
	rabbit_queue_name := os.Getenv("RABBIT_QUEUE_NAME")

	conn, err := amqp091.Dial(rabbit_uql)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(rabbit_queue_name, true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	sm := http.NewServeMux()

	sm.HandleFunc("/pub", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[PUB] request received at path /post with method ", r.Method)

		if r.Method != http.MethodPost {
			w.WriteHeader(405)
			w.Write([]byte(`{"message": "method not allowed"}`))
			return
		}

		b, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(`{"message": "bad request"}`))
			return
		}

		err = ch.PublishWithContext(context.TODO(), "", q.Name, false, false, amqp091.Publishing{
			DeliveryMode: amqp091.Persistent,
			ContentType:  "application/json",
			Body:         b,
		})

		log.Println("[PUB] body set to queue: ", string(b))

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(`{"message": "internal server error"}`))
			return
		}

		w.WriteHeader(200)
		w.Write([]byte(`{"message": "sent to queue"}`))
	})

	err = ch.Qos(1, 0, false)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(rabbit_queue_name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for d := range msgs {
			log.Printf("[SUB] Received msg: %s", string(d.Body))
			d.Ack(true)
		}
	}()

	log.Println("Server running at port 8080")
	log.Fatal(http.ListenAndServe(":8080", sm))
}
