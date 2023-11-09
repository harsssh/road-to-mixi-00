package handlers

import "problem1/models"

type FriendListEntry struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func convertToFriendList(users []*models.User) []*FriendListEntry {
	friends := make([]*FriendListEntry, len(users))
	for i, user := range users {
		friends[i] = &FriendListEntry{
			ID:   user.ID,
			Name: user.Name,
		}
	}
	return friends
}
