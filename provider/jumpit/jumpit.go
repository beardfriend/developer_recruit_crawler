package jumpit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"sync"

	"dev_recruitment_crawler/model"
)

type Jumpit struct{}

func NewJumpit() *Jumpit {
	return &Jumpit{}
}

const baseUrl = "https://api.jumpit.co.kr/api/positions"

func (j *Jumpit) jobCodeName(name string) int {
	switch name {
	case "backend":
		return 1
	case "frontend":
		return 2
	case "dataEngineer":
		return 19
	default:
		return 0
	}
}

type Position struct {
	ID               int64    `json:"id"`
	JobCategory      string   `json:"jobCategory"`
	Logo             string   `json:"logo"`
	ImagePath        string   `json:"imagePath"`
	Title            string   `json:"title"`
	CompanyName      string   `json:"companyName"`
	TechStacks       []string `json:"techStacks"`
	ScrapCount       int64    `json:"scrapCount"`
	ViewCount        int64    `json:"viewCount"`
	Newcomer         bool     `json:"newcomer"`
	MinCareer        int64    `json:"minCareer"`
	MaxCareer        int64    `json:"maxCareer"`
	Locations        []string `json:"locations"`
	AlwaysOpen       bool     `json:"alwaysOpen"`
	ClosedAt         string   `json:"closedAt"`
	CompanyProfileID int64    `json:"companyProfileId"`
	Celebration      int64    `json:"celebration"`
	Scraped          bool     `json:"scraped"`
}

type response struct {
	Result struct {
		TotalCount  int64      `json:"totalCount"`
		Page        int64      `json:"page"`
		Keyword     string     `json:"keyword"`
		KeywordType string     `json:"keywordType"`
		Positions   []Position `json:"positions"`
	} `json:"result"`
}

func (j *Jumpit) GetRecruitment(minCareer int, job string) []*model.Recruitment {
	resp := j.get(minCareer, job, 1)

	page := math.Ceil(float64(resp.Result.TotalCount) / 15)
	response := make([]*model.Recruitment, 0)
	response = j.addResponse(resp.Result.Positions, response)

	var wg sync.WaitGroup

	for i := 2; i <= int(page); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			resp := j.get(minCareer, job, i)

			response = j.addResponse(resp.Result.Positions, response)
		}(i)
	}
	wg.Wait()

	return response
}

func (j *Jumpit) addResponse(data []Position, response []*model.Recruitment) []*model.Recruitment {
	for _, v := range data {
		location := ""
		if len(v.Locations) > 0 {
			location = v.Locations[0]
		}

		response = append(response, &model.Recruitment{
			Title:       v.Title,
			Provider:    "jumpit",
			Url:         fmt.Sprintf("https://www.jumpit.co.kr/position/%d", v.ID),
			CompanyName: v.CompanyName,
			Location:    location,
			ImageUrl:    v.Logo,
		})
	}

	return response
}

func (j *Jumpit) get(minCareer int, job string, pageNo int) *response {
	req, _ := http.NewRequest(http.MethodGet, baseUrl, nil)
	query := req.URL.Query()
	query.Add("sort", "rsp_rate")
	query.Add("highlight", "false")
	query.Add("page", fmt.Sprint(pageNo))
	query.Add("jobCategory", fmt.Sprint(j.jobCodeName(job)))
	if minCareer != 0 {
		query.Add("career", fmt.Sprint(minCareer))
	}

	req.URL.RawQuery = query.Encode()

	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	result := new(response)
	json.Unmarshal(data, result)

	return result
}
