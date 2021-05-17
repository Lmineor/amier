package response

import "ziyue/model"

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
	UUID       string   `json:"uuid"`
	Paragraphs []string `json:"paragraphs"`
	Poem       string   `json:"poem"`
	PoetUUID   string   `json:"poet"`
}
