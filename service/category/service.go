package category

import (
	"brodo-demo/entity"
	"brodo-demo/repository"
	"brodo-demo/service/errservice"
	"errors"
)

var (
	ErrCategoryNameTooShort = errors.New("category name too shorts")
	ErrCategoryNotFound     = errors.New("category not found")
)

type CategoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

type CreateCategoryPayload struct {
	UserId       int
	CategoryName string
}

type UpdateCategoryPayload struct {
	UserId       int
	CategoryName string
	CategoryId   int
}

func isValidCategoryName(name string) bool {
	const minPasswordLength = 3
	return len(name) > minPasswordLength
}

func (service *CategoryService) CreateCategory(payload CreateCategoryPayload) (*entity.Category, error) {
	if !isValidCategoryName(payload.CategoryName) {
		return nil, ErrCategoryNameTooShort
	}

	newCategory := entity.Category{
		Name:      payload.CategoryName,
		CreatedBy: payload.UserId,
	}

	createdCategory, err := service.categoryRepository.InsertOne(newCategory)

	if err != nil {
		return nil, err
	}

	return createdCategory, nil
}

func (service *CategoryService) UpdateCategoryById(payload UpdateCategoryPayload) (*entity.Category, error) {
	if !isValidCategoryName(payload.CategoryName) {
		return nil, ErrCategoryNameTooShort
	}

	foundCategory, err := service.categoryRepository.FindById(payload.CategoryId)

	if err != nil {
		return nil, ErrCategoryNotFound
	}

	if foundCategory.CreatedBy != payload.UserId {
		return nil, errservice.ErrForbidden
	}

	foundCategory.Name = payload.CategoryName

	updatedCategory, err := service.categoryRepository.UpdateById(foundCategory)

	if err != nil {
		return nil, err
	}
	return updatedCategory, nil
}

func (service *CategoryService) GetAll() ([]entity.Category, error) {
	categories, err := service.categoryRepository.FindAll()

	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (service *CategoryService) DeleteCategoryById(categoryId int, userId int) error {
	category, err := service.categoryRepository.FindById(categoryId)

	if err != nil {
		return ErrCategoryNotFound
	}

	if category.CreatedBy != userId {
		return errservice.ErrForbidden
	}

	err = service.categoryRepository.DeleteById(categoryId)

	if err != nil {
		return err
	}

	return nil
}
