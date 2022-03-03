package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"ucas/xza/controller"
	"ucas/xza/dao/db"
)

//TODO Test资源文件下载
func DownloadFileService(c *gin.Context) {
	fileDir := c.Query("fileDir")
	fileName := c.Query("fileName")
	//打开文件
	_, errByOpenFile := os.Open(fileDir + "/" + fileName)
	_=errByOpenFile
	//非空处理
	//if common.IsEmpty(fileDir) || common.IsEmpty(fileName) || errByOpenFile != nil {
	//	/*c.JSON(http.StatusOK, gin.H{
	//	    "success": false,
	//	    "message": "失败",
	//	    "error":   "资源不存在",
	//	})*/
	//	c.Redirect(http.StatusFound, "/404")
	//	return
	//}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(fileDir + "/" + fileName)
	return
}
func GetFileList(c *gin.Context){
	//x1:= c.PostForm("path")
	//fmt.Println("x1: ",x1)
	//c.JSON(200,gin.H{"123":321})
	windows_base_path := "E:\\就业\\golang\\project\\remote_sense_go\\Files\\DownLoadDir"
	linux_base_path := "/home/ubuntu/xza_file"
	_=windows_base_path
	_=linux_base_path
	files, _ := ioutil.ReadDir(windows_base_path)
	//files, _ := ioutil.ReadDir(windows_base_path)
	//files,_:=ioutil.ReadDir("C:\\Users\\86188\\Desktop\\就业\\golang\\project\\GIN_1\\view")
	res:=make(map[int]string)
	fmt.Println("xxxxxx: ",len(files))

	//res:=make([]string,len(files))
	for  i,file:=range files{
		//res[file.Name()] = "1"
		res[i] = file.Name()
		fmt.Printf(file.Name())
		//res[i] = file.Name()
	}
	c.JSON(200,res)
}
func main() {

	//加载配置文件
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("")
	}

	//数据库
	db.InitDb()


	//创建默认的引擎
	router := gin.Default()
	//告诉gin框架去哪加载模板文件 此处可以使用正则表达式
	//当前是加载view/这个文件夹下的所有文件
	router.LoadHTMLGlob("view/*")

	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	//engine := gin.Default()
	//设置全局跨域访问
	router.Use(AccessJsMiddleware())
	// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
	//router.GET("/user/:name", func(c *gin.Context) {
	//	name := c.Param("name")
	//	c.String(http.StatusOK, "Hello %s", name)
	//})

	// 但是，这个规则既能匹配/user/john/格式也能匹配/user/john/send这种格式
	// 如果没有其他路由器匹配/user/john，它将重定向到/user/john/
	//router.GET("/user/:name/*action", func(c *gin.Context) {
	//	name := c.Param("name")
	//	action := c.Param("action")
	//	message := name + " is " + action
	//	c.String(http.StatusOK, message)
	//})
	//router.POST("/upload", func(c *gin.Context) {
	//	// 单文件
	//	file, _ := c.FormFile("file")
	//	log.Println(file.Filename)
	//
	//	// 上传文件到指定的路径
	//	// c.SaveUploadedFile(file, dst)
	//
	//	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	//})
	router.GET("/downloadFiles",DownloadFileService)
	router.GET("/GetFileList",GetFileList)  //你注意这个路由怎么写的！ 你特么写了个get 你有一个提交的是post啊兄弟
	router.POST("/GetFileList",GetFileList)
	router.GET("/",GetIndex)
		//get请求返回显示页面 index.html
		router.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})

		//创建请求 当访问地址为/uploadfile时执行后面的函数
		//router.POST("/uploadfiles", Uploadfiles)
	router.POST("/uploadfiles",controller.UploadImage)

	router.Run(":8080")
}

func getting(c *gin.Context){
	c.JSON(200,gin.H{"hello": "world"})
}
func GetIndex(c *gin.Context){
	cur, _ := os.Getwd()
	//fmt.Println(filepath.Join(cur, "view/index.html"))
	c.File(filepath.Join(cur, "view/index.html"))
	//c.HTML(http.StatusOK, "index.html",gin.H{})
}

func Uploadfiles(c *gin.Context) {
	//多个文件列表获取
	form,err:=c.MultipartForm()
	if err==nil{
		files:=form.File["f1"]
		for _,f:=range files{
			c.SaveUploadedFile(f, filepath.Join("Files/UploadDir",f.Filename))
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

}
func Uploadfile(c *gin.Context) {
	//获取表单数据 参数为name值
	f, err := c.FormFile("f1")
	////错误处理
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	} else {
		//将文件保存至本项目根目录中
		c.SaveUploadedFile(f, f.Filename)
		//保存成功返回正确的Json数据
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	}

}
//func Cors() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		origin := c.Request.Header.Get("origin")	//请求头部
//		if len(origin) == 0 {
//			origin = c.Request.Header.Get("Origin")
//		}
//		//接收客户端发送的origin （重要！）
//		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
//		//允许客户端传递校验信息比如 cookie (重要)
//		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
//		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
//		//服务器支持的所有跨域请求的方法
//		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE, UPDATE")
//		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
//		// 设置预验请求有效期为 86400 秒
//		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
//		if c.Request.Method == "OPTIONS" {
//			c.AbortWithStatus(204)
//			return
//		}
//		c.Next()
//	}
//}
func AccessJsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := c.Writer
		r:=c.Request
		// 处理js-ajax跨域问题
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Add("Access-Control-Allow-Headers", "Access-Token")
		c.Next()
	}
}
//func CrosHandler() gin.HandlerFunc {
//	return func(context *gin.Context) {
//		method := context.Request.Method
//		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
//		context.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
//		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
//		context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma,token,openid,opentoken")
//		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
//		context.Header("Access-Control-Max-Age", "172800")
//		context.Header("Access-Control-Allow-Credentials", "false")
//		context.Set("content-type", "application/json")//设置返回格式是json
//
//		if method == "OPTIONS" {
//			context.JSON(http.StatusOK, gin.H{"Code": 200, "Data": "Options Request!"})
//		}
//
//		//处理请求
//		context.Next()
//	}
//}
