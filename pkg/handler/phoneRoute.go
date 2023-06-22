package handler

import (
	Sarkor_test "Sarkor-test"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Create new phone
func (h *Handler) createPhone(c *gin.Context) {
	var input Sarkor_test.Phone

	id, err := getUserId(c)
	if err != nil {
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	// set phone user id
	input.UserId = id
	phoneId, err := h.services.CreatePhone(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"phoneId": phoneId,
	})
}

func (h *Handler) getPhones(c *gin.Context) {

}

func (h *Handler) editPhone(c *gin.Context) {

}

func (h *Handler) deletePhone(c *gin.Context) {

}
