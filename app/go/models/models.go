package models

type User struct {
	ID      uint64  `gorm:"autoIncrement"`
	UserID  int     `gorm:"not null"`
	Name    string  `gorm:"not null;size:64;default:''"`
	Friends []*User `gorm:"many2many:friend_link;joinForeignKey:User1ID;joinReferences:User2ID"`
}

type FriendLink struct {
	ID      uint64 `gorm:"autoIncrement"`
	User1ID int    `gorm:"not null"`
	User2ID int    `gorm:"not null"`
}

type BlockList struct {
	ID      uint64 `gorm:"autoIncrement"`
	User1ID int    `gorm:"not null"`
	User2ID int    `gorm:"not null"`
}
