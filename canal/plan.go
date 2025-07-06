package canal

import (
	"panda-trip/internal/model"
)

type Plan struct {
	model.PlanCache
}

func NewPlan() *Plan {
	return &Plan{
		PlanCache: model.PlanCache{},
	}
}
