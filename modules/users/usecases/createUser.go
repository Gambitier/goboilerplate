package usecases

import (
	"context"
	"database/sql"
	"log"

	"github.com/gambitier/gocomm/db/dal/authors"
	"github.com/gambitier/gocomm/modules/users/dto"
	"github.com/gambitier/gocomm/modules/users/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
)

func CreateUser(authorsQueries *authors.Queries, requestData *dto.CreateUserRequest) (*authors.Author, *fiber.Error) {
	log.Print(requestData)

	// create an author
	insertedAuthor, err := authorsQueries.Create(context.TODO(), authors.CreateParams{
		Name: requestData.Name,
		Bio:  pgtype.Text{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})

	if err != nil {
		log.Println(err)
		var useCaseError *fiber.Error = nil

		switch err {
		case sql.ErrNoRows:
			useCaseError = errors.UserNotFound
		default:
			useCaseError = errors.FailedToCreateUser
		}

		return nil, useCaseError
	}

	return &insertedAuthor, nil
}
