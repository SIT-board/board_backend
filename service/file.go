package service

import (
	"fmt"
	"mime/multipart"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FileService struct {
	Directory string
}

func (service *FileService) SaveFile(ctx *gin.Context, fileHeader *multipart.FileHeader) (filename string, err error) {
	filename = fmt.Sprintf("%s_%s_%s", time.Now().Format("2006_01_02"), uuid.New(), fileHeader.Filename)
	prefixFilename := fmt.Sprintf("%s/%s", service.Directory, filename)
	if err = ctx.SaveUploadedFile(fileHeader, prefixFilename); err != nil {
		return "", err
	}
	return filename, nil

}
