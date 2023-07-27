package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	recordController "github.com/rAndrade360/go-learning/kafka/automation-game/api/handler/record"
	kafkapkg "github.com/rAndrade360/go-learning/kafka/automation-game/api/internal/pkg/kafka"
	recordService "github.com/rAndrade360/go-learning/kafka/automation-game/api/internal/service/record"
	"github.com/segmentio/kafka-go"
)

func init() {
	godotenv.Load()
}

func main() {
	r := mux.NewRouter()

	topic := os.Getenv("TOPIC")
	partition, err := strconv.Atoi(os.Getenv("PARTITION"))
	if err != nil {
		log.Fatal("failed to load env:", err)
	}

	conn, err := kafka.DialLeader(context.Background(), "tcp", fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT")), topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	defer conn.Close()

	kfkconn := kafkapkg.NewKafkaPkg(conn)
	recordSvc := recordService.NewRecordService(kfkconn)
	recordCtrl := recordController.NewRecordController(recordSvc)

	r.Handle("/record", recordCtrl.Create()).Methods("POST")

	http.ListenAndServe(":8080", r)
}
