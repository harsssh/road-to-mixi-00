package models

type User struct {
    ID     uint64 `gorm:"column:id;primaryKey;autoIncrement"`
	UserID int    `gorm:"column:user_id;not null"`
	Name   string `gorm:"column:name;not null;size:64;default:''"`
}

type FriendLink struct {
    ID      uint64 `gorm:"primaryKey;autoIncrement;column:id"`
    User1ID int    `gorm:"column:user1_id;not null"`
    User2ID int    `gorm:"column:user2_id;not null"`
}

type BlockList struct {
    ID      uint64 `gorm:"primaryKey;autoIncrement;column:id"`
    User1ID int    `gorm:"column:user1_id;not null"`
    User2ID int    `gorm:"column:user2_id;not null"`
}
