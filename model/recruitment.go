package model

import "time"

type Recruitment struct {
	Title       string
	Provider    string
	Url         string
	ImageUrl    string
	CompanyName string
	Location    string
}

type MongoRecruitment struct {
	Id        string `bson:"_id"`
	Data      []*Recruitment
	Position  string
	Career    int
	CreatedAt time.Time
}
