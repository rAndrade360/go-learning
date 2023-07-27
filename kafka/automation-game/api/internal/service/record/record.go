package record

import (
	"github.com/rAndrade360/go-learning/kafka/automation-game/api/internal/entity"
	"github.com/rAndrade360/go-learning/kafka/automation-game/api/internal/pkg/json"
	"github.com/rAndrade360/go-learning/kafka/automation-game/api/internal/pkg/kafka"
)

type service struct {
	kfk kafka.KafkaPkg
}

type RecordService interface {
	Create(r *entity.Record) error
}

func NewRecordService(kfk kafka.KafkaPkg) RecordService {
	return &service{
		kfk: kfk,
	}
}

func (s *service) Create(r *entity.Record) error {
	b, err := json.Marshal(*r)
	if err != nil {
		return err
	}

	return s.kfk.Produce(b)
}
