package db

import (
	"go_logger/model"
	"testing"
	"time"
)

func init() {
	//parseTime=true 将mysql中时间类型，自动解析为go结构体中的时间类型
	dsn := "root:3Ydc6lbdnf.c1@tcp(localhost:3306)/go_blog?parseTime=true"
	err := Init(dsn)
	if err!=nil{
		panic(err)
	}
}

// 测试插入文章
func TestInsertArticle(t *testing.T) {
	// 构建对象
	article := &model.ArticleDetail{}
	article.ArticleInfo.CategoryId = 1
	article.ArticleInfo.CommentCount=0
	article.Content = "abcdefageweqw adweqsa"
	article.ArticleInfo.CreateTime = time.Now()
	article.ArticleInfo.Title="测试用例"
	article.ArticleInfo.Username="summd"
	article.ArticleInfo.Summary="abcfd"
	article.ArticleInfo.ViewCount=1
	id, err := InsertArticle(article)
	if err != nil {
		return
	}
	t.Logf("articleId:%#v\n", id)
}

func TestGetArticleList(t *testing.T) {
	articleList, err := GetArticleList(1, 2)
	if err != nil {
		t.Errorf(`GetArticleList(1, 2) failed %v`, err)
		return
	}
	t.Logf("article : %d\n", len(articleList))
}

func TestGetArticleDetail(t *testing.T) {
	article, err := GetArticleDetail(1)
	if err != nil {
		t.Errorf(`GetArticleDetail(1) failed %v`, err)
		return
	}
	t.Logf("article: %#v\n", article)
}

func TestGetArticleListByCategoryId(t *testing.T) {
	articleList, err := GetArticleListByCategoryId(1,1,4)
	if err != nil {
		t.Errorf(`GetArticleListByCategoryId(1,1,4) failed %v`, err)
		return
	}
	t.Logf("article : %d\n", len(articleList))

}