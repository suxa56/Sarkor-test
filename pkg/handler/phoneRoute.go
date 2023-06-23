package handler

import (
	Sarkor_test "Sarkor-test"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strconv"
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

// Get and display phone dto, search by name
func (h *Handler) getPhones(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	va := c.Request.URL.RawQuery
	params, err := url.ParseQuery(va)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "cannot parse params")
		return
	}
	phone := params.Get("q")

	phoneDto, err := h.services.GetPhoneInfo(phone)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"phone": phoneDto,
	})
}

// Update phone
func (h *Handler) editPhone(c *gin.Context) {
	var input Sarkor_test.UpdatePhone

	id, err := getUserId(c)
	if err != nil {
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.EditPhone(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}

// Delete phone, get phone id from url, get user id from context
func (h *Handler) deletePhone(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	phoneId := c.Param("phoneId")
	phoneIdInt, err := strconv.Atoi(phoneId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	err = h.services.DeletePhone(phoneIdInt, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
