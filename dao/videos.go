package dao

import (
	"log"
)

type Video struct {
	VideoID		uint 		`gorm:"primarykey"`
	UserID 		uint
	PlayUrl 	string		`gorm:"not null"`
	CoverUrl	string		`gorm:"not null"`
	Title 		string		`gorm:"not null"`
	Comments	[] Comment
	Likes		[] User	`gorm:"many2many:like;joinForeignKey:video_id;joinReferences:user_id;"`
}

func AddVideo(video *Video) error {
	DB := GetDB()
	tx := DB.Begin()
	if err := tx.Model(&Video{}).Create(&video).Error; err != nil {
		tx.Rollback()
		log.Println(err.Error())
		return err
	}
	tx.Commit()
	return nil
}

func GetVideoByUserAndTitle(UID uint, title string) (Video, error) {
	DB := GetDB()
	tx := DB.Begin()
	video := Video{}
	if err := tx.Where("user_id = ? AND title = ?", UID, title).First(&video).Error; err != nil {
		tx.Rollback()
		log.Println(err.Error())
		return video, err
	}
	tx.Commit()
	return video, nil
}

func GetVideoIdsByUserId(UserId int) [] uint {
	DB := GetDB()
	tx := DB.Begin()
	video_ids := []uint{}
	tx.Model(&Video{}).Where("user_id = ?", UserId).Select("video_id").Find(&video_ids)
	tx.Commit()
	return video_ids
}

func GetVideoCountByUserId(UserId int) int64 {
	DB := GetDB()
	tx := DB.Begin()
	var count int64
	tx.Model(&Video{}).Where("user_id = ?", UserId).Count(&count)
	tx.Commit()
	return count
}

func Get30Videos() []Video{
	DB := GetDB()
	tx := DB.Begin()
	var videos []Video
	tx.Order("video_id desc").Limit(30).Find(&videos)
	tx.Commit()
	return videos
}

func GetUserVideosByID(userID int) []Video {
	var user User
	DB := GetDB()
	tx := DB.Begin()
	tx.Preload("Videos").First(&user, userID)
	videos := user.Videos
	tx.Commit()
	return videos
}

