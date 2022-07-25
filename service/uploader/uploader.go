package uploader

import "mime/multipart"

type Uploader interface {
	Upload(fileHeader *multipart.FileHeader, file multipart.File) (string, error)
}