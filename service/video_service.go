package service

import (
	"4096Tiktok/dao"
	//"fmt"
	"log"
)

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

func AddVideo(video *dao.Video) error {
	if err := dao.AddVideo(video); err != nil {
		log.Println("AddVideo failure")
		return err
	}
	return nil
}

func GetVideoByUserIDAndTitle(UID uint, title string) bool {
	if video, err := dao.GetVideoByUserAndTitle(UID, title); err == nil && video.UserID != 0{
		return true
	}
	return false
}

func Get30Videos(wawtcher_id int) []Video {
	videos := dao.Get30Videos()
	videoInfos := GetVideosInfo(videos, wawtcher_id)
	return videoInfos
}

func GetVideosInfo(videos [] dao.Video, watcher_id int) []Video {
	videoinfos := []Video{}
	for _, v :=  range videos {
		user := GetUserInfoById(int(v.UserID))
		videoID := int(v.VideoID)
		var isFavorited = false
		if watcher_id != 0 {
			isFavorited = dao.IsVideoFavorited(videoID, watcher_id)
		}
		favoriteCount := dao.VideoLikeCount(videoID)
		videoInfo := Video{
			Id:            int64(videoID),
			Author:        user,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: favoriteCount,
			CommentCount:  int64(len(v.Comments)),
			IsFavorite:    isFavorited,
		}
		videoinfos = append(videoinfos, videoInfo)
	}
	return videoinfos
}

func GetUserVideosByID(watcherID, userID int) []Video {
	videos := dao.GetUserVideosByID(userID)
	videoInfos := GetVideosInfo(videos, watcherID)
	return videoInfos
}