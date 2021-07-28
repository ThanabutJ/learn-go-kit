package main

//uppercaseRequest
type uppercaseRequest struct {
	S string `json:"s"`
}

//uppercaseResponse
type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

//countRequest
type countRequest struct {
	S string `json:"s"`
}

//countRepsonse
type countRepsonse struct {
	V int `json:"v"`
}
