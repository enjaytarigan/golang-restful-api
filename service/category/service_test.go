package category

import (
	"brodo-demo/entity"
	repository "brodo-demo/repository/mocks"
	"brodo-demo/service/errservice"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsValidCategoryName_Failed(t *testing.T) {
	tests := []struct {
		testName string
		input    string
		want     bool
	}{
		{
			testName: "should return false when given empty string",
			input:    "",
			want:     false,
		},
		{
			testName: "should return false when given categoryName length less than 3",
			input:    "HP",
			want:     false,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			isValid := isValidCategoryName(test.input)

			assert.Equal(t, test.want, isValid)
		})
	}
}

func TestIsValidCategoryName_Success(t *testing.T) {
	isValid := isValidCategoryName("gadget")

	assert.True(t, isValid)
}

func TestCreateCategory(t *testing.T) {
	t.Run("should return ErrCategoryNameTooShort when given categoryName with length less than 3", func(t *testing.T) {
		payload := CreateCategoryPayload{
			UserId:       1,
			CategoryName: "hp",
		}

		service := CategoryService{
			categoryRepository: nil,
		}

		newCategory, err := service.CreateCategory(payload)

		assert.Nil(t, newCategory)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, ErrCategoryNameTooShort)
	})

	t.Run("should return category entity correctly", func(t *testing.T) {
		mockCategoryRepository := new(repository.CategoryRepository)

		payload := CreateCategoryPayload{
			UserId:       1,
			CategoryName: "gadget",
		}

		category := entity.Category{
			CreatedBy: payload.UserId,
			Name:      payload.CategoryName,
		}

		mockCategoryRepository.On("InsertOne", category).Return(&entity.Category{}, nil)

		categoryService := CategoryService{
			categoryRepository: mockCategoryRepository,
		}

		newCategory, err := categoryService.CreateCategory(payload)

		assert.NotNil(t, newCategory)
		assert.Nil(t, err)
	})
}

func TestUpdateCategoryById(t *testing.T) {
	t.Run("should return ErrCategoryNameTooShort when given categoryName with length less than 3", func(t *testing.T) {
		payload := UpdateCategoryPayload{
			UserId:       1,
			CategoryName: "hp",
			CategoryId:   1,
		}

		service := CategoryService{
			categoryRepository: nil,
		}

		newCategory, err := service.UpdateCategoryById(payload)

		assert.Nil(t, newCategory)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, ErrCategoryNameTooShort)
	})

	t.Run("should return ErrCategoryNotFound when given invalid categoryId", func(t *testing.T) {
		payload := UpdateCategoryPayload{
			UserId:       1,
			CategoryName: "gadget",
			CategoryId:   0,
		}

		mockCategoryRepository := new(repository.CategoryRepository)

		mockCategoryRepository.On("FindById", payload.CategoryId).Return(entity.Category{}, errors.New("category not found"))

		categoryService := CategoryService{
			categoryRepository: mockCategoryRepository,
		}

		updatedCategory, err := categoryService.UpdateCategoryById(payload)

		assert.Nil(t, updatedCategory)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, ErrCategoryNotFound)
	})

	t.Run("should return errservice.ErrForbideen when payload.UserId is not match to foundCategory.CreatedBy", func(t *testing.T) {
		payload := UpdateCategoryPayload{
			UserId:       1,
			CategoryName: "gadget",
			CategoryId:   1,
		}

		mockCategoryRepository := new(repository.CategoryRepository)

		foundCategory := entity.Category{
			ID:        payload.CategoryId,
			CreatedBy: 22,
		}

		mockCategoryRepository.On("FindById", payload.CategoryId).Return(foundCategory, nil)

		categoryService := CategoryService{
			categoryRepository: mockCategoryRepository,
		}

		updatedCategory, err := categoryService.UpdateCategoryById(payload)

		assert.Nil(t, updatedCategory)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, errservice.ErrForbidden)
	})

	t.Run("should return updated category correclty", func(t *testing.T) {
		category := entity.Category{
			ID:        1,
			CreatedBy: 1,
			Name:      "Old Name",
			CreatedAt: time.Now(),
		}

		payload := UpdateCategoryPayload{
			UserId:       1,
			CategoryName: "Fashion",
		}

		mockCategoryRepository := new(repository.CategoryRepository)

		expectedUpdatedResult := &entity.Category{
			ID:        category.ID,
			Name:      payload.CategoryName,
			CreatedAt: category.CreatedAt,
			CreatedBy: category.CreatedBy,
		}

		category.Name = payload.CategoryName

		mockCategoryRepository.On("FindById", payload.CategoryId).Return(category, nil)
		mockCategoryRepository.On("UpdateById", category).Return(expectedUpdatedResult, nil)

		service := CategoryService{
			categoryRepository: mockCategoryRepository,
		}

		updatedCategory, err := service.UpdateCategoryById(payload)

		assert.Nil(t, err)
		assert.NotNil(t, updatedCategory)
		assert.Equal(t, payload.CategoryName, updatedCategory.Name)
	})
}

