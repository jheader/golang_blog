package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jheader/golang_blog/controller"
	"github.com/jheader/golang_blog/middleware"
)

func SetupRoutes() *gin.Engine {

	r := gin.New()
	r.Use(middleware.LoggerMiddleWare())
	r.Use(middleware.ErrorHandleMiddleWare())
	r.Use(gin.Recovery())

	api := r.Group("/api/v1")
	{
		// 认证相关路由（无需认证）
		auth := api.Group("/auth")
		{
			auth.POST("/register", (&controller.AuthController{}).Register)
			auth.POST("/login", (&controller.AuthController{}).Login)
		}
		// 需要认证的路由
		authenticated := api.Group("")
		authenticated.Use(middleware.AuthMiddleware())
		{
			authenticated.GET("/profile", (&controller.User{}).GetProfile)
			//文章
			postsRout := authenticated.Group("/posts")
			{ //发表或者更新
				postsRout.POST("/saveOrUpdate", (&controller.PostController{}).CreateOrUpdate)
				//删除
				postsRout.DELETE("/:post_id", (&controller.PostController{}).DeletedById)

			}
			//评论授权路由 实现评论的创建功能，已认证的用户可以对文章发表评论。
			addcomment := authenticated.Group("/posts/:post_id/comment")
			{
				addcomment.POST("", (&controller.CommentController{}).CreateComment)

			}

		}

		// 公开路由（无需认证）
		public := api.Group("")
		{
			// 文章公开路由获取所有文章列表和
			public.GET("/posts", (&controller.PostController{}).GetAllPosts)
			//单个文章的详细信息
			public.GET("/posts/:post_id", (&controller.PostController{}).GetPostById)
			//获取某篇文章的所有评论列表。
			public.GET("/posts/commentsByPostId", func(ctx *gin.Context) {})
		}

		// 健康检查
		r.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "Blog API is running",
			})
		})

	}

	return r
}
