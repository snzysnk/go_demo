package docs

import "x_go_swagger/swagger/api"

// swagger:route GET /user User getUserRequest
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
