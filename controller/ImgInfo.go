package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"path/filepath"
	"time"
	"ucas/xza/dao/db"
	"ucas/xza/model"
)

func UploadImage(c *gin.Context) {

	//多个文件列表获取
	form,err:=c.MultipartForm()
	if err==nil{
		files:=form.File["f1"]
		for _,f:=range files{
			c.SaveUploadedFile(f, filepath.Join("Files/UploadDir",f.Filename))
			//band := c.PostForm("band")
			//thingType := c.PostForm("author")
			band := "3"
			//thingType := "plane"

			ImageInfo := &model.ImageInfo{
				Band: band,
				//ThingType: thingType,
				//Angle
				//Uuid: uuid.New().String(),
				Uuid: uuid.NewString(),
				CreateTime: time.Now(),
			}
			flag, _ := db.InsertImage(ImageInfo)
			if !flag {
				//c.HTML(http.StatusInternalServerError, "view/500.html", nil)
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err,
				})
				return
			}
		}
		//将文件保存至本项目根目录中
		//保存成功返回正确的Json数据
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	}else{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}


	//band := c.PostForm("band")
	//thingType := c.PostForm("author")
	//ImageInfo := &model.ImageInfo{
	//	Band: band,
	//	ThingType: thingType,
	//	//Angle
	//	//Uuid: uuid.New().String(),
	//	Uuid: uuid.NewString(),
	//	CreateTime: time.Now(),
	//}
	//flag, _ := db.InsertImage(ImageInfo)
	//if !flag {
	//	c.HTML(http.StatusInternalServerError, "views/500.html", nil)
	//}
}
