package handlers

import (
	"context"

	"github.com/gambitier/gocomm/modules/users/dto"
	"github.com/gambitier/gocomm/modules/users/proto"
	"github.com/gambitier/gocomm/modules/users/usecases"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateUser creates a new user.
// @Summary Create a new user
// @Description Create a new user based on the provided user data
// @Tags Users
// @Accept json
// @Produce json
// @Param body body dto.CreateUserRequest true "User data to create"
// @Success 201 {object} string //TODO: fix response type
// @Router /users/ [post]
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	requestData := c.Locals("body").(*dto.CreateUserRequest)
	insertedAuthor, err := usecases.CreateUser(h.dbRepo.AuthorsQueries, requestData)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"message": err.Message,
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"data":    insertedAuthor,
	})
}

func (h *UserServiceServer) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	insertedAuthor, err := usecases.CreateUser(h.AppContext.DbRepo.AuthorsQueries, req.ToDto())
	if err != nil {
		return nil, status.Error(
			codes.Internal,
			err.Error(),
		)
	}

	return &proto.CreateUserResponse{
		ID:   insertedAuthor.ID,
		Name: insertedAuthor.Name,
		Bio:  "TODO",
	}, nil
}
