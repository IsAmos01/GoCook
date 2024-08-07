package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx/v3"
	"mime/multipart"
	"strings"
)

func Upload(ctx *gin.Context) {
	// 方式一：ctx.Request.FormFile("file") 获取file类型以及header
	file, header, err := ctx.Request.FormFile("file")
	
	// 方式二：ctx.FormFile("file") 获取 fileHeader，再通过open方法获取file类型
	// formFile, err := ctx.FormFile("file")
	// open, err := formFile.Open()
	// if err != nil {
	// 	return
	// }
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	
	// 判断文件类型
	filename := header.Filename
	fileExt := strings.ToLower(filename[strings.LastIndex(filename, "."):])
	if fileExt != ".xls" && fileExt != ".xlsx" {
		return
	}
	
	xlFile, err := xlsx.OpenReaderAt(file, header.Size)
	if err != nil {
		return
	}
	// 获取sheet0数据
	sheet0 := xlFile.Sheets[0]
	// 遍历方式一
	err = sheet0.ForEachRow(func(r *xlsx.Row) error {
		err = r.ForEachCell(func(c *xlsx.Cell) error {
			fmt.Println(c.Value)
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}
	
	// 遍历方式二:转为三维数组
	// output, err := xlFile.ToSliceUnmerged() // output  [sheet][row][cell]string
	// for i := 0; i < len(output); i++ {
	// 	for j := 0; j < len(output[i]); j++ {
	// 		// 操作数据
	// 	}
	// }
}
