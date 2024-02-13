package router

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/photos")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreatePhoto)
		productRouter.GET("/", controllers.GetAllPhotos)
		productRouter.GET("/:photoId", controllers.GetPhoto)
		productRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		productRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.GET("/", controllers.GetAllComments)
		commentRouter.GET("/:commentId", controllers.GetComment)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	socialmediaRouter := r.Group("/socialmedias")
	{
		socialmediaRouter.Use(middlewares.Authentication())
		socialmediaRouter.POST("/", controllers.CreateSocialMedia)
		socialmediaRouter.GET("/", controllers.GetAllSocialMedias)
		socialmediaRouter.GET("/:socialmediaId", controllers.GetSocialMedia)
		socialmediaRouter.PUT("/:socialmediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialmediaRouter.DELETE("/:socialmediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	return r
}
