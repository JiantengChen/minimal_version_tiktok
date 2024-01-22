package main

import (
	"4096Tiktok/controller"
	"4096Tiktok/middleware"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", middleware.JwtMiddleWarePass(), controller.Feed)
	apiRouter.GET("/user/", middleware.JwtMiddleWare(), controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action/", middleware.JwtMiddleWare(), controller.Publish)
	apiRouter.GET("/publish/list/",  middleware.JwtMiddleWare(), controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", middleware.JwtMiddleWare(), controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", middleware.JwtMiddleWare(), controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	apiRouter.GET("/relation/friend/list/", controller.FriendList)
	apiRouter.GET("/message/chat/", controller.MessageChat)
	apiRouter.POST("/message/action/", controller.MessageAction)
}
