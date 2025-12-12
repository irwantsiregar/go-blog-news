package response

type ErrorResponseDefault struct {
	Meta Meta `json:"meta"`
	// Data any  `json:"data,omitempty"`
}

type Meta struct {
	Status bool `json:"status"`
	// Code    int    `json:"code"`
	Message string `json:"message"`
}