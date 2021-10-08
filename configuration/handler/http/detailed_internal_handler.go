package http

import (
	"api-gorm-setting/entity"
	"api-gorm-setting/service"
	"net/http"
	nethttp "net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

// CreateDetailedBodyRequest defines all body attributes needed to add Detailed.
type CreateDetailedBodyRequest struct {
	Detailed_id uuid.UUID `json:"detailed_id" binding:"required"`
	Season      int       `json:"season" binding:"required"`
	Episodes    int       `json:"episodes" binding:"required"`
	Year        int       `json:"year" binding:"required"`
}

// DetailedRowResponse defines all attributes needed to fulfill for Detailed row entity.
type DetailedRowResponse struct {
	Id          uuid.UUID `json:"id"`
	Detailed_id uuid.UUID `json:"detailed_id"`
	Season      int       `json:"season"`
	Episodes    int       `json:"episodes"`
	Year        int       `json:"year"`
}

// DetailedResponse defines all attributes needed to fulfill for pic Detailed entity.
type DetailedDetailResponse struct {
	Id          uuid.UUID `json:"id"`
	Detailed_id uuid.UUID `json:"detailed_id"`
	Season      int       `json:"season"`
	Episodes    int       `json:"episodes"`
	Year        int       `json:"year"`
}

func buildDetailedRowResponse(Detailed *entity.Detailed) DetailedRowResponse {
	form := DetailedRowResponse{
		Id:          Detailed.Id,
		Detailed_id: Detailed.Detailed_id,
		Season:      Detailed.Season,
		Episodes:    Detailed.Episodes,
		Year:        Detailed.Year,
	}

	return form
}

func buildDetailedDetailResponse(Detailed *entity.Detailed) DetailedDetailResponse {
	form := DetailedDetailResponse{
		Id:          Detailed.Id,
		Detailed_id: Detailed.Detailed_id,
		Season:      Detailed.Season,
		Episodes:    Detailed.Episodes,
		Year:        Detailed.Year,
	}

	return form
}

// QueryParamsDetailed defines all attributes for input query params
type QueryParamsDetailed struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaDetailed define attributes needed for Meta
type MetaDetailed struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaDetailed creates an instance of Meta response.
func NewMetaDetailed(limit, offset int, total int64) *MetaDetailed {
	return &MetaDetailed{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// DetailedHandler handles HTTP request related to user flow.
type DetailedHandler struct {
	service service.DetailedUseCase
}

// NewDetailedHandler creates an instance of DetailedHandler.
func NewDetailedHandler(service service.DetailedUseCase) *DetailedHandler {
	return &DetailedHandler{
		service: service,
	}
}

// Create handles article creation.
// It will reject the request if the request doesn't have required data,
func (handler *DetailedHandler) CreateDetailed(echoCtx echo.Context) error {
	var form CreateDetailedBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	DetailedEntity := entity.NewDetailed(
		uuid.Nil,
		form.Detailed_id,
		form.Season,
		form.Episodes,
		form.Year,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), DetailedEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", DetailedEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *DetailedHandler) GetListDetailed(echoCtx echo.Context) error {
	var form QueryParamsDetailed
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	Detailed, err := handler.service.GetListDetailed(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Detailed)
	return echoCtx.JSON(res.Status, res)

}

func (handler *DetailedHandler) GetDetailDetailed(echoCtx echo.Context) error {
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

	Detailed, err := handler.service.GetDetailDetailed(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Detailed)
	return echoCtx.JSON(res.Status, res)
}

func (handler *DetailedHandler) UpdateDetailed(echoCtx echo.Context) error {
	var form CreateDetailedBodyRequest
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

	_, err = handler.service.GetDetailDetailed(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	DetailedEntity := entity.NewDetailed(
		id,
		form.Detailed_id,
		form.Season,
		form.Episodes,
		form.Year,
	)

	if err := handler.service.UpdateDetailed(echoCtx.Request().Context(), DetailedEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *DetailedHandler) DeleteDetailed(echoCtx echo.Context) error {
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

	err = handler.service.DeleteDetailed(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
