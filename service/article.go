package service

import (
	"go_logger/dao/db"
	"go_logger/model"
)

// 获取文章和对应的分类
func GetArticleRecordList(pageNum, pageSize int)(articleRecordList []*model.ArticleRecord, err error) {
	// 1、获取文章列表
	articleInfoList, err := db.GetArticleList(pageNum, pageSize)
	if err != nil || len(articleInfoList) <=0 {
		return
	}
	// 2、获取文章对应的分类(多个)
	categoryIds := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		return
	}

	// 返回页面，做聚合
	for _, article := range articleInfoList {
		// 根据当前的文章生成结构体
		articleRecord := &model.ArticleRecord{
			ArticleInfo:*article,
		}
		for _,category := range categoryList {
			if category.CategoryId == article.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

// 根据多个文章的id，获取多个分类id的集合
func getCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64){
	// 变量文章，得到每一个文章
	for _, article := range articleInfoList {
		categoryHas := false	//已经有这个id了
		// 从当前文章获取分类id
		categoryId := article.CategoryId

		// 去重，防止重复
		for _,id := range ids {
			//看当前id是否存在
			if id == categoryId {
				categoryHas = true
			}
		}
		if !categoryHas {
			ids = append(ids, categoryId)
		}
	}
	return
}

// 根据分类id, 获取该文章和他们对应的分类信息
func GetArticleRecordListByCategoryId(categoryId int64, pageNum, pageSize int)(articleRecordList []*model.ArticleRecord, err error) {
	articleInfoList, err := db.GetArticleListByCategoryId(categoryId,pageNum, pageSize)
	if err != nil || len(articleInfoList) <=0 {
		return
	}
	// 2、获取文章对应的分类(多个)
	categoryIds := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		return
	}

	// 返回页面，做聚合
	for _, article := range articleInfoList {
		// 根据当前的文章生成结构体
		articleRecord := &model.ArticleRecord{
			ArticleInfo:*article,
		}
		for _,category := range categoryList {
			if category.CategoryId == article.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}