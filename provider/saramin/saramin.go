package saramin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"sync"

	"dev_recruitment_crawler/model"

	"github.com/PuerkitoBio/goquery"
)

type Saramin struct{}

func NewSaramin() *Saramin {
	return &Saramin{}
}

func getCode(job string) string {
	switch job {
	case "backend":
		return "84"
	case "frontend":
		return "92"
	case "dataEngineer":
		return "83"
	default:
		return "0"
	}
}

const (
	baseUrl     = "https://www.saramin.co.kr/zf_user/jobs/list/domestic"
	totalApiurl = "https://www.saramin.co.kr/zf_user/jobs/api/get-search-count"
	detailurl   = "https://www.saramin.co.kr/zf_user/jobs/relay/view"
)

type Totalresponse struct {
	TodayCnt    int64  `json:"today_cnt"`
	ResultCnt   int64  `json:"result_cnt"`
	SearchQuery string `json:"search_query"`
}

func (j *Saramin) GetRecruitment(minCareer int, job string) []*model.Recruitment {
	total := j.total(minCareer, job)
	page := math.Ceil(float64(total) / 50)

	response := make([]*model.Recruitment, 0)
	var wg sync.WaitGroup
	for i := 1; i < int(page); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			resp := j.get(minCareer, job, i)
			response = append(response, resp...)
		}(i)
	}
	wg.Wait()

	return response
}

func (j *Saramin) total(minCareer int, job string) int64 {
	req, _ := http.NewRequest(http.MethodGet, totalApiurl, nil)
	query := req.URL.Query()
	query.Add("loc_mcd", "101000") // 서울
	query.Add("loc_cd", "102190")  // 판교
	query.Add("cat_kewd", getCode(job))
	if minCareer == 0 {
		query.Add("exp_cd", fmt.Sprint(1))
	} else {
		query.Add("exp_max", fmt.Sprint(minCareer)) // 경력
		query.Add("exp_cd", fmt.Sprint(2))
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

	result := new(Totalresponse)
	json.Unmarshal(data, result)

	return result.ResultCnt
}

func (j *Saramin) get(minCareer int, job string, pageNo int) (result []*model.Recruitment) {
	req, _ := http.NewRequest(http.MethodGet, baseUrl, nil)
	query := req.URL.Query()
	query.Add("loc_mcd", "101000") // 서울
	query.Add("loc_cd", "102190")  // 판교
	query.Add("cat_kewd", getCode(job))

	if minCareer == 0 {
		query.Add("exp_cd", fmt.Sprint(1))
	} else {
		query.Add("exp_max", fmt.Sprint(minCareer)) // 경력
		query.Add("exp_cd", fmt.Sprint(2))
	}

	query.Add("search_optional_item", "y")
	query.Add("search_done", "y")
	query.Add("panel_count", "y")
	query.Add("page_count", "50")
	query.Add("page", fmt.Sprint(pageNo))
	query.Add("preview", "y")

	req.URL.RawQuery = query.Encode()

	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// HTML 읽기
	html, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	section := html.Find("div#sri_section")
	list := section.Find("section.list_recruiting")
	l := list.Find("div.list_body")
	d := l.Find("div.list_item")
	d.Each(func(i int, s *goquery.Selection) {
		recruit := s.Find("div.notification_info")
		if len(recruit.Has("div.flag_reward_wrap").Nodes) == 0 {
			company := s.Find("div.company_nm")
			cName := company.Find("a.str_tit > span").Text()

			r := recruit.Find("div.job_tit > a.str_tit")
			title, _ := r.Attr("title")
			href, _ := r.Attr("href")

			cInfo := s.Find("div.company_info")
			area := cInfo.Find("p.work_place").Text()
			result = append(result, &model.Recruitment{
				Title:       title,
				Provider:    "saramin",
				Url:         fmt.Sprintf("https://www.saramin.co.kr%s", href),
				ImageUrl:    "https://play-lh.googleusercontent.com/C-Rk5j68xQIgi1apCuupseecXCquaNb-VZdnBecmjYK4_LHPg-ytgk7BTSe8JHSyjoY",
				CompanyName: cName,
				Location:    area,
			})
		}
	})

	return
}
