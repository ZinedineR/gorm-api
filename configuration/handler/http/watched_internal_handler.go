package http

import (
	"api-gorm-setting/entity"
	"api-gorm-setting/service"
	"net/http"
	nethttp "net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

// CreateWatchedBodyRequest defines all body attributes needed to add Watched.
type CreateWatchedBodyRequest struct {
	Watched_id uuid.UUID `json:"watched_id" binding:"required"`
	Season     int       `json:"season" binding:"required"`
	Episodes   int       `json:"episodes" binding:"required"`
}

// WatchedRowResponse defines all attributes needed to fulfill for Watched row entity.
type WatchedRowResponse struct {
	Id         uuid.UUID `json:"id"`
	Watched_id uuid.UUID `json:"watched_id"`
	Season     int       `json:"season"`
	Episodes   int       `json:"episodes"`
}

// WatchedResponse defines all attributes needed to fulfill for pic Watched entity.
type WatchedDetailResponse struct {
	Id         uuid.UUID `json:"id"`
	Watched_id uuid.UUID `json:"watched_id"`
	Season     int       `json:"season"`
	Episodes   int       `json:"episodes"`
}

func buildWatchedRowResponse(Watched *entity.Watched) WatchedRowResponse {
	form := WatchedRowResponse{
		Id:         Watched.Id,
		Watched_id: Watched.Watched_id,
		Season:     Watched.Season,
		Episodes:   Watched.Episodes,
	}

	return form
}

func buildWatchedDetailResponse(Watched *entity.Watched) WatchedDetailResponse {
	form := WatchedDetailResponse{
		Id:         Watched.Id,
		Watched_id: Watched.Watched_id,
		Season:     Watched.Season,
		Episodes:   Watched.Episodes,
	}

	return form
}

// QueryParamsWatched defines all attributes for input query params
type QueryParamsWatched struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaWatched define attributes needed for Meta
type MetaWatched struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaWatched creates an instance of Meta response.
func NewMetaWatched(limit, offset int, total int64) *MetaWatched {
	return &MetaWatched{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// WatchedHandler handles HTTP request related to user flow.
type WatchedHandler struct {
	service service.WatchedUseCase
}

// NewWatchedHandler creates an instance of WatchedHandler.
func NewWatchedHandler(service service.WatchedUseCase) *WatchedHandler {
	return &WatchedHandler{
		service: service,
	}
}

// Create handles article creation.
// It will reject the request if the request doesn't have required data,
func (handler *WatchedHandler) CreateWatched(echoCtx echo.Context) error {
	var form CreateWatchedBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	WatchedEntity := entity.NewWatched(
		uuid.Nil,
		form.Watched_id,
		form.Season,
		form.Episodes,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), WatchedEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", WatchedEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *WatchedHandler) GetListWatched(echoCtx echo.Context) error {
	var form QueryParamsWatched
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	Watched, err := handler.service.GetListWatched(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Watched)
	return echoCtx.JSON(res.Status, res)

}

func (handler *WatchedHandler) GetDetailWatched(echoCtx echo.Context) error {
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

	Watched, err := handler.service.GetDetailWatched(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Watched)
	return echoCtx.JSON(res.Status, res)
}

func (handler *WatchedHandler) UpdateWatched(echoCtx echo.Context) error {
	var form CreateWatchedBodyRequest
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

	_, err = handler.service.GetDetailWatched(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	WatchedEntity := entity.NewWatched(
		id,
		form.Watched_id,
		form.Season,
		form.Episodes,
	)

	if err := handler.service.UpdateWatched(echoCtx.Request().Context(), WatchedEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *WatchedHandler) DeleteWatched(echoCtx echo.Context) error {
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

	err = handler.service.DeleteWatched(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
