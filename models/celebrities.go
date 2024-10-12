package models

type Celebrity struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int16  `json:"age"`
}
