package record

import (
	"io"
	"net/http"

	"github.com/rAndrade360/go-learning/kafka/automation-game/api/internal/entity"
	"github.com/rAndrade360/go-learning/kafka/automation-game/api/internal/pkg/json"
	recordService "github.com/rAndrade360/go-learning/kafka/automation-game/api/internal/service/record"
)

type controller struct {
	service recordService.RecordService
}

type RecordController interface {
	Create() http.Handler
}

func NewRecordController(service recordService.RecordService) RecordController {
	return &controller{
		service: service,
	}
}

func (c *controller) Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, `{"message":"Bad Request"}`, http.StatusBadRequest)
			return
		}

		var record entity.Record

		err = json.Unmarshal(b, &record)
		if err != nil {
			http.Error(w, `{"message":"Bad Request"}`, http.StatusBadRequest)
			return
		}

		err = c.service.Create(&record)
		if err != nil {
			http.Error(w, `{"message":"Internal Server Error"}`, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})
}
