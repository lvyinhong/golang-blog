package controller

import (
	"github.com/gin-gonic/gin"
	"go_logger/service"
	"net/http"
	"strconv"
)

// 访问主页的控制器
func IndexHandle(c *gin.Context) {
	// 从service取数据
	articleRecordList, err := service.GetArticleRecordList(0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "view/500.html", nil)
		return
	}
	// 加载分类数据
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "view/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "view/index.html", gin.H{
		"article_list": articleRecordList,
		"category_list": categoryList,
	})
}

// 点击分类云，进行分类
func CategoryList(c *gin.Context) {
	categoryIdStr := c.Query("category_id")
	// 转成int
	categoryId,err:= strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "view/500.html", nil)
		return
	}

	// 根据分类id， 获取文章列表
	articleRecordList, err := service.GetArticleRecordListByCategoryId(categoryId, 0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "view/500.html", nil)
		return
	}

	// 再次加载所有分类数据，用于分类云显示
	categoryList, err := service.GetAllCategoryList()
	c.HTML(http.StatusOK, "view/index.html", gin.H{
		"article_list": articleRecordList,
		"category_list": categoryList,
	})
}