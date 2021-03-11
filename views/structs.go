package views

type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

type TodoRequest struct {
	Name   string `json:"name"`
	Todo   string `json:"todo"`
	Status int    `json:"status"`
}

type TodoReponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Todo   string `json:"todo"`
	Status int    `json:"status"`
}

type TodoUpdate struct {
	Id     int    `json:"id"`
	Status int    `json:"status"`
	Name   string `json:"name"`
	Todo   string `json:"todo"`
}

type TodoDelete struct {
	Id int `json:"id"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
