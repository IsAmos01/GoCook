package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	MAX_UPLOAD_SIZE = 1024 * 1024 * 20 // 50MB
)

func Upload(ctx *gin.Context) {
	ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, MAX_UPLOAD_SIZE)
	
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	contentType := header.Header["Content-Type"][0]
	if contentType != "image/jpeg" {
		return
	}
	
	data, err := io.ReadAll(file)
	if err != nil {
		return
	}
	
	err = os.WriteFile("./hhh"+header.Filename, data, 0777)
	if err != nil {
		return
	}
	
	ctx.Writer.WriteHeader(http.StatusCreated)
	_, _ = io.WriteString(ctx.Writer, "successfully uploaded")
}
