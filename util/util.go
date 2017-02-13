package util

import "github.com/satori/go.uuid"

type Msg struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GenerateUUID() []byte {

	u := uuid.NewV4().String()
	return []byte(u)
}
