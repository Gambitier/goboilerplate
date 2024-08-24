package proto

import "github.com/gambitier/gocomm/modules/users/dto"

func (req *CreateUserRequest) ToDto() *dto.CreateUserRequest {
	requestData := &dto.CreateUserRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	return requestData
}
