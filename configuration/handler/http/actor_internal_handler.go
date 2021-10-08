package http

import (
	"api-gorm-setting/entity"
	"api-gorm-setting/service"
	"net/http"
	nethttp "net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CreateActorBodyRequest defines all body attributes needed to add Actor.
type CreateActorBodyRequest struct {
	TV_id uuid.UUID `json:"tv_id" binding:"required"`
	Name  string    `json:"name" binding:"required"`
}

// ActorRowResponse defines all attributes needed to fulfill for Actor row entity.
type ActorRowResponse struct {
	Id    uuid.UUID `json:"id"`
	TV_id uuid.UUID `json:"tv_id"`
	Name  string    `json:"name"`
}

// ActorResponse defines all attributes needed to fulfill for pic Actor entity.
type ActorDetailResponse struct {
	Id    uuid.UUID `json:"id"`
	TV_id uuid.UUID `json:"tv_id"`
	Name  string    `json:"name"`
}

func buildActorRowResponse(Actor *entity.Actor) ActorRowResponse {
	form := ActorRowResponse{
		Id:    Actor.Id,
		TV_id: Actor.TV_id,
		Name:  Actor.Name,
	}

	return form
}

func buildActorDetailResponse(Actor *entity.Actor) ActorDetailResponse {
	form := ActorDetailResponse{
		Id:    Actor.Id,
		TV_id: Actor.TV_id,
		Name:  Actor.Name,
	}

	return form
}

// QueryParamsActor defines all attributes for input query params
type QueryParamsActor struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaActor define attributes needed for Meta
type MetaActor struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaActor creates an instance of Meta response.
func NewMetaActor(limit, offset int, total int64) *MetaActor {
	return &MetaActor{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// ActorHandler handles HTTP request related to user flow.
type ActorHandler struct {
	service service.ActorUseCase
}

// NewActorHandler creates an instance of ActorHandler.
func NewActorHandler(service service.ActorUseCase) *ActorHandler {
	return &ActorHandler{
		service: service,
	}
}

// Create handles article creation.
// It will reject the request if the request doesn't have required data,
func (handler *ActorHandler) CreateActor(echoCtx echo.Context) error {
	var form CreateActorBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	ActorEntity := entity.NewActor(
		uuid.Nil,
		form.TV_id,
		form.Name,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), ActorEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", ActorEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *ActorHandler) GetListActor(echoCtx echo.Context) error {
	var form QueryParamsActor
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	Actor, err := handler.service.GetListActor(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Actor)
	return echoCtx.JSON(res.Status, res)

}

func (handler *ActorHandler) GetDetailActor(echoCtx echo.Context) error {
	idParam := echoCtx.Param("id")
	if len(idParam) == 0 {
		errorResponse := buildErrorResponse(nil, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	Actor, err := handler.service.GetDetailActor(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Actor)
	return echoCtx.JSON(res.Status, res)
}

func (handler *ActorHandler) UpdateActor(echoCtx echo.Context) error {
	var form CreateActorBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	idParam := echoCtx.Param("id")

	if len(idParam) == 0 {
		errorResponse := buildErrorResponse(nil, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	_, err = handler.service.GetDetailActor(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	ActorEntity := entity.NewActor(
		id,
		form.TV_id,
		form.Name,
	)

	if err := handler.service.UpdateActor(echoCtx.Request().Context(), ActorEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *ActorHandler) DeleteActor(echoCtx echo.Context) error {
	idParam := echoCtx.Param("id")
	if len(idParam) == 0 {
		errorResponse := buildErrorResponse(nil, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	err = handler.service.DeleteActor(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