func TestGetAllCategories(t *testing.T) {
	t.Run("should return categories data", func(t *testing.T) {
		mockCategoryRepository := new(repository.CategoryRepository)

		categories := []entity.Category{
			{ID: 1, Name: "Sport", CreatedAt: time.Now(), CreatedBy: 1},
			{ID: 2, Name: "Fashion", CreatedAt: time.Now(), CreatedBy: 2},
		}

		mockCategoryRepository.On("FindAll").Return(categories, nil)

		categoryService := CategoryService{
			categoryRepository: mockCategoryRepository,
		}

		foundCategories, err := categoryService.GetAll()

		assert.Nil(t, err)
		assert.Equal(t, len(categories), len(foundCategories))
	})
}

func TestDeleteCategoryById(t *testing.T) {
	t.Run("should return ErrCategoryNotFound when given invalid categoryId", func(t *testing.T) {
		mockCategoryRepository := new(repository.CategoryRepository)

		categoryId := 1

		mockCategoryRepository.On("FindById", categoryId).Return(entity.Category{}, errors.New("category not found"))

		categoryService := CategoryService{
			categoryRepository: mockCategoryRepository,
		}

		err := categoryService.DeleteCategoryById(categoryId, 1)

		assert.NotNil(t, err)
		assert.ErrorIs(t, err, ErrCategoryNotFound)
	})

	t.Run("should return errservice.ErrForbidden when given wrong userId", func(t *testing.T) {
		mockCategoryRepository := new(repository.CategoryRepository)

		categoryId := 1

		category := entity.Category{
			ID:        1,
			Name:      "Example",
			CreatedAt: time.Now(),
			CreatedBy: 1,
		}

		mockCategoryRepository.On("FindById", categoryId).Return(category, nil)

		categoryService := CategoryService{
			categoryRepository: mockCategoryRepository,
		}

		err := categoryService.DeleteCategoryById(categoryId, 0)

		assert.NotNil(t, err)
		assert.ErrorIs(t, err, errservice.ErrForbidden)
	})

	t.Run("should not return error when given valid categoryId and userId", func(t *testing.T) {
		mockCategoryRepository := new(repository.CategoryRepository)

		categoryId := 1
		userId := 1

		category := entity.Category{
			ID:        1,
			Name:      "Example",
			CreatedAt: time.Now(),
			CreatedBy: userId,
		}

		mockCategoryRepository.On("FindById", categoryId).Return(category, nil)
		mockCategoryRepository.On("DeleteById", category.ID).Return(nil)

		categoryService := CategoryService{
			categoryRepository: mockCategoryRepository,
		}

		err := categoryService.DeleteCategoryById(categoryId, userId)

		assert.Nil(t, err)
	})
}
