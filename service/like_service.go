package service

import "4096Tiktok/dao"

func GetVideoById(VideoId int) (dao.Video, error) {
	var video dao.Video
	if video, err := dao.GetVideoById(VideoId); err == nil && video.UserID != 0 {
		return video, err
	}
	return video, nil
}

func FavorVideo(UserId int, action int, video *dao.Video) error {
	var err error
	switch action {
		case 1:
			err = dao.LikeVideo(video, UserId)
		case 2:
			err = dao.DislikeVideo(video, UserId)
	}
	return err
}

func GetUserLikeCount(Id int) int64 {
	count := dao.UserLikeCount(Id)
	return count
}

func GetUserLikedCount(Id int) int64 {
	count := dao.GetUserLikedCount(Id)
	return count
}

func GetVideoCountByUserId(Id int) int64 {
	count := dao.GetVideoCountByUserId(Id)
	return count
}

func GetUserLikeVideos(Id int) []Video {
	videos := dao.GetUserLikeVideos(Id)
	videoInfos := GetVideosInfo(videos, Id)
	return videoInfos
}