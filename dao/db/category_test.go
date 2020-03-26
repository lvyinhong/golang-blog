package db

import (
	"testing"
)

func init() {
	//parseTime=true 将mysql中时间类型，自动解析为go结构体中的时间类型
	dsn := "root:3Ydc6lbdnf.c1@tcp(localhost:3306)/go_blog?parseTime=true"
	err := Init(dsn)
	if err!=nil{
		panic(err)
	}
}

func TestGetCategoryById(t *testing.T) {
	category, err := GetCategoryById(1)
	if err != nil {
		panic(err)
	}
	t.Logf("category:%#v", category)
}

func TestGetCategoryList(t *testing.T) {
	var categoryIds []int64

	categoryIds = append(categoryIds, 1,2,3)
	categoryList, err := GetCategoryList(categoryIds)
	if err != nil {
		panic(err)
	}

	for _,v := range categoryList {
		t.Logf("id:%v category:%#v\n", v.CategoryId, v.CategoryName)
	}

}

func TestGetAllCategoryList(t *testing.T) {
	categoryList, err := GetAllCategoryList()
	if err != nil {
		panic(err)
	}
	for _,v := range categoryList {
		t.Logf("id:%v category:%#v\n", v.CategoryId, v.CategoryName)
	}

}