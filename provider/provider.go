package provider

import "dev_recruitment_crawler/model"

type Provider interface {
	GetRecruitment(minCareer int, job string) []*model.Recruitment
}
