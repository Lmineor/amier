package response

import (
	"ziyue/model"
)

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
