package response

import (
	"encoding/json"
	"github.com/xuqingfeng/mockr/db"
	"github.com/xuqingfeng/mockr/util"
)

type Response struct {
	Status       Status   `json:"status"`
	Headers      []Header `json:"headers"`
	Cookies      []Cookie `json:"cookies"`
	ResponseBody string   `json:"responseBody"`
	Method       string   `json:"method"`
	MimeType     string   `json:"mimeType"`
}

type Status struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Cookie struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// TODO: 16/12/20 label
type Label string

const (
	responseBucketName = "response"
)

func (r *Response) AddResponse() error {

	d, err := db.New(db.DbPath, []byte(responseBucketName))
	if err != nil {
		return err
	}
	defer d.Close()

	uuid := util.GenerateUUID()
	rpInByte, err := json.Marshal(r)
	if err != nil {
		return err
	}
	err = d.Set(uuid, rpInByte, []byte(responseBucketName))
	if err != nil {
		return err
	}

	return nil
}
