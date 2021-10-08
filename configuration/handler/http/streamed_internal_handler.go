package http

import (
	"api-gorm-setting/entity"
	"api-gorm-setting/service"
	"net/http"
	nethttp "net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CreateStreamedBodyRequest defines all body attributes needed to add Streamed.
type CreateStreamedBodyRequest struct {
	Streamed_id uuid.UUID `json:"streamed_id" binding:"required"`
	Platform    string    `json:"platform" binding:"required"`
}

// StreamedRowResponse defines all attributes needed to fulfill for Streamed row entity.
type StreamedRowResponse struct {
	Id          uuid.UUID `json:"id"`
	Streamed_id uuid.UUID `json:"streamed_id"`
	Platform    string    `json:"platform"`
}

// StreamedResponse defines all attributes needed to fulfill for pic Streamed entity.
type StreamedDetailResponse struct {
	Id          uuid.UUID `json:"id"`
	Streamed_id uuid.UUID `json:"streamed_id"`
	Platform    string    `json:"platform"`
}

func buildStreamedRowResponse(Streamed *entity.Streamed) StreamedRowResponse {
	form := StreamedRowResponse{
		Id:          Streamed.Id,
		Streamed_id: Streamed.Streamed_id,
		Platform:    Streamed.Platform,
	}

	return form
}

func buildStreamedDetailResponse(Streamed *entity.Streamed) StreamedDetailResponse {
	form := StreamedDetailResponse{
		Id:          Streamed.Id,
		Streamed_id: Streamed.Streamed_id,
		Platform:    Streamed.Platform,
	}

	return form
}

// QueryParamsStreamed defines all attributes for input query params
type QueryParamsStreamed struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaStreamed define attributes needed for Meta
type MetaStreamed struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaStreamed creates an instance of Meta response.
func NewMetaStreamed(limit, offset int, total int64) *MetaStreamed {
	return &MetaStreamed{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// StreamedHandler handles HTTP request related to user flow.
type StreamedHandler struct {
	service service.StreamedUseCase
}

// NewStreamedHandler creates an instance of StreamedHandler.
func NewStreamedHandler(service service.StreamedUseCase) *StreamedHandler {
	return &StreamedHandler{
		service: service,
	}
}

// Create handles article creation.
// It will reject the request if the request doesn't have required data,
func (handler *StreamedHandler) CreateStreamed(echoCtx echo.Context) error {
	var form CreateStreamedBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	StreamedEntity := entity.NewStreamed(
		uuid.Nil,
		form.Streamed_id,
		form.Platform,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), StreamedEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", StreamedEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *StreamedHandler) GetListStreamed(echoCtx echo.Context) error {
	var form QueryParamsStreamed
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	Streamed, err := handler.service.GetListStreamed(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Streamed)
	return echoCtx.JSON(res.Status, res)

}

func (handler *StreamedHandler) GetDetailStreamed(echoCtx echo.Context) error {
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

	Streamed, err := handler.service.GetDetailStreamed(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Streamed)
	return echoCtx.JSON(res.Status, res)
}

func (handler *StreamedHandler) UpdateStreamed(echoCtx echo.Context) error {
	var form CreateStreamedBodyRequest
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

	_, err = handler.service.GetDetailStreamed(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	StreamedEntity := entity.NewStreamed(
		id,
		form.Streamed_id,
		form.Platform,
	)

	if err := handler.service.UpdateStreamed(echoCtx.Request().Context(), StreamedEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *StreamedHandler) DeleteStreamed(echoCtx echo.Context) error {
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

	err = handler.service.DeleteStreamed(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
