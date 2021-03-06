package handler

import (
	"ecommerce/application"
	"ecommerce/domain/entity"
	"ecommerce/handler/utils"
	"ecommerce/infrastructure/concurrency"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductService struct {
	Product application.ProductAppInterface
	Retail  application.RetailerAppInterface
}

func NewProductService(pdt application.ProductAppInterface, rt application.RetailerAppInterface) *ProductService {
	return &ProductService{Product: pdt, Retail: rt}
}

func (pd *ProductService) AddProduct(c *gin.Context) {
	var product entity.Product
	// Bind request body
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if retailer valid
	if _, err := pd.Retail.GetRetailerByID(uuid.UUID(product.RetailerId)); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	// Validate product
	validationErr := product.ValidateInput()
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
		return
	}

	productNew, err := pd.Product.AddProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	response := utils.FormatProductAddedResponse(*productNew)
	c.JSON(http.StatusCreated, response)
}

func (pd *ProductService) GetAllProducts(c *gin.Context) {

	products, err := pd.Product.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := utils.FormatProductListResponse(products)
	c.JSONP(http.StatusOK, gin.H{"products": result})
}

func (pd *ProductService) GetProductByID(c *gin.Context) {
	param := c.Params.ByName("id")

	fmt.Print("Param : ", param)
	id, e := uuid.Parse(param)

	// Parse uuid recieved in param
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
	} else {
		if product, err := pd.Product.GetProductByID(id); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSONP(http.StatusOK, product)
		}
	}

}

func (pd *ProductService) UpdateProduct(c *gin.Context) {
	var updateProduct entity.ProductPatch

	// Parse uuid
	param := c.Params.ByName("id")
	fmt.Print("Param : ", param)
	id, e := uuid.Parse(param)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	// Bind json
	if err := c.ShouldBindJSON(&updateProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get product
	newProduct, err := pd.Product.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusPreconditionFailed, gin.H{"error": err.Error()})
		return
	}
	newProduct.ProductQuantity = updateProduct.ProductQuantity
	newProduct.ProductPrice = updateProduct.ProductPrice

	fmt.Println(newProduct)
	if validationErr := newProduct.ValidateInput(); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
		return
	}

	// Concurrency management
	if locked := concurrency.Mutex.Lock(newProduct.ProductId); !locked {
		c.JSON(http.StatusLocked, gin.H{"error": "can't apply lock"})
		return
	}
	defer concurrency.Mutex.Unlock(newProduct.ProductId)
	// time.Sleep(5 * time.Second)

	// Patch product
	if _, err := pd.Product.UpdateProduct(newProduct); err != nil {
		c.JSON(http.StatusNotModified, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, newProduct)
	}
}
