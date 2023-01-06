package utils

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func WriteLogToFile(filename string) {
	f, _ := os.Create(filename)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
