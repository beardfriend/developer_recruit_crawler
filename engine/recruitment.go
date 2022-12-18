package engine

import (
	"sync"

	"dev_recruitment_crawler/model"
	"dev_recruitment_crawler/provider"
	"dev_recruitment_crawler/provider/jumpit"
	"dev_recruitment_crawler/provider/wanted"
)

type Engine struct {
	Providers map[string]provider.Provider
}

func NewEngine() *Engine {
	providers := make(map[string]provider.Provider, 0)
	providers["jumpit"] = jumpit.NewJumpit()
	providers["wanted"] = wanted.NewWanted()
	return &Engine{
		Providers: providers,
	}
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
