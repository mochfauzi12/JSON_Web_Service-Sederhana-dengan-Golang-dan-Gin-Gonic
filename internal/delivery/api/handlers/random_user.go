package handlers

import (
	interfaceservice "Web_Service/domain/interface_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	randomUserService interfaceservice.RandomUserService
}

func NewAPIHandler(randomUserService interfaceservice.RandomUserService) *APIHandler {
	return &APIHandler{
		randomUserService: randomUserService,
	}
}

func (h *APIHandler) GetPerson(c *gin.Context) {

	response, err := h.randomUserService.GetRandomUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)

}
