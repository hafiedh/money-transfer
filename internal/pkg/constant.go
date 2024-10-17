package pkg

type (
	DefaultResponse struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
		Data    any    `json:"data"`
	}
)
