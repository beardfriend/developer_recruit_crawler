package wanted

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"dev_recruitment_crawler/model"
)

type Wanted struct{}

func NewWanted() *Wanted {
	return &Wanted{}
}

const baseUrl = "https://www.wanted.co.kr/api/v4/jobs"

func (j *Wanted) jobCodeName(name string) int {
	switch name {
	case "backend":
		return 872
	case "frontend":
		return 669
	case "dataEngineer":
		return 655
	default:
		return 518
	}
}

type response struct {
	ModelStatus           interface{} `json:"model_status"`
	Links                 Links       `json:"links"`
	IsCallableExternalJob bool        `json:"is_callable_external_job"`
	DataModel             interface{} `json:"data_model"`
	Data                  []Datum     `json:"data"`
	IsScore               bool        `json:"is_score"`
}

type Datum struct {
	Status         string        `json:"status"`
	Reward         Reward        `json:"reward"`
	IsLike         bool          `json:"is_like"`
	IsBookmark     bool          `json:"is_bookmark"`
	Company        Company       `json:"company"`
	TitleImg       Img           `json:"title_img"`
	CompareCountry bool          `json:"compare_country"`
	DueTime        *string       `json:"due_time"`
	LikeCount      int64         `json:"like_count"`
	ID             int64         `json:"id"`
	LogoImg        Img           `json:"logo_img"`
	Address        Address       `json:"address"`
	MatchingScore  interface{}   `json:"matching_score"`
	Position       string        `json:"position"`
	Score          interface{}   `json:"score"`
	CategoryTags   []CategoryTag `json:"category_tags"`
}

type Address struct {
	Country  string `json:"country"`
	Location string `json:"location"`
}

type CategoryTag struct {
	ParentID int64 `json:"parent_id"`
	ID       int64 `json:"id"`
}

type Company struct {
	ID                       int64                    `json:"id"`
	IndustryName             string                   `json:"industry_name"`
	ApplicationResponseStats ApplicationResponseStats `json:"application_response_stats"`
	Name                     string                   `json:"name"`
}

type ApplicationResponseStats struct {
	AvgRate       int64  `json:"avg_rate"`
	Level         string `json:"level"`
	DelayedCount  int64  `json:"delayed_count"`
	AvgDay        *int64 `json:"avg_day"`
	RemainedCount int64  `json:"remained_count"`
	Type          string `json:"type"`
}

type Img struct {
	Origin string `json:"origin"`
	Thumb  string `json:"thumb"`
}

type Reward struct {
	FormattedTotal       string `json:"formatted_total"`
	FormattedRecommender string `json:"formatted_recommender"`
	FormattedRecommendee string `json:"formatted_recommendee"`
}

type Links struct {
	Prev string  `json:"prev"`
	Next *string `json:"next"`
}

func (j *Wanted) GetRecruitment(minCareer int, job string) []*model.Recruitment {
	response := make([]*model.Recruitment, 0)
	var wg sync.WaitGroup

	i := 0
	chanResponseLength := make(chan int)

	for {
		wg.Add(1)
		i++
		go func(i int) {
			defer wg.Done()
			resp := j.get(minCareer, job, i)
			response = j.addResponse(resp.Data, response)
			chanResponseLength <- len(resp.Data)
		}(i)
		length := <-chanResponseLength

		if length < 30 {
			break
		}
	}

	wg.Wait()
	return response
}

func (j *Wanted) addResponse(data []Datum, response []*model.Recruitment) []*model.Recruitment {
	for _, v := range data {
		if v.DueTime != nil {
			continue
		}

		if v.Status != "active" {
			continue
		}
		response = append(response, &model.Recruitment{
			Title:       v.Position,
			Provider:    "wanted",
			Url:         fmt.Sprintf("https://www.wanted.co.kr/wd/%d", v.ID),
			CompanyName: v.Company.Name,
			Location:    v.Address.Location,
			ImageUrl:    v.LogoImg.Thumb,
		})
	}

	return response
}

func (j *Wanted) get(minCareer int, job string, pageNo int) *response {
	limit := 30
	unix := time.Now().Unix()
	req, _ := http.NewRequest(http.MethodGet, baseUrl, nil)
	query := req.URL.Query()
	query.Add("country", "kr")
	query.Add("tag_type_ids", fmt.Sprint(j.jobCodeName(job))) // 개발직군
	query.Add("job_sort", "company.response_rate_order")
	query.Add("locations", "all")
	query.Add("years", "0")
	query.Add("years", fmt.Sprint(minCareer))
	query.Add("limit", fmt.Sprint(limit))
	query.Add("offset", fmt.Sprint(limit*pageNo))

	req.URL.RawQuery = query.Encode()
	req.URL.RawQuery = fmt.Sprint(unix) + "&" + req.URL.RawQuery

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
