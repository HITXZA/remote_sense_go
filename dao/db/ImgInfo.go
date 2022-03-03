package db

import (
	model2 "ucas/xza/model"
)

//插入图片基本信息
func InsertImage(imageInfo *model2.ImageInfo) (bool, error) {
	if err := DB.Table("ImageInfo").Create(imageInfo).Error; err != nil {
		return false, err
	}
	return true, nil
}
