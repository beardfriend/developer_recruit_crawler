package engine

import (
	"context"
	"fmt"
	"sync"
	"time"

	"dev_recruitment_crawler/model"
	"dev_recruitment_crawler/provider"
	"dev_recruitment_crawler/provider/jumpit"
	"dev_recruitment_crawler/provider/programmers"
	"dev_recruitment_crawler/provider/saramin"
	"dev_recruitment_crawler/provider/wanted"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Engine struct {
	Providers map[string]provider.Provider
	db        *mongo.Client
}

func NewEngine(db *mongo.Client) *Engine {
	providers := make(map[string]provider.Provider, 0)
	providers["jumpit"] = jumpit.NewJumpit()
	providers["wanted"] = wanted.NewWanted()
	providers["programmers"] = programmers.NewProgrammers()
	providers["saramin"] = saramin.NewSaramin()
	return &Engine{
		db:        db,
		Providers: providers,
	}
}

func (e *Engine) CronRecruitment() {
	positions := []string{"backend", "frontend"}
	careers := []int{0, 1, 2, 3, 4, 5, 6, 7}
	respCh := make(chan model.MongoRecruitment)
	var wg sync.WaitGroup
	for _, p := range positions {
		for _, c := range careers {
			time.Sleep(5 * time.Second)
			wg.Add(1)
			go func(c int, p string) {
				defer wg.Done()
				m := e.crwaler(c, p)
				d := model.MongoRecruitment{Data: m, Position: p, Career: c}
				respCh <- d
			}(c, p)
		}
	}

	go func() {
		wg.Wait()
		close(respCh)
	}()

	coll := e.db.Database("db").Collection("recruitments")
	for resp := range respCh {
		recruitments := bson.D{{"data", resp.Data}, {"c", resp.Career}, {"p", resp.Position}}

		coll.InsertOne(context.TODO(), recruitments)
	}
}

func (e *Engine) GetRecruitment(career int, position string) []*model.Recruitment {
	coll := e.db.Database("db").Collection("recruitments")
	r := coll.FindOne(context.Background(), bson.M{"c": career, "p": position}, &options.FindOneOptions{
		Sort: bson.M{"_id": -1},
	})
	res := model.MongoRecruitment{}
	r.Decode(&res)
	fmt.Println(res.Id)
	return res.Data
}

func (e *Engine) crwaler(career int, position string) []*model.Recruitment {
	var wg sync.WaitGroup
	result := make([]*model.Recruitment, 0)
	respCh := make(chan []*model.Recruitment)
	for _, v := range e.Providers {
		wg.Add(1)
		go func(provider provider.Provider) {
			defer wg.Done()
			result := provider.GetRecruitment(career, position)
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
