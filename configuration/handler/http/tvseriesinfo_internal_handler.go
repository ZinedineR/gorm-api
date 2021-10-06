package http

import (
	"api-gorm-setting/entity"
	"api-gorm-setting/service"
	"net/http"
	nethttp "net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CreateTVBodyRequest defines all body attributes needed to add TV.
type CreateTVBodyRequest struct {
	Title    string `json:"title" binding:"required"`
	Producer string `json:"producer" binding:"required"`
}

// TVRowResponse defines all attributes needed to fulfill for TV row entity.
type TVRowResponse struct {
	Id       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Producer string    `json:"producer"`
}

// TVResponse defines all attributes needed to fulfill for pic TV entity.
type TVDetailResponse struct {
	Id       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Producer string    `json:"producer"`
}

func buildTVRowResponse(TV *entity.TV) TVRowResponse {
	form := TVRowResponse{
		Id:       TV.Id,
		Title:    TV.Title,
		Producer: TV.Producer,
	}

	return form
}

func buildTVDetailResponse(TV *entity.TV) TVDetailResponse {
	form := TVDetailResponse{
		Id:       TV.Id,
		Title:    TV.Title,
		Producer: TV.Producer,
	}

	return form
}

// QueryParamsTV defines all attributes for input query params
type QueryParamsTV struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaTV define attributes needed for Meta
type MetaTV struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaTV creates an instance of Meta response.
func NewMetaTV(limit, offset int, total int64) *MetaTV {
	return &MetaTV{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// TVHandler handles HTTP request related to user flow.
type TVHandler struct {
	service service.TVUseCase
}

// NewTVHandler creates an instance of TVHandler.
func NewTVHandler(service service.TVUseCase) *TVHandler {
	return &TVHandler{
		service: service,
	}
}

// Create handles article creation.
// It will reject the request if the request doesn't have required data,
func (handler *TVHandler) CreateTV(echoCtx echo.Context) error {
	var form CreateTVBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	TVEntity := entity.NewTV(
		uuid.Nil,
		form.Title,
		form.Producer,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), TVEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", TVEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *TVHandler) GetListTV(echoCtx echo.Context) error {
	var form QueryParamsTV
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	TV, err := handler.service.GetListTV(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", TV)
	return echoCtx.JSON(res.Status, res)

}

func (handler *TVHandler) GetDetailTV(echoCtx echo.Context) error {
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

	TV, err := handler.service.GetDetailTV(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", TV)
	return echoCtx.JSON(res.Status, res)
}

func (handler *TVHandler) UpdateTV(echoCtx echo.Context) error {
	var form CreateTVBodyRequest
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

	_, err = handler.service.GetDetailTV(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	TVEntity := &entity.TV{
		id,
		form.Title,
		form.Producer,
	}

	if err := handler.service.UpdateTV(echoCtx.Request().Context(), TVEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *TVHandler) DeleteTV(echoCtx echo.Context) error {
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

	err = handler.service.DeleteTV(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
