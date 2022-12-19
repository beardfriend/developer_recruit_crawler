package engine

import (
	"context"
	"fmt"
	"sync"

	"dev_recruitment_crawler/model"
	"dev_recruitment_crawler/provider"
	"dev_recruitment_crawler/provider/jumpit"
	"dev_recruitment_crawler/provider/programmers"
	"dev_recruitment_crawler/provider/saramin"
	"dev_recruitment_crawler/provider/wanted"

	"go.mongodb.org/mongo-driver/mongo"
)

type Engine struct {
	Providers map[string]provider.Provider
	db        *mongo.Client
}

func NewEngine(mongo *mongo.Client) *Engine {
	providers := make(map[string]provider.Provider, 0)
	providers["jumpit"] = jumpit.NewJumpit()
	providers["wanted"] = wanted.NewWanted()
	providers["programmers"] = programmers.NewProgrammers()
	providers["saramin"] = saramin.NewSaramin()
	return &Engine{
		Providers: providers,
		db:        mongo,
	}
}

func (e *Engine) CronRecruitment() {
	coll := e.db.Database("db").Collection("recruitments")
	m := e.GetRecruitment()
	newRecruit := make([]interface{}, 0)
	for _, v := range m {
		newRecruit = append(newRecruit, v)
	}
	fmt.Println(len(newRecruit))
	coll.InsertMany(context.TODO(), newRecruit)
}

func (e *Engine) GetRecruitment() []*model.Recruitment {
	var wg sync.WaitGroup
	result := make([]*model.Recruitment, 0)
	respCh := make(chan []*model.Recruitment)
	for _, v := range e.Providers {
		wg.Add(1)
		go func(provider provider.Provider) {
			defer wg.Done()
			result := provider.GetRecruitment(1, "backend")
			respCh <- result
		}(v)
	}

	go func() {
		wg.Wait()
		close(respCh)
	}()

	for resp := range respCh {
		result = append(result, resp...)
	}

	return result
}
