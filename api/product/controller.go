package product

import (
	"brodo-demo/api/common"
	"brodo-demo/api/product/request"
	"brodo-demo/api/product/response"
	"brodo-demo/entity"
	"brodo-demo/service/product"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService *product.ProductService
}

func NewProductController(productService *product.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (c *ProductController) PostProduct(ctx *gin.Context) {
	body := request.PostProductRequest{}

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return

	}

	product := entity.Product{
		Name:        body.Name,
		Description: body.Description,
		CategoryId:  body.CategoryId,
		Price:       body.Price,
		CreatedBy:   ctx.GetInt("userId"),
	}

	mainImgHeader, err := ctx.FormFile("mainImg")

	if err != nil {
		return
	}

	mainImgFile, err := mainImgHeader.Open()

	if err != nil {
		return
	}

	newProduct, err := c.productService.CreateProduct(product, mainImgHeader, mainImgFile)

	if err != nil {
		fmt.Println(err)
		statusCode, response := common.NewErrorServiceResponse(err)
		ctx.JSON(statusCode, response)
		return
	}

	ctx.JSON(http.StatusCreated, common.NewSuccessResponse(response.ProductResponse{Product: newProduct}))
}

func (c *ProductController) GetProducts(ctx *gin.Context) {
	size, _ := strconv.Atoi(ctx.Query("size"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	minPrice, _ := strconv.Atoi(ctx.Query("minPrice"))
	maxPrice, _ := strconv.Atoi(ctx.Query("maxPrice"))

	param := product.GetAllProductsParam{
		Page: page,
		Size: size,
		MinPrice: minPrice,
		MaxPrice: maxPrice,
	}

	products, pagination, err := c.productService.GetProducts(param)

	if err != nil {
		statusCode, response := common.NewErrorServiceResponse(err)
		ctx.JSON(statusCode, response)
		return
	}

	dataResponse := response.GetProductsResponse{
		Products:   products,
		Pagination: pagination,
	}

	ctx.JSON(http.StatusOK, common.NewSuccessResponse(&dataResponse))
}

func (c *ProductController) GetProductById(ctx *gin.Context) {
	productId, _ := strconv.Atoi(ctx.Param("productId"))

	product, err := c.productService.GetProductById(productId)

	if err != nil {
		statusCode, body := common.NewErrorServiceResponse(err)
		ctx.JSON(statusCode, body)
		return
	}

	response := response.ProductResponse{
		Product: &product,
	}

	ctx.JSON(http.StatusOK, common.NewSuccessResponse(&response))
}
