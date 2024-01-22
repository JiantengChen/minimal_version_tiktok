package dao

type Comment struct {
	CommentID uint `gorm:"primarykey"`
	UserId int
	VideoId int
	Text string
	CreatedAt int
}
