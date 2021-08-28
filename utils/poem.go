package utils

import (
	"strings"
	"ziyue/model"
	"ziyue/model/response"
)

func ParsePoemSplit(poem *model.Poem, pUUID string) *response.PoemResponse {
	return &response.PoemResponse{
		Poem:       poem.Poem,
		UUID:       poem.UUID,
		Paragraphs: strings.Split(poem.Paragraphs, "|"),
		PoetUUID:   pUUID,
		Like:       poem.Like}

}
