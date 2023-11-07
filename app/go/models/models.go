package models

type User struct {
	ID     int64
	UserID int `db:"user_id"`
	Name   string
}

// FriendLink user1 and user2 are friends
type FriendLink struct {
	ID      int64
	User1ID int `db:"user1_id"`
	User2ID int `db:"user2_id"`
}

// BlockList user1 blocks user2
type BlockList struct {
	ID      int64
	User1ID int `db:"user1_id"`
	User2ID int `db:"user2_id"`
}
