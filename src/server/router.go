package server

import (
	"github.com/gin-gonic/gin"
	"go-blog/src/api"
)

// NewRouter 路由配置
// 关于路由地址的几点想法
//	1. 路由地址写法
//   	a. 有的喜欢使用公共前缀然后拼接上自己的部分路由,有的使用 全路径
//  	   比如 index := v1.Group("/api/v1/index"){ index.GET("carousel", api.QueryCarousel)}
//         比如 index := v1.Group(""){ index.GET("/api/v1/index/carousel", api.QueryCarousel)}
//		   Java 里面也有类似的，比如我们 在 Controller 的上面有时候会标注 公共路由 @RequestMapping("api/v1")
// 		b. 这两种各有优劣，第一种层次比较清晰简洁，但是不利于搜索；第二种排查问题方便，直接搜索路由就行；
// 		   我们在项目中leader 比较推荐第二种，这样团队搜索路由很方面，能很快定位。因为这边我也全部改成第二种了
func NewRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/login", api.Login)
	//r.GET("/verify/:token", api.Verify)
	//r.GET("/refresh/:token", api.Refresh)
	//r.GET("/sayHello/:token", api.SayHello)
	return r
}
