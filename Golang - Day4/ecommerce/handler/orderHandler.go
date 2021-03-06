package handler

import (
	"ecommerce/application"
	"ecommerce/domain/entity"
	"ecommerce/infrastructure/concurrency"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderService struct {
	Order    application.OrderAppInterface
	Customer application.CustomerAppInterface
	Product  application.ProductAppInterface
}

func NewOrderService(orderApp application.OrderAppInterface,
	customerApp application.CustomerAppInterface,
	productApp application.ProductAppInterface) *OrderService {
	return &OrderService{Order: orderApp, Customer: customerApp, Product: productApp}
}

func (od *OrderService) AddMultipleOrders(c *gin.Context) {
	var orderRecieved entity.OrderRecieved
	var result []entity.OrderDetail

	// Bind request body
	if err := c.ShouldBindJSON(&orderRecieved); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate customer id
	_, err := od.Customer.GetCustomerByID(orderRecieved.CustomerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add to order table
	order, e := od.Order.AddOrder(&entity.Order{OrderId: orderRecieved.OrderId, CustomerId: orderRecieved.CustomerId})
	if e != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": e.Error()})
		return
	}

	// Add order details
	for _, ord := range orderRecieved.Specs {

		// Validate order spec
		if validationErr := ord.ValidateInput(); validationErr != nil {
			continue
		}
		status := "not placed"
		product, err := od.Product.GetProductByID(ord.ProductId)

		// Place order
		if err == nil && product.ProductQuantity-ord.ProductQuantity >= 0 {
			product.ProductQuantity -= ord.ProductQuantity

			// Apply lock
			if locked := concurrency.Mutex.Lock(product.ProductId); !locked {
				c.JSON(http.StatusLocked, gin.H{"error": "can't apply lock"})
				return
			}
			defer concurrency.Mutex.Unlock(product.ProductId)
			// time.Sleep(10 * time.Second)

			// Patch product
			if _, patchErr := od.Product.UpdateProduct(product); patchErr == nil {
				status = "placed"
			}
		}
		ordPlaced, err := od.Order.AddOrderDetail(&entity.OrderDetail{
			OrderId:         order.OrderId,
			OrderStatus:     status,
			ProductId:       ord.ProductId,
			QuantityOrdered: ord.ProductQuantity,
		})
		if err == nil {
			result = append(result, *ordPlaced)
		}

	}
	c.JSON(http.StatusOK, result)
}

func (od *OrderService) GetOrderByID(c *gin.Context) {
	param := c.Params.ByName("id")

	fmt.Print("Param : ", param)
	id, e := uuid.Parse(param)

	// Parse uuid recieved in param
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
	} else {
		if order, err := od.Order.GetOrderByID(id); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSONP(http.StatusOK, order)
		}
	}

}
