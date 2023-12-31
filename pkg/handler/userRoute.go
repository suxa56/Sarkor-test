package handler

import (
	Sarkor_test "Sarkor-test"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register new user
func (h *Handler) register(c *gin.Context) {
	var input Sarkor_test.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type signInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Auth user
func (h *Handler) auth(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.GenerateToken(input.Login, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"SESSTOKEN": token,
	})

}

// Get user dto by name
func (h *Handler) getUser(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}
	name := c.Param("name")

	userInfo, err := h.services.GetUserInfo(name)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, userInfo)

}
