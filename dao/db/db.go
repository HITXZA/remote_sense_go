package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"net/url"
	"os"
)

var (
	DB  *gorm.DB
	err error
)
//当InitDb没有被main函数使用的时候   这个gorm这个东西也是没有被引入工程的
//这个时候你去执行god mod tidy -v  没有效果 因为这个文件没被主函数直接或者间接的使用
func InitDb() error {

	//dsn := "root:123.@tcp(127.0.0.1:3306)/blogs?parseTime=true&charset=utf8mb4&loc=Local"

	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err!=nil{
		return err
	}
	//viper.GetString("")

	driverName := "mysql"

	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))
	fmt.Printf(args)
	DB, err := gorm.Open(driverName, args)
	//if err != nil {
	//	panic("fail to connect database, err: " + err.Error())
	//}
	//if DB, err = gorm.Open(driverName, dsn); err != nil {
	//	return err //此error只是判断上述参数是否正确，并不代表真的连接成功
	//}
	if err = DB.DB().Ping(); err != nil {
		return err
	}
	DB.DB().SetMaxIdleConns(16)
	DB.DB().SetMaxOpenConns(100)
	//DB.AutoMigrate(&model.ArticleDetail{})
	//DB.AutoMigrate(&model.Comment{})
	//defer DB.Close()
	return nil
}
