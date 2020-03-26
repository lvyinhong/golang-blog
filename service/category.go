package service

import (
	"go_logger/dao/db"
	"go_logger/model"
)

// 获取所有分类
func GetAllCategoryList()(categoryList []*model.Category, err error) {
	categoryList, err = db.GetAllCategoryList()
	return
}