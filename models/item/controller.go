package item

import (
	"assigment2/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type itemController struct {
	service Service
}

func NewItemController(service Service) *itemController {
	return &itemController{service}
}

func (h *itemController) Get(c *gin.Context) {
	items, err := h.service.Get()
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Get Items Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatted := ItemsFormatter(items)
	response := helper.APIResponse("Get Items Success", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}

func (h *itemController) Create(c *gin.Context) {
	var input CreateItem
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Get Items Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	item, err := h.service.Create(input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Get Items Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatted := ItemFormatter(item)
	response := helper.APIResponse("Create Item Success", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}
