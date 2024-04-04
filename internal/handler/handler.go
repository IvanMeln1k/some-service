package handler

import (
	"context"
	"fmt"
	"os"

	"github.com/IvanMeln1k/some-service/internal/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handler struct {
}

func (h *Handler) GetApiV1User(ctx context.Context,
	request GetApiV1UserRequestObject) (GetApiV1UserResponseObject, error) {
	return GetApiV1User200JSONResponse(domain.User{
		Id:       uuid.New(),
		Email:    "email@email.com",
		Name:     "Сергей",
		Username: "Platon Karataev",
	}), nil
}

func (h *Handler) PostApiV1User(ctx context.Context,
	request PostApiV1UserRequestObject) (PostApiV1UserResponseObject, error) {
	Id := uuid.New()
	return PostApiV1User200JSONResponse{
		Id: &Id,
	}, nil
}

func (h *Handler) GetApiV1Photo(ctx context.Context, request GetApiV1PhotoRequestObject) (GetApiV1PhotoResponseObject, error) {
	file, err := os.Open("assets/some-photo.jpeg")
	if err != nil {
		fmt.Println(err)
		return GetApiV1Photo500JSONResponse{
			Message: "Internal server error",
		}, nil
	}
	return GetApiV1Photo200ImagepngResponse{
		Body: file,
	}, nil
}

type Deps struct {
}

func NewHandler(deps Deps) *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())

	srv := NewStrictHandler(h, nil)
	RegisterHandlers(e, srv)

	return e
}
