package service

import (
	"fmt"
	"mime/multipart"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ImageService struct {
	RawDirectory       string
	ThumbnailDirectory string
}

type fileName struct {
	prefixPath string
	format     string
	time       time.Time
	uuid       uuid.UUID
}

func (r *fileName) String() string {
	return fmt.Sprintf("%s/%s_%s.%s", r.prefixPath, r.time.Format("2006_01_02"), r.uuid, r.format)
}

func (r *ImageService) SaveImage(c *gin.Context, fileHeader *multipart.FileHeader) (string, error) {
	splited := strings.Split(fileHeader.Filename, ".")
	fn := fileName{
		prefixPath: r.RawDirectory,
		format:     splited[len(splited)-1],
		time:       time.Now(),
		uuid:       uuid.New(),
	}
	filename := fn.String()
	err := c.SaveUploadedFile(fileHeader, filename)
	if err != nil {
		return "", err
	}
	return filename, err
}
