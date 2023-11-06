package models

type User struct {
	ID      uint64  `gorm:"autoIncrement"`
	UserID  int     `gorm:"not null"`
	Name    string  `gorm:"not null;size:64;default:''"`
	Friends []*User `gorm:"many2many:friend_link;joinForeignKey:User1ID;joinReferences:User2ID"`
	Blocks  []*User `gorm:"many2many:block_list;joinForeignKey:User1ID;joinReferences:User2ID"`
}

// FriendLink user1 and user2 are friends
type FriendLink struct {
	ID      uint64 `gorm:"autoIncrement"`
	User1ID int    `gorm:"not null"`
	User2ID int    `gorm:"not null"`
}

// BlockList user1 blocks user2
type BlockList struct {
	ID      uint64 `gorm:"autoIncrement"`
	User1ID int    `gorm:"not null"`
	User2ID int    `gorm:"not null"`
}
