package dao

import (
	"log"
)

type User struct {
	UserID 		uint			`gorm:"primarykey"`
	Username 	string 		`gorm:"index:,unique"`
	Password 	string
	Videos		[] Video
	Comments	[] Comment
	Likes		[] Video 	`gorm:"many2many:like;joinForeignKey:user_id;joinReferences:video_id;"`
	Fans		[] User 	`gorm:"many2many:follow;joinForeignKey:user_id;joinReferences:fan_id;"`
}

func AddUser(user *User) error{
	DB := GetDB()
	tx := DB.Begin()
	if err := tx.Model(&User{}).Create(&user).Error; err != nil {
		tx.Rollback()
		log.Println(err.Error())
		return err
	}
	tx.Commit()
	return nil
}

func GetUserByName(name string) (User, error) {
	DB := GetDB()
	tx := DB.Begin()
	user := User{}
	if err := tx.Model(&User{}).Where("username = ?", name).First(&user).Error; err != nil {
		tx.Rollback()
		log.Println(err.Error())
		return user, err
	}
	tx.Commit()
	return user, nil
}

func GetUserById(Id int) (User, error) {
	DB := GetDB()
	tx := DB.Begin()
	user := User{}
	if err := tx.Where("user_id = ?", Id).First(&user).Error; err != nil {
		tx.Rollback()
		log.Println(err.Error())
		return user, err
	}
	tx.Commit()
	return user, nil
}

func GetFollowCount(fanId int) (int64, error) {
	//var count int64
	//DB := GetDB()
	//tx := DB.Begin()
	//if err := tx.Model(&Follow{}).Where("fan_id = ?", fanId).Count(&count).Error; err != nil {
	//	tx.Rollback()
	//	return 0, err
	//}
	//tx.Commit()
	//return count, nil
	return 0, nil
}

func GetFanCount(userId int64) (int64, error) {
	//var count int64
	//DB := GetDB()
	//tx := DB.Begin()
	//tx.Model(&User{}).Association("Fans").Count()
	//tx.Commit()
	//return count, nil
	return 0, nil
}

func IsFollow(fanId, userId int64) (bool, error) {
	return false, nil
}












