package request

type Poem struct {
	Poem       string   `json:"poem"`
	Paragraphs []string `json:"paragraphs"`
	Poet       string   `json:"poet"`
	Dynasty    string   `json:"dynasty"`
}
