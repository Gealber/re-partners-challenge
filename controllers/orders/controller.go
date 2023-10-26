package orders

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type dimensionsRepository interface {
	// retrieve current dimensions
	Get() ([]int, error)
	// update current dimensions
	Update([]int) error
}

type orderService interface {
	PackOrder(n int, dimensions []int) map[int]int
}

type orderController struct {
	dimensionsRepo dimensionsRepository
	orderSrv       orderService
}

func New(repo dimensionsRepository, orderSrv orderService) *orderController {
	return &orderController{dimensionsRepo: repo, orderSrv: orderSrv}
}

func (ctr *orderController) GetOrderPacking(c *gin.Context) {
	// number of items
	nStr := c.Query("items")

	n, err := strconv.Atoi(nStr)
	if err != nil {
		c.JSON(400, gin.H{"message": "invalid items supplied"})

		return
	}

	// retrieve current dimensions
	dimensions, err := ctr.dimensionsRepo.Get()
	if err != nil {
		c.JSON(500, gin.H{"message": "internal server error"})

		return
	}

	if len(dimensions) == 0 {
		c.JSON(200, nil)

		return
	}

	packing := ctr.orderSrv.PackOrder(n, dimensions)
	c.JSON(200, gin.H{
		"packs": packing,
	})
}

func (ctr *orderController) PutOrdersDimensions(c *gin.Context) {
	var body CreateDimensionsRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"message": "bad request"})

		return
	}

	if err := ctr.dimensionsRepo.Update(body.Dimensions); err != nil {
		c.JSON(500, gin.H{"message": "internal server error"})
	}
}

func (ctr *orderController) GetOrdersDimensions(c *gin.Context) {
	// retrieve current dimensions
	dimensions, err := ctr.dimensionsRepo.Get()
	if err != nil {
		c.JSON(500, gin.H{"message": "internal server error"})

		return
	}

	c.JSON(200, gin.H{
		"dimensions": dimensions,
	})
}
