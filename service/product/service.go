package product

import (
	"brodo-demo/entity"
	"brodo-demo/repository"
	"brodo-demo/service/category"
	"brodo-demo/service/pagination"
	"brodo-demo/service/uploader"
	"errors"
	"fmt"
	"mime/multipart"
)

var (
	ErrProductPrice    = errors.New("product price must be greater than Rp1.000")
	ErrProductName     = errors.New("product name length must be greater than 5")
	ErrMainImg         = errors.New("invalid main img")
	ErrProductType     = errors.New("invalid product type")
	ErrProductNotFound = errors.New("product not found")
)

type ProductService struct {
	productRepository  repository.ProductRepository
	categoryRepository repository.CategoryRepository
	uploader           uploader.Uploader
}

type GetAllProductsParam struct {
	Page int
	Size int
	MinPrice int
	MaxPrice int
}

func NewProductService(productRepository repository.ProductRepository, categoryRepository repository.CategoryRepository, uploader uploader.Uploader) *ProductService {
	return &ProductService{
		productRepository:  productRepository,
		categoryRepository: categoryRepository,
		uploader:           uploader,
	}
}

func isValidPrice(price int) bool {
	return price > 1000
}

func isValidProductName(productName string) bool {
	const minimumLengthProductName = 5
	return len(productName) > minimumLengthProductName
}

func (service *ProductService) CreateProduct(product entity.Product, mainImageHeader *multipart.FileHeader, mainImageFile multipart.File) (*entity.Product, error) {
	if !isValidPrice(product.Price) {
		return nil, ErrProductPrice
	}

	if !isValidProductName(product.Name) {
		return nil, ErrProductName
	}

	if _, err := service.categoryRepository.FindById(product.CategoryId); err != nil {
		return nil, category.ErrCategoryNotFound
	}

	mainUrl, err := service.uploader.Upload(mainImageHeader, mainImageFile)

	if err != nil {
		return nil, err
	}

	product.MainImg = mainUrl

	newProduct, err := service.productRepository.InsertOne(product)

	if err != nil {
		return nil, err
	}

	return newProduct, nil
}

func (service *ProductService) GetProducts(params GetAllProductsParam) ([]entity.Product, *pagination.Pagination, error) {
	skip, limit, currentPage := pagination.CreatePagination(params.Page, params.Size)

	products, count, err := service.productRepository.FindAllAndCount(repository.FindAllProductsParam{
		Skip:  skip,
		Limit: limit,
		MinPrice: params.MinPrice,
		MaxPrice: params.MaxPrice,
	})

	if err != nil {
		return []entity.Product{}, nil, err
	}

	return products, pagination.NewPagination(count, currentPage, limit), nil
}

func (service *ProductService) GetProductById(productId int) (entity.Product, error) {
	product, err := service.productRepository.FindById(productId)

	if err != nil {
		fmt.Println(err)
		return product, ErrProductNotFound
	}

	return product, nil
}
