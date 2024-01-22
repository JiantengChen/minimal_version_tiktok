package dao

import (
	"fmt"
	"log"
)

func GetVideoById(VideoId int) (Video, error) {
	DB := GetDB()
	tx := DB.Begin()
	var video Video
	if err := tx.Model(&Video{}).Where("video_id = ?", VideoId).First(&video).Error; err != nil {
		tx.Rollback()
		log.Println(err.Error())
		return video, err
	}

	tx.Commit()
	return video, nil
}

func LikeVideo(video *Video, UserId int) error {
	DB := GetDB()
	tx := DB.Begin()
	if err := tx.Model(&video).Association("Likes").Append(&User{UserID: uint(UserId)}); err != nil {
		tx.Rollback()
		log.Println(err.Error())
		return err
	}
	tx.Commit()
	return nil
}

func DislikeVideo(video *Video, UserId int) error {
	DB := GetDB()
	tx := DB.Begin()
	if err := tx.Model(&video).Association("Likes").Delete(&User{UserID: uint(UserId)}); err != nil {
		tx.Rollback()
		log.Println()
		return err
	}
	tx.Commit()
	return nil
}

func IsVideoFavorited(videoID, userID int) bool{
	DB := GetDB()
	tx := DB.Begin()
	var user User
	tx.Preload("Likes").First(&user, userID)
	var isLiked bool
	for _, video := range user.Likes {
		if video.VideoID == uint(videoID) {
			isLiked = true
			break
		}
	}
	tx.Commit()
	return isLiked
}

func VideoLikeCount(VideoId int) int64 {
	DB := GetDB()
	tx := DB.Begin()
	var video Video
	tx.Preload("Likes").First(&video, VideoId)
	likesCount := len(video.Likes)
	fmt.Println("video ", VideoId, " count is ", likesCount)
	tx.Commit()
	return int64(likesCount)
}

func UserLikeCount(UserId int) int64 {
	DB := GetDB()
	tx := DB.Begin()
	var user User
	tx.Preload("Likes").First(&user, UserId)
	var likesCount int64
	for _, video := range user.Likes {
		likesCount += tx.Model(&video).Association("Likes").Count()
	}
	tx.Commit()
	return likesCount
}

func GetUserLikedCount(Id int) int64 {
	DB := GetDB()
	tx := DB.Begin()

	var user User
	tx.Preload("Videos.Likes").Find(&user, Id)

	var totalLikes int64
	for _, video := range user.Videos {
		totalLikes += int64(len(video.Likes))
	}
	tx.Commit()
	return totalLikes
}

func GetUserLikeVideos(Id int) []Video {
	DB := GetDB()
	tx := DB.Begin()
	var user User
	tx.Preload("Likes").First(&user, Id)
	videos := user.Likes
	tx.Commit()
	return videos
}