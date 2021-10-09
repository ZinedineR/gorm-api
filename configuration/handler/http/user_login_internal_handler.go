package http

import (
	"api-gorm-setting/entity"
	"api-gorm-setting/service"
	nethttp "net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

// CreateUserBodyRequest defines all body attributes needed to add User.
type CreateUserBodyRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Admin    bool   `json:"admin"`
}

// UserRowResponse defines all attributes needed to fulfill for User row entity.
type UserRowResponse struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Admin    bool      `json:"admin"`
}

// UserResponse defines all attributes needed to fulfill for pic User entity.
type UserDetailResponse struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Admin    bool      `json:"admin"`
}

func buildUserRowResponse(User *entity.User) UserRowResponse {
	form := UserRowResponse{
		Id:       User.Id,
		Username: User.Username,
		Password: User.Password,
		Admin:    User.Admin,
	}

	return form
}

func buildUserDetailResponse(User *entity.User) UserDetailResponse {
	form := UserDetailResponse{
		Id:       User.Id,
		Username: User.Username,
		Password: User.Password,
		Admin:    User.Admin,
	}

	return form
}

// QueryParamsUser defines all attributes for input query params
type QueryParamsUser struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaUser define attributes needed for Meta
type MetaUser struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaUser creates an instance of Meta response.
func NewMetaUser(limit, offset int, total int64) *MetaUser {
	return &MetaUser{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// UserHandler handles HTTP request related to user flow.
type UserHandler struct {
	service service.UserUseCase
}

// NewUserHandler creates an instance of UserHandler.
func NewUserHandler(service service.UserUseCase) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// Create handles article creation.
// It will reject the request if the request doesn't have required data,
func (handler *UserHandler) CreateUser(echoCtx echo.Context) error {
	var form CreateUserBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	UserEntity := entity.NewUser(
		uuid.Nil,
		form.Username,
		form.Password,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), UserEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", UserEntity)
	return echoCtx.JSON(res.Status, res)
}
