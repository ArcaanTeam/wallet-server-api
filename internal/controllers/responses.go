package controllers

// SuccessResponse standard API response
type SuccessResponse struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"Operation successful"`
	Data    interface{} `json:"data"`
}

// ErrorResponse standard API response
type ErrorResponse struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Operation failed"`
	Error   string `json:"error"`
}

func NewSuccessResponse(data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Success: true,
		Data:    data,
		Message: "Operation successful",
	}
}

func NewErrorResponse(err error) *ErrorResponse {
	appErr, ok := err.(interface {
		GetCode() int
		GetMessage() string
	})
	if ok {
		return &ErrorResponse{
			Success: false,
			Message: "Operation failed",
			Error:   appErr.GetMessage(),
		}
	}
	return &ErrorResponse{
		Success: false,
		Message: "Operation failed",
		Error:   err.Error(),
	}
}

// PaginatedResponse paginated data response
type paginatedResponse struct {
	Data       interface{} `json:"data"`
	Pagination pagination  `json:"pagination"`
}

type pagination struct {
	Page     int `json:"page" example:"1"`
	PageSize int `json:"pageSize" example:"10"`
	Total    int `json:"total" example:"100"`
}

func NewPaginatedResponse(data interface{}, message string) *SuccessResponse {
	// TODO: implement the pagination
	defaultPage := 1
	defaultPageSize := 10
	defaultTotal := 1

	return &SuccessResponse{
		Success: true,
		Message: message,
		Data: paginatedResponse{
			Data: data,
			Pagination: pagination{
				Page:     defaultPage,
				PageSize: defaultPageSize,
				Total:    defaultTotal,
			},
		},
	}
}
