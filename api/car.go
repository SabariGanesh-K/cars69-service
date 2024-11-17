package api

import (
	"errors"
	"net/http"
"database/sql"
	"github.com/gin-gonic/gin"
	db "github.com/SabariGanesh-K/cars69-service.git/db/sqlc"
	"github.com/SabariGanesh-K/cars69-service.git/token"
)

type CreateCarRequest struct {
	ID            int32    `json:"id"`
	OwnerID       int32    `json:"owner_id"`
	Title         string   `json:"title"`
	ModelID       string   `json:"model_id"`
	Color         string   `json:"color"`
	ModelName     string   `json:"model_name"`
	Description   string   `json:"description"`
	Tags          []string `json:"tags"`
	ListingStatus string   `json:"listing_status"`
	Images        []string `json:"images"`
	OwnerUsername string   `json:"owner_username"`
}

func (server *Server) createCar(ctx *gin.Context) {
	var req CreateCarRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCarParams{
		ID          : req.ID,
	OwnerID     : req.OwnerID,
	Title        : req.Title,
	ModelID     : req.ModelID,
	Color     :req.Color,
	ModelName    : req.ModelName,
	Description  : req.Description,
	Tags          : req.Tags,
	ListingStatus: req.ListingStatus,
	Images       : req.Images,
	OwnerUsername:req.OwnerUsername,
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if (req.OwnerUsername != authPayload.Username ){
		err := errors.New("account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}


	car, err := server.store.CreateCar(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, car)
}


type getCarbyIdRequest struct {
	id int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getCarById(ctx *gin.Context) {
	var req getCarbyIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	car, err := server.store.GetCarById(ctx, req.id)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if (car.OwnerUsername != authPayload.Username ){
		err := errors.New("account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}


	ctx.JSON(http.StatusOK, car)
}

type deleteCarRequest struct {
	id int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteCar(ctx *gin.Context) {
	var req deleteCarRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

 server.store.DeleteCar(ctx, req.id)
	// if err != nil {
	// 	if errors.Is(err, db.ErrRecordNotFound) {
	// 		ctx.JSON(http.StatusNotFound, errorResponse(err))
	// 		return
	// 	}

	// 	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	// 	return
	// }

	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	// if account.Owner != authPayload.Username {
	// 	err := errors.New("account doesn't belong to the authenticated user")
	// 	ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	// 	return
	// }

	ctx.JSON(http.StatusOK,1)
}


type getUserOwnedCarsRequest struct {
	ownerID int32 `uri:"ownerID" binding:"required,min=1"`
}

func (server *Server) getUserOwnedCars(ctx *gin.Context) {
	var req getUserOwnedCarsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetUserOwnedCars(ctx, req.ownerID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if (len(account)!=0 && account[0].OwnerUsername != authPayload.Username ){
		err := errors.New("account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type searchCarsFTSRequest struct {
	ID      int32          `json:"id"`
	Column2 sql.NullString `json:"column_2"`
	Offset  int32          `json:"offset"`
	Limit   int32          `json:"limit"`
}


func (server *Server) searchCarsFTS(ctx *gin.Context) {
	var req searchCarsFTSRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.SearchCarsFTSParams{
		ID          : req.ID,
		Column2          : req.Column2,
		Offset          : req.Offset,
		Limit          : req.Limit,
	

	}

	car, err := server.store.SearchCarsFTS(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if (car[0].OwnerUsername != authPayload.Username ){
		err := errors.New("account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, car)
}

type updateCarRequest struct {
	Title         string `json:"title"`
	Tags          []string       `json:"tags"`
	ListingStatus string `json:"listing_status"`
	Description   string `json:"description"`
	Color         string `json:"color"`
	ModelID       string `json:"model_id"`
	Images        []string       `json:"images"`
	ModelName     string `json:"model_name"`
	ID            int32          `json:"id"`
}
func (server *Server) updateCar(ctx *gin.Context) {
	var req updateCarRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCarParams{
		Title        : req.Title,
		ModelID     : req.ModelID,
		Color     :req.Color,
		ModelName    : req.ModelName,
		Description  : req.Description,
		Tags          : req.Tags,
		ListingStatus: req.ListingStatus,
	ID            :req.ID   ,
	Images:req.Images ,


	}

	cars, err := server.store.UpdateCar(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, cars)
}

