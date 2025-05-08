package user

import (
	"Chat-Websocket/monitoring"
	"Chat-Websocket/pkg/authPkg"
	"Chat-Websocket/pkg/loggerPkg"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type userHandler struct {
	userSvc IUserService
	logger  loggerPkg.ILogger
}

func NewHandler(s IUserService, logger loggerPkg.ILogger) IUserHandler {
	return &userHandler{
		userSvc: s, logger: logger,
	}
}

// CreateUser godoc
// @Summary      Register a new user
// @Description  Create and register a new user in the system
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      CreateUserReq             true  "User data"
// @Success      201   {object}  map[string]interface{}    "Created user"
// @Failure      400   {object}  map[string]string         "Invalid Input Data"
// @Failure      500   {object}  map[string]string         "Internal Server Error"
// @Router       /api/auth/register [post]
func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var input CreateUserReq

	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input data: %s", err.Error())})
		return
	}

	res, err := h.userSvc.CreateUserService(c.Request.Context(), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create user: %s", err.Error())})
		return
	}

	monitoring.CreateCounter.Inc()

	c.JSON(http.StatusCreated, gin.H{"data": res})
}

// ListUser godoc
// @Summary      Get list of users
// @Description  Retrieve all users from the system
// @Tags         auth
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "List of users"
// @Failure      500  {object}  map[string]string       "Internal Server Error"
// @Router       /api/auth/users [get]
func (h *userHandler) ListUserHandler(c *gin.Context) {

	users, err := h.userSvc.ListUserService(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	monitoring.GetUsersCounter.Inc()

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// LoginUser godoc
// @Summary      User login
// @Description  Authenticate a user and generate a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      LoginUserReq             true  "User login data"
// @Success      200   {object}  map[string]interface{}    "Login successful"
// @Failure      400   {object}  map[string]string         "Invalid Input Data"
// @Failure      401   {object}  map[string]string         "Unauthorized - Invalid credentials"
// @Failure      500   {object}  map[string]string         "Internal Server Error"
// @Router       /api/auth/login [post]
func (h *userHandler) LoginUserHandler(c *gin.Context) {
	var input LoginUserReq

	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input data: %s", err.Error())})
		return
	}

	res, err := h.userSvc.LoginUserService(c.Request.Context(), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to login user: %s", err)})
		return
	}

	id, err := strconv.ParseInt(res.ID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	monitoring.LoginCounter.Inc()

	token, err := authPkg.CreateToken(id, res.Email)
	res.Token = token
	res.Email = input.Email
	c.JSON(http.StatusOK, gin.H{"data": res})
}

// EditUser godoc
// @Summary      Edit user information
// @Description  Update the information of the logged-in user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      EditUserReq             true  "User data to be updated"
// @Success      200   {object}  map[string]interface{}    "User updated successfully"
// @Failure      400   {object}  map[string]string         "Invalid Input Data"
// @Failure      401   {object}  map[string]string         "Unauthorized - Invalid user"
// @Failure      500   {object}  map[string]string         "Internal Server Error"
// @Security     BearerAuth
// @Router       /api/auth/edit [put]
func (h *userHandler) EditUserHandler(c *gin.Context) {
	userIDRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in token"})
		return
	}

	userID, ok := userIDRaw.(int64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid userID format"})
		return
	}

	var input EditUserReq
	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input data: %s", err.Error())})
		return
	}

	input.ID = int32(userID)

	updatedUser, err := h.userSvc.UpdateUserService(c.Request.Context(), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update user: %s", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": updatedUser})
}

// DeleteUser godoc
// @Summary      Delete user
// @Description  Delete the account of the logged-in user
// @Tags         auth
// @Produce      json
// @Success      204  {object}  map[string]string  "User deleted successfully"
// @Failure      401  {object}  map[string]string  "Unauthorized - Invalid user"
// @Failure      500  {object}  map[string]string  "Internal Server Error"
// @Security     BearerAuth
// @Router       /api/auth/delete [delete]
func (h *userHandler) DeleteUserHandler(c *gin.Context) {
	userIDRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in token"})
		return
	}

	userID, ok := userIDRaw.(int64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid userID format"})
		return
	}
	userIDInt32 := int32(userID)

	err := h.userSvc.DeleteUserService(c.Request.Context(), userIDInt32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Successfully deleted user"})

}
