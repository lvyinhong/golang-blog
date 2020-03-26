package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go_logger/controller"
	"go_logger/dao/db"
	"os"
)

// 初始化加载配置文件
func initConfig()() {
	workDir,_:= os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func main() {
	initConfig()
	err :=db.Init(
		viper.GetString("database.driver"),
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
			viper.GetString("database.host"),
			viper.GetInt("database.port"),
			viper.GetString("database.dbname"),
			viper.GetString("database.charset")))
	if err!=nil {
		panic(fmt.Errorf("connect mysql error %v\n", err))
		return
	}

	r := gin.Default()

	// 加载静态文件
	r.Static("/static/", "./static")

	//加载模板
	r.LoadHTMLGlob("views/*")

	r.GET("/", controller.IndexHandle)
	r.GET("/category/", controller.CategoryList)

	_:r.Run("localhost:8080")
	

}