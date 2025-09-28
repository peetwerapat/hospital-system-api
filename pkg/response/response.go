package response

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"totalItems"`
	TotalPages int `json:"totalPages"`
}

type BaseHttpResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type HttpResponseData[T any] struct {
	BaseHttpResponse
	Data T `json:"data"`
}

type HttpResponseWithPagination[T any] struct {
	BaseHttpResponse
	Data       T           `json:"data"`
	Pagination *Pagination `json:"pagination"`
}
