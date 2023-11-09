package repository

import (
	"github.com/jmoiron/sqlx"
	"problem1/models"
	"problem1/services"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) userExists(id int64) (bool, error) {
	var exist bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)"
	query = u.db.Rebind(query)
	err := u.db.Get(&exist, query, id)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (u *UserRepository) FindFriendsByID(id int64) ([]*models.User, error) {
	userExists, err := u.userExists(id)
	if err != nil {
		return nil, err
	}
	if !userExists {
		return nil, services.ErrUserNotFound
	}

	var friends []*models.User
	query := `
		SELECT u.id, u.user_id, u.name
		FROM friend_link AS fl
		JOIN users AS u
			ON (fl.user1_id = :id OR fl.user2_id = :id) AND
			    u.id = IF(fl.user1_id = :id, fl.user2_id, fl.user1_id)
		ORDER BY u.id
	`
	query, args, err := sqlx.Named(query, models.User{ID: id})
	if err != nil {
		return nil, err
	}
	query = u.db.Rebind(query)
	err = u.db.Select(&friends, query, args...)
	if err != nil {
		return nil, err
	}
	return friends, nil
}

func (u *UserRepository) FindBlockUsersByID(id int64) ([]*models.User, error) {
	userExists, err := u.userExists(id)
	if err != nil {
		return nil, err
	}
	if !userExists {
		return nil, services.ErrUserNotFound
	}

	var blocks []*models.User
	query := `
		SELECT u.id, u.user_id, u.name 
		FROM users AS u
		JOIN block_list AS bl
			ON bl.user1_id = :id AND bl.user2_id = u.id
		ORDER BY u.id
	`
	query, args, err := sqlx.Named(query, models.User{ID: id})
	if err != nil {
		return nil, err
	}
	query = u.db.Rebind(query)
	err = u.db.Select(&blocks, query, args...)
	if err != nil {
		return nil, err
	}
	return blocks, nil
}

// FindFriendsOfFriendsByID does not exclude blocked users etc.
func (u *UserRepository) FindFriendsOfFriendsByID(id int64) ([]*models.User, error) {
	userExists, err := u.userExists(id)
	if err != nil {
		return nil, err
	}
	if !userExists {
		return nil, services.ErrUserNotFound
	}

	var friendsOfFriends []*models.User
	if err != nil {
		return nil, err
	}
	query := `
		SELECT u2.id, u2.user_id, u2.name
		FROM friend_link AS fl1
		JOIN users AS u1
			ON (fl1.user1_id = :id OR fl1.user2_id = :id) AND
			   u1.id = IF(fl1.user1_id = :id, fl1.user2_id, fl1.user1_id)
		JOIN friend_link AS fl2
		    ON fl2.user1_id != :id AND fl2.user2_id != :id AND
		       (fl2.user1_id = u1.id OR fl2.user2_id = u1.id)
		JOIN users AS u2
			ON u2.id != :id AND
			   (fl2.user1_id = u1.id OR fl2.user2_id = u1.id) AND
			   u2.id = IF(fl2.user1_id = u1.id, fl2.user2_id, fl2.user1_id)
		ORDER BY u2.id
	`
	query, args, err := sqlx.Named(query, models.User{ID: id})
	if err != nil {
		return nil, err
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}
	query = u.db.Rebind(query)
	err = u.db.Select(&friendsOfFriends, query, args...)
	if err != nil {
		return nil, err
	}
	return friendsOfFriends, nil
}

func (u *UserRepository) FindFriendsOfFriendsExcludingSomeUsersByIDWithPagination(
	id int64, excludeIDs []int64, page int, limit int,
) ([]*models.User, error) {
	userExists, err := u.userExists(id)
	if err != nil {
		return nil, err
	}
	if !userExists {
		return nil, services.ErrUserNotFound
	}

	var result []*models.User
	if err != nil {
		return nil, err
	}
	arg := map[string]interface{}{
		"id":          id,
		"exclude_ids": excludeIDs,
		"limit":       limit,
		"offset":      (page - 1) * limit,
	}
	query := `
		SELECT u2.id, u2.user_id, u2.name
		FROM friend_link AS fl1
		JOIN users AS u1
			ON (fl1.user1_id = :id OR fl1.user2_id = :id) AND
			   u1.id = IF(fl1.user1_id = :id, fl1.user2_id, fl1.user1_id)
		JOIN friend_link AS fl2
		    ON fl2.user1_id != :id AND fl2.user2_id != :id AND
		       (fl2.user1_id = u1.id OR fl2.user2_id = u1.id)
		JOIN users AS u2
			ON u2.id != :id AND
			   u2.id NOT IN (:exclude_ids) AND
			   (fl2.user1_id = u1.id OR fl2.user2_id = u1.id) AND
			   u2.id = IF(fl2.user1_id = u1.id, fl2.user2_id, fl2.user1_id)
		ORDER BY u2.id
		LIMIT :limit OFFSET :offset
	`
	query, args, err := sqlx.Named(query, arg)
	if err != nil {
		return nil, err
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}
	query = u.db.Rebind(query)
	err = u.db.Select(&result, query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}
