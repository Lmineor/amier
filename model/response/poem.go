package response

import "ziyue/model"

type PoetResponse struct {
	Poet model.Poet `json:"poet"`
}

type PoetsResponse struct {
	Poets []model.Poet `json:"poets"`
	Total int64        `json:"total"`
}
