package web

type CategoryRequest struct {
	Name string `validate:"required" json:"name"`
}

type CategoryCreateResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CategoryUpdateServiceRequest struct {
	Name string `json:"name"`
}
