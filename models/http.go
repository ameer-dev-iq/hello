package models

type BaseResponse[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}
