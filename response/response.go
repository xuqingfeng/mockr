package response

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

func (r *Response) AddResponse() {

}
