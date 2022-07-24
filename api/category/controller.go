package category

import (
	"brodo-demo/api/category/request"
	"brodo-demo/api/category/response"
	"brodo-demo/api/common"
	"brodo-demo/service/category"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService *category.CategoryService
}

func NewCategoryController(categoryService *category.CategoryService) *CategoryController {
	return &CategoryController{
		categoryService: categoryService,
	}
}

func (c *CategoryController) PostCategory(ctx *gin.Context) {
	body := request.PostCategoryRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	payload := category.CreateCategoryPayload{
		UserId:       ctx.GetInt("userId"),
		CategoryName: body.Name,
	}

	newCategory, err := c.categoryService.CreateCategory(payload)

	if err != nil {
		statusCode, response := common.NewErrorServiceResponse(err)
		ctx.JSON(statusCode, response)
		return
	}

	response := common.NewSuccessResponse(response.PostCategoryResponse{
		Category: newCategory,
	})

	ctx.JSON(http.StatusCreated, response)
}

func (c *CategoryController) PutCategoryById(ctx *gin.Context) {
	body := request.PostCategoryRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	categoryId, _ := strconv.Atoi(ctx.Param("categoryId"))

	payload := category.UpdateCategoryPayload{
		UserId:       ctx.GetInt("userId"),
		CategoryId:   categoryId,
		CategoryName: body.Name,
	}

	updatedCategory, err := c.categoryService.UpdateCategoryById(payload)

	if err != nil {
		statusCode, response := common.NewErrorServiceResponse(err)
		ctx.JSON(statusCode, response)
		return
	}

	response := common.NewSuccessResponse(response.PostCategoryResponse{
		Category: updatedCategory,
	})

	ctx.JSON(http.StatusCreated, response)
}

func (c *CategoryController) GetCategories(ctx *gin.Context) {
	categories, err := c.categoryService.GetAll()

	if err != nil {
		code, response := common.NewErrorServiceResponse(err)
		ctx.JSON(code, response)
	}

	body := response.GetCategories{
		Categories: categories,
	}

	response := common.NewSuccessResponse(body)

	ctx.JSON(http.StatusOK, response)
}

func (c *CategoryController) DeleteCategoryById(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Param("categoryId"))

	err := c.categoryService.DeleteCategoryById(categoryId, ctx.GetInt("userId"))

	if err != nil {
		statusCode, body := common.NewErrorServiceResponse(err)
		ctx.JSON(statusCode, body)
		return
	}

	response := common.NewSuccessResponse(nil)
	ctx.JSON(http.StatusOK, response)
}