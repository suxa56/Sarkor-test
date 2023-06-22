package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createPhone(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getPhones(c *gin.Context) {

}

func (h *Handler) editPhone(c *gin.Context) {

}

func (h *Handler) deletePhone(c *gin.Context) {

}
