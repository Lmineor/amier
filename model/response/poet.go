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
