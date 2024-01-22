package service

import (
	"4096Tiktok/dao"
	"4096Tiktok/middleware"
	"golang.org/x/crypto/bcrypt"
	"log"
	"regexp"
	"strconv"
)

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
	Avatar 			string 	`json:"avatar"`
	BackgroundImage string 	`json:"background_image"`
	Signature 		string 	`json:"signature"`
	TotalFavorited 	string 	`json:"total_favorited"`
	WorkCount 		int 	`json:"work_count"`
	FavoriteCount 	int 	`json:"favorite_count"`
}

func CheckString(string string) bool {
	if ok, _ := regexp.MatchString("^[\\w_-]{6,32}$", string); !ok {
		return false
	}
	return true
}

func VerifyNameAndPwd(username, password string) bool {
	if err := CheckString(username) && CheckString(password); err != true {
		return false
	}
	return true
}

func ReleaseToken(user *dao.User) (string, error){
	token, err := middleware.TokenRelease(*user)
	return token, err
}

func EncryptPwd(password string) string {
	EncryptedPwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(EncryptedPwd)
}

func DecryptPwd(password, encryptedPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPwd), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func AddUser(user *dao.User) error {
	if err := dao.AddUser(user); err != nil {
		log.Println("AddUser failure")
		return err
	}
	return nil
}

func GetUserByName(username string) (dao.User, error){
	if user, err := dao.GetUserByName(username); err != nil {
		log.Println("Get user failure")
		return dao.User{}, err
	}else {
		return user, nil
	}
}

func GetUserById(Id int) (dao.User, error) {
	if user, err := dao.GetUserById(Id); err != nil {
		log.Println("Get user failure")
		return dao.User{}, err
	}else {
		return user, nil
	}
}


func GetUserInfoById(UserId int) User {
	user_db, _ := dao.GetUserById(UserId)
	total_favorited := GetUserLikedCount(UserId)
	work_count := GetVideoCountByUserId(UserId)
	favorite_Count := GetUserLikeCount(UserId)


	Userinfo := User{
		Id:              int64(UserId),
		Name:            user_db.Username,
		Avatar:          avatar,
		BackgroundImage: background_image,
		Signature:       signature,
		TotalFavorited: strconv.FormatInt(total_favorited, 10),
		WorkCount:       int(work_count),
		FavoriteCount:   int(favorite_Count),
	}
	return Userinfo
}
//func Get


