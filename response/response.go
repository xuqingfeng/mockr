package response

type Response struct {
	Status  Status   `json:"status"`
	Headers []Header `json:"headers"`
	Cookies []Cookie `json:"cookies"`
	Body    Body     `json:"body"`
	Method  Method   `json:"method"`
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

type Body struct {
	Content  string `json:"content"`
	MimeType string `json:"mimeType"`
}

type Method string

// TODO: 16/12/20 label
type Label string
