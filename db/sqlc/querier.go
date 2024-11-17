// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	AddImagesToCar(ctx context.Context, arg AddImagesToCarParams) (Cars, error)
	CreateCar(ctx context.Context, arg CreateCarParams) (Cars, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Sessions, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (Users, error)
	CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParams) (VerifyEmails, error)
	DeleteCar(ctx context.Context, id int32) error
	GetCarById(ctx context.Context, id int32) (Cars, error)
	GetSession(ctx context.Context, id uuid.UUID) (Sessions, error)
	GetUser(ctx context.Context, username string) (Users, error)
	GetUserOwnedCars(ctx context.Context, ownerID int32) ([]Cars, error)
	// go-type: title string
	// go-type: description string
	// go-type: model_name string
	// go-type: model_id string
	// go-type: color string
	// go-type: tags []string
	SearchCarsFTS(ctx context.Context, arg SearchCarsFTSParams) ([]Cars, error)
	UpdateCar(ctx context.Context, arg UpdateCarParams) (Cars, error)
	UpdateVerifyEmail(ctx context.Context, arg UpdateVerifyEmailParams) (VerifyEmails, error)
}

var _ Querier = (*Queries)(nil)