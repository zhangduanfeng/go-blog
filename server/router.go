package server

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go-blog/errno"
	"go-blog/handler"
	"go-blog/store"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	// 可自定义盐值
	TokenSalt = "default_salt"
)

func MD5(data []byte) string {
	_md5 := md5.New()
	_md5.Write(data)
	return hex.EncodeToString(_md5.Sum([]byte("")))
}

// NewRouter 路由配置
// 关于路由地址的几点想法
//  1. 路由地址写法
//     a. 有的喜欢使用公共前缀然后拼接上自己的部分路由,有的使用 全路径
//     比如 index := v1.Group("/api/v1/index"){ index.GET("carousel", api.QueryCarousel)}
//     比如 index := v1.Group(""){ index.GET("/api/v1/index/carousel", api.QueryCarousel)}
//     Java 里面也有类似的，比如我们 在 Controller 的上面有时候会标注 公共路由 @RequestMapping("api/v1")
//     b. 这两种各有优劣，第一种层次比较清晰简洁，但是不利于搜索；第二种排查问题方便，直接搜索路由就行；
//     我们在项目中leader 比较推荐第二种，这样团队搜索路由很方面，能很快定位。因为这边我也全部改成第二种了
func NewRouter() *gin.Engine {
	// 禁用控制台颜色
	gin.DisableConsoleColor()
	// 创建记录日志的文件
	path := "blog"
	writer, _ := rotatelogs.New(
		path+"%Y%m%d.log",
		rotatelogs.WithRotationCount(7),
		rotatelogs.WithRotationSize(100*1024*1024),
		//这里设置1分钟产生一个日志文件
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)
	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	gin.DefaultWriter = io.MultiWriter(writer, os.Stdout)

	r := gin.Default()
	r.POST("/login", handler.Login)
	r.POST("/register", handler.Register)

	//以下的接口，都使用Authorize()中间件身份验证
	r.Use(Authorize())

	r.GET("/article/list", handler.ListArticles)
	r.GET("/article/details", handler.ArticleDetails)
	r.POST("/article/create", handler.CreateArticle)
	r.POST("/article/publish", handler.PublishArticle)
	r.POST("/article/save", handler.SaveArticle)
	r.POST("/upload", handler.Upload)
	r.GET("/category/list", handler.ListCategoryAndTag)
	return r
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token") // 访问令牌
		if token != "" {
			result, err := store.RedisClient.Get(store.Ctx, token).Result()
			if err != nil || result == "" {
				// 验证不通过，不再调用后续的函数处理
				c.Abort()
				c.JSON(http.StatusUnauthorized, errno.ConstructErrResp("1", "token认证失败"))
			}
			c.Next()
		} else {
			// 验证不通过，不再调用后续的函数处理
			c.Abort()
			c.JSON(http.StatusUnauthorized, errno.ConstructErrResp("1", "请求头不存在token"))
		}
	}
}
