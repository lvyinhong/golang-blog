package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go_logger/model"
)

// 插入文章
func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {
	// 加个验证
	if article == nil {
		return
	}
	sqlstr := `insert into article(comment, summary, title, username, category_id, view_count, comment_count) values(?,?,?,?,?,?,?)`
	result, err := DB.Exec(sqlstr,
		article.Content,article.Summary,article.Title,article.Username,
		article.ArticleInfo.CategoryId,article.ViewCount,article.CommentCount)
	if err != nil {
		return
	}
	articleId, err = result.LastInsertId()
	return
}

// 查询文章列表，分页
func GetArticleList(pageNum,pageSize int) (ArticleList []*model.ArticleInfo, err error){
	if pageNum < 0 || pageSize <=0 {
		return
	}
	// 时间降序排序
	sqlstr := "select id, summary, title, view_count,create_time,comment_count,username,category_id " +
		"from article" +
		"where status = 1 order by create_time desc limit ?,?"
	err = DB.Select(&ArticleList,sqlstr,pageNum,pageSize)
	return
}

// 根据文章id，查询单个文章
func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	if articleId < 0 {
		err = fmt.Errorf("invalid parameter,article_id:%d", articleId)
		return
	}
	articleDetail = &model.ArticleDetail{}
	sqlstr := `select 
							id, summary, title, view_count, content,
							 create_time, comment_count, username, category_id
						from 
							article 
						where 
							id = ?
						and
							status = 1
						`
	err = DB.Get(articleDetail, sqlstr, articleId)
	return
}

// 根据分类id，查询这一类的文章
func GetArticleListByCategoryId(categoryId int64, pageNum, pageSize int)(articleList []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize <=0 {
		return
	}
	// 时间降序排序
	sqlstr := "select id, summary, title, view_count,create_time,comment_count,username,category_id " +
		"from " +
		"article " +
		"where status = 1 and category_id = ? order by create_time desc limit ?,?"
	err = DB.Select(&articleList,sqlstr,categoryId,pageNum,pageSize)
	return
}
