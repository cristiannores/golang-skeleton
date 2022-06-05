package response

type ErrorResponse struct {
	Kind        string `json:"kind"`
	Description string `json:"description"`
}
