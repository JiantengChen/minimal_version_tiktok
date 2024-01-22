package controller

import (
	"4096Tiktok/dao"
	"4096Tiktok/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")

	vid, _ := strconv.Atoi(videoId)
	action, _ := strconv.Atoi(actionType)
	User, _ := c.Get("user")
	user := User.(dao.User)

	// check whether video exists
	var video dao.Video
	var err error
	if video, err = service.GetVideoById(vid); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 401, StatusMsg: "video doesn't exist",
		})
		return
	}

	if err := service.FavorVideo(int(user.UserID), action, &video); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 402, StatusMsg: "action failed",
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0, StatusMsg: "action success",
	})
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	userId := c.Query("user_id")
	id, _ := strconv.Atoi(userId)
	_, err := service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusOK, UserInfoFailResponse{
			Response: Response{StatusCode: 205, StatusMsg: "user doesn't exist"},
			Userinfo: nil,
		})
		return
	}
	Videos := service.GetUserLikeVideos(id)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: Videos,
	})
}
