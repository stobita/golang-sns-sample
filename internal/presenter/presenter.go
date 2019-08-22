package presenter

import (
	"github.com/stobita/golang-sns-sample/internal/model"
)

type listJSON struct {
	Items []interface{} `json:"items"`
}

type postJSON struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type errorJSON struct {
	Error errorItem `json:"error"`
}

type errorItem struct {
	Message string `json:"message"`
}

func ErrorResponse(text string) errorJSON {
	return errorJSON{
		Error: errorItem{
			Message: text,
		},
	}
}

func PostsResponse(list []*model.Post) (listJSON, error) {
	var o listJSON
	for _, v := range list {
		o.Items = append(o.Items, &postJSON{
			ID:      v.ID,
			Title:   v.Title,
			Content: v.Content,
		})
	}
	return o, nil
}
