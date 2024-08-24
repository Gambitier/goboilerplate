package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)

type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

type ErrorList []ErrorResponse

func (e ErrorList) Error() string {
	errJSON, _ := json.Marshal(e)
	return string(errJSON)
}

// TODO: move this middleware to separate pkg, since init is done for pkg
func init() {
	validate = validator.New()

	// Set up the English translator
	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	trans, _ = uni.GetTranslator("en")

	// Register English translations
	en_translations.RegisterDefaultTranslations(validate, trans)
}

// ValidateRequest creates a middleware to validate the incoming request against the given DTO type
func ValidateRequest(dtoType any) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Create a new instance of the DTO type
		req := reflect.New(reflect.TypeOf(dtoType)).Interface()

		// Parse the request body into the DTO
		if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		if err := ValidateDto(req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errors": err,
			})
		}

		// Store the validated request in locals
		c.Locals("body", req)
		return c.Next()
	}
}

func ValidateDto(dto any) ErrorList {
	err := validate.Struct(dto)
	if err != nil {
		errors := mapErrors(err)
		return errors
	}
	return nil
}

func mapErrors(err error) ErrorList {
	var errors ErrorList
	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, ErrorResponse{
			Field:   err.Field(),
			Message: err.Translate(trans),
		})
	}
	return errors
}

type GrpcRequestProto interface {
	ToDto() any
}

func ValidationInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		protoReq, ok := req.(GrpcRequestProto)
		if !ok {
			return nil, status.Error(codes.Internal, "request does not implement GrpcRequestProto interface")
		}

		if err := ValidateDto(protoReq.ToDto()); err != nil {
			return nil, status.Error(
				codes.InvalidArgument,
				err.Error(),
			)
		}
		resp, err := handler(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
