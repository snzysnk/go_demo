package docs

import (
	"x_go_swagger/api"
)

// swagger:route GET /user user getUserRequest
// 查找用户
// responses:
//	200: getUserResponse
//	default: errResponse

// swagger:parameters getUserRequest
type GetUserRequest struct {
	// in.path
	// required: true
	Name string `json:"name"`
}

// swagger:response getUserResponse
type GetUserResponse struct {
	// in.body
	Body api.User
}

// swagger:response errResponse
type ErrResponse struct {
	// in.body
	Body api.ErrResponse
}
