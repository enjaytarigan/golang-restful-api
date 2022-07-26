package product

import (
	"brodo-demo/entity"
	repository "brodo-demo/repository/mocks"
	"brodo-demo/service/category"
	uploader "brodo-demo/service/uploader/mocks"
	"errors"
	"mime/multipart"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsValidPrice(t *testing.T) {
	tableTest := []struct {
		name  string
		price int
		want  bool
	}{
		{
			name:  "should return false when given price with negative value",
			price: -1,
			want:  false,
		},
		{
			name:  "should return false when given price with 0 value",
			price: 0,
			want:  false,
		},
		{
			name:  "should return false when given price less than 1000",
			price: 0,
			want:  false,
		},
		{
			name:  "should return true when given price greater than 1000",
			price: 150_000,
			want:  true,
		},
	}

	for _, test := range tableTest {
		t.Run(test.name, func(t *testing.T) {
			isValid := isValidPrice(test.price)

			assert.Equal(t, test.want, isValid)
		})
	}
}

func TestIsValidProductName(t *testing.T) {
	testTable := []struct {
		name        string
		productName string
		want        bool
	}{
		{
			name:        "should return false when given empty productName",
			productName: "",
			want:        false,
		},
		{
			name:        "should return false when given productName with length less than 5",
			productName: "jojo",
			want:        false,
		},
		{
			name:        "should return true when given valid productName",
			productName: "Alpha Legacy Vegtan",
			want:        true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			isValid := isValidProductName(test.productName)

			assert.Equal(t, test.want, isValid)
		})
	}
}

func TestCreateProduc(t *testing.T) {
	t.Run("should return ErrProductPrice when given invalid price", func(t *testing.T) {
		service := ProductService{
			productRepository:  nil,
			categoryRepository: nil,
		}

		product, err := service.CreateProduct(entity.Product{Price: 999}, nil, nil)

		assert.Nil(t, product)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, ErrProductPrice)
	})

	t.Run("should return ErrProductName when given invalid name of product", func(t *testing.T) {
		service := ProductService{
			productRepository:  nil,
			categoryRepository: nil,
		}

		product, err := service.CreateProduct(entity.Product{Price: 150000, Name: "h"}, nil, nil)

		assert.Nil(t, product)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, ErrProductName)
	})

	t.Run("should return ErrCategoryNotFound when given invalid categoryId", func(t *testing.T) {
		categoryRepository := new(repository.CategoryRepository)

		product := entity.Product{
			Name:       "Alpha Legacy Vegtan",
			Price:      150_000,
			MainImg:    "http://example.com/main.jpg",
			CategoryId: 0,
		}

		categoryRepository.On("FindById", product.CategoryId).Return(entity.Category{}, errors.New("category not found"))

		service := ProductService{
			productRepository:  nil,
			categoryRepository: categoryRepository,
		}

		newProduct, err := service.CreateProduct(product, nil, nil)

		assert.Nil(t, newProduct)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, category.ErrCategoryNotFound)
	})

	t.Run("should return new product correctly", func(t *testing.T) {
		categoryRepository := new(repository.CategoryRepository)
		productRepository := new(repository.ProductRepository)
		mockUploader := new(uploader.Uploader)

		product := entity.Product{
			Name:        "Alpha Legacy Vegtan",
			Price:       150_000,
			CategoryId:  1,
			CreatedBy:   1,
			Description: "Hello World Description",
		}

		createdProduct := &entity.Product{
			ID:          1,
			Name:        "Alpha Legacy Vegtan",
			Price:       150_000,
			MainImg:     "main_img_url",
			CategoryId:  1,
			CreatedBy:   1,
			Description: "Hello World Description",
			CreatedAt:   time.Now(),
		}

		fileHeader := &multipart.FileHeader{}

		categoryRepository.On("FindById", product.CategoryId).Return(entity.Category{}, nil)
		mockUploader.On("Upload", fileHeader, nil).Return("main_img_url", nil)

		product.MainImg = "main_img_url"

		productRepository.On("InsertOne", product).Return(createdProduct, nil)

		service := ProductService{
			productRepository:  productRepository,
			categoryRepository: categoryRepository,
			uploader:           mockUploader,
		}

		newProduct, err := service.CreateProduct(product, fileHeader, nil)

		assert.NotNil(t, newProduct)
		assert.Nil(t, err)
		assert.Equal(t, product.Name, newProduct.Name)
		assert.Equal(t, product.Price, newProduct.Price)
		assert.Equal(t, product.Description, newProduct.Description)
		assert.Equal(t, product.CategoryId, newProduct.CategoryId)
		assert.Equal(t, product.CreatedBy, newProduct.CreatedBy)
	})
}
