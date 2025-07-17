package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"empty/internal/model"
	"empty/internal/service"
	"empty/internal/utils"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service}
}

// GetModelsByBrandID godoc
// @Summary      Get models by brand ID for create cars
// @Description  Returns a list of car models for a given brand ID, optionally filtered by text
// @Tags         users
// @Produce      json
// @Param        id    path      int     true   "Brand ID"
// @Param        user  body     model.UserCrete  false  "coroll"
// @Success      200   {array}  model.User
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user model.UserCrete

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := h.UserService.CreateUser(c.Request.Context(), &user)

	utils.GinResponse(c, data)
}

// GetModelsByUserID godoc
// @Summary      Get models by user ID for create cars
// @Description  Returns a list of car models for a given user ID, optionally filtered by text
// @Tags         users
// @Produce      json
// @Param        id    path      int     true   "User ID"
// @Param        text  query     string  false  "coroll"
// @Success      200   {array}  model.User
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  pkg.ErrorResponse
// @Failure		 403   {object} pkg.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/users/{id} [post]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		utils.GinResponse(c, &model.Response{
			Status: http.StatusBadRequest,
			Error:  err,
		})
		return
	}

	user, err := h.UserService.GetUserByID(c.Request.Context(), idInt)

	if err != nil {
		utils.GinResponse(c, &model.Response{
			Status: http.StatusInternalServerError,
			Error:  err,
		})
		return
	}

	utils.GinResponse(c, &model.Response{
		Status: 200,
		Data:   user,
	})
}
