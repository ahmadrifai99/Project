package order

import (
	"assigment2/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type orderController struct {
	service Service
}

func NewOrderController(service Service) *orderController {
	return &orderController{service}
}

func (h *orderController) GetAll(c *gin.Context) {
	orders, err := h.service.Get()
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Get Orders Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatted := OrdersFormatter(orders)
	response := helper.APIResponse("Get Orders Success", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}

func (h *orderController) Create(c *gin.Context) {
	var input CreateOrder
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Create Order Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	order, err := h.service.Create(input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Create Order Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatted := OrderFormatter(order)
	response := helper.APIResponse("Create Order Success", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}

func (h *orderController) Update(c *gin.Context) {
	var input UpdateOrder
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Update Order Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	order, err := h.service.Update(input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Update Order Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatted := OrderFormatter(order)
	response := helper.APIResponse("Update Order Success", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}

func (h *orderController) Delete(c *gin.Context) {
	var input DeleteOrder
	err := c.ShouldBindUri(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Delete Order Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	order, err := h.service.Delete(input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Delete Order Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatted := OrderFormatter(order)
	response := helper.APIResponse("Delete Order Success", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}
