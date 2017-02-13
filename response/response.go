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

func (r *Response) AddResponse() ([]byte, error) {

	d, err := db.New(db.DbPath, []byte(responseBucketName))
	if err != nil {
		return nil, err
	}
	defer d.Close()

	uuid := util.GenerateUUID()
	rpInByte, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	err = d.Set(uuid, rpInByte, []byte(responseBucketName))
	if err != nil {
		return nil, err
	}

	return uuid, nil
}

func GetResponse(uuid []byte) (Response, error) {

	d, err := db.New(db.DbPath, []byte(responseBucketName))
	if err != nil {
		return Response{}, err
	}
	defer d.Close()

	ret, err := d.Get(uuid, []byte(responseBucketName))
	if err != nil {
		return Response{}, err
	}
	var rp Response
	err = json.Unmarshal(ret, &rp)
	if err != nil {
		return Response{}, err
	}

	return rp, nil
}
