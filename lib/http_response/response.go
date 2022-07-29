package http_response

type Response struct {
StatusCode	int				`json:"status_code"`
Message		string			`json:"message"`
Status		string			`json:"status"`
Timestamp	string			`json:"timestamp"`
Data		interface{}		`json:"data"`
}