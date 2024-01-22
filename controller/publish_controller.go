package controller

import (
	"4096Tiktok/dao"
	"4096Tiktok/ossDB"
	"4096Tiktok/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
)

type VideoListResponse struct {
	Response
	VideoList []service.Video `json:"video_list"`
}


// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	title := c.Query("title")
	fileHeader, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 301,
			StatusMsg:  err.Error(),
		})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 302,
			StatusMsg:  err.Error(),
		})
		return
	}

	User, _ := c.Get("user")
	user := User.(dao.User)
	filename := filepath.Base(fileHeader.Filename)

	if service.GetVideoByUserIDAndTitle(user.UserID, title) == true {
		c.JSON(http.StatusOK, Response{
			StatusCode: 303,
			StatusMsg:  "duplicate video name",
		})
		return
	}

	playName := fmt.Sprintf("video/%d_%s_%s", user.UserID, title, filename)

	if err := oss.PutObject(playName, file); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 304,
			StatusMsg:  err.Error(),
		})
		return
	}

	playUrl := oss.GeneratePlayUrl(playName)
	coverUrl := oss.GenerateCoverUrl(playUrl)

	video := dao.Video{
		UserID: user.UserID,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		Title:    title,
	}
	if err = service.AddVideo(&video); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 305,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  playName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	userId := c.Query("user_id")
	id, _ := strconv.Atoi(userId)
	User, _ := c.Get("user")
	user := User.(dao.User)
	_, err := service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusOK, UserInfoFailResponse{
			Response: Response{StatusCode: 205, StatusMsg: "user doesn't exist"},
			Userinfo: nil,
		})
		return
	}
	videos := service.GetUserVideosByID(int(user.UserID), id)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}

