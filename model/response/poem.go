package response

import (
	"ziyue/model"
)

type PoetsResponse struct {
	Poets []model.Poet `json:"poets"`
	Total int64        `json:"total"`
}

type PoetResponse struct {
	Poet    string   `json:"poet"`
	Dynasty string   `json:"dynasty"`
	Descb   string   `json:"descb"`
	Poems   []string `json:"poems"`
	UUID    string   `json:"uuid"`
}

type PoemResponse struct {
	Paragraphs []string `json:"paragraphs"`
	Poem       string   `json:"poem"`
	Like       uint     `json:"ilike"`
	UUID       string   `json:"uuid"`
	PoetUUID   string   `json:"poet"`
}

type PoemsResponse struct {
	Poems []model.Poem `json:"poems"`
}
