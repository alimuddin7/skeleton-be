package models

type MetaData struct {
	Code    string `json:"code"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Meta  MetaData    `json:"meta"`
	Count int64       `json:"count,omitempty"`
	Page  *Page       `json:"page,omitempty"`
}

type Page struct {
	CurrentPage  int `json:"curPage,omitempty"`
	PreviousPage int `json:"prevPage,omitempty"`
	NextPage     int `json:"nextPage,omitempty"`
}

type ErrorCodeElement struct {
	Key     string   `json:"key"`
	Content MetaData `json:"content"`
}
