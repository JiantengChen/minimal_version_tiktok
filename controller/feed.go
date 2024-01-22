package controller

import (
	"4096Tiktok/dao"
	"4096Tiktok/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []service.Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	var watcher_id int
	User, err := c.Get("user")
	if err == false {
		watcher_id = 0
	}else {
		user := User.(dao.User)
		watcher_id = int(user.UserID)
	}

	videos := service.Get30Videos(watcher_id)
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  time.Now().Unix(),
	})
}
