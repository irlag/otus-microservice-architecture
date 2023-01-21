package processors

import (
	"go.uber.org/zap"
)

type Processors struct {
	HealthcheckProcessor HealthcheckProcessor
}

func NewProcessor(
	log *zap.Logger,
) *Processors {
	return &Processors{
		HealthcheckProcessor: NewHealtcheckProcessor(),
	}
}
