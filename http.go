package hello_world

type CreatePersonRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type CreatePersonResponse struct {
	Message string `json:"message"`
}
