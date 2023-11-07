package repository

import (
	"database/sql"
	"errors"
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

func (u *UserRepository) findUserByUserID(userID int) (*models.User, error) {
	var user models.User
	query := "SELECT id, user_id, name FROM users WHERE user_id = ?"
	query = u.db.Rebind(query)
	err := u.db.Get(&user, query, userID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) FindFriendsByUserID(userID int) ([]*models.User, error) {
	user, err := u.findUserByUserID(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, services.ErrUserNotFound
		}
		return nil, err
	}

	var friends []*models.User
	query := `
		SELECT u.id, u.user_id, u.name
		FROM users AS u 
		WHERE u.id IN (
			SELECT IF(fl.user1_id = :id, fl.user2_id, fl.user1_id)
			FROM friend_link AS fl
			WHERE fl.user1_id = :id OR fl.user2_id = :id
		)
		ORDER BY u.id
	`
	query, args, err := sqlx.Named(query, user)
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

func (u *UserRepository) FindBlockUsersByUserID(userID int) ([]*models.User, error) {
	user, err := u.findUserByUserID(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, services.ErrUserNotFound
		}
		return nil, err
	}

	var blocks []*models.User
	query := `
		SELECT u.id, u.user_id, u.name 
		FROM users AS u
		WHERE u.id IN (
		    SELECT bl.user2_id 
		    FROM block_list AS bl
		    WHERE user1_id = :id
		)
		ORDER BY u.id
	`
	query, args, err := sqlx.Named(query, user)
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

func (u *UserRepository) findFriendIDsByID(id int64) ([]int64, error) {
	var friendIDs []int64
	query := `
		SELECT IF(fl.user1_id = :id, fl.user2_id, fl.user1_id)
		FROM friend_link AS fl
		WHERE fl.user1_id = :id OR fl.user2_id = :id
	`
	query, args, err := sqlx.Named(query, models.User{ID: id})
	if err != nil {
		return nil, err
	}
	query = u.db.Rebind(query)
	err = u.db.Select(&friendIDs, query, args...)
	if err != nil {
		return nil, err
	}
	return friendIDs, nil
}

func (u *UserRepository) findBlockIDsByID(id int64) ([]int64, error) {
	var blockIDs []int64
	query := "SELECT user2_id FROM block_list WHERE user1_id = ?"
	query = u.db.Rebind(query)
	err := u.db.Select(&blockIDs, query, id)
	if err != nil {
		return nil, err
	}
	return blockIDs, nil
}

// FindFriendsOfFriendsByUserID does not exclude blocked users etc.
func (u *UserRepository) FindFriendsOfFriendsByUserID(userID int) ([]*models.User, error) {
	user, err := u.findUserByUserID(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, services.ErrUserNotFound
		}
		return nil, err
	}

	var friendsOfFriends []*models.User
	friendIDs, err := u.findFriendIDsByID(user.ID)
	if err != nil {
		return nil, err
	}
	arg := map[string]interface{}{
		"id":         user.ID,
		"friend_ids": friendIDs,
	}
	query := `
		SELECT u.id, u.user_id, u.name
		FROM users AS u
		WHERE
		    u.id != :id AND
		    u.id IN (
				SELECT fl.user2_id
				FROM friend_link AS fl
				WHERE fl.user1_id IN (:friend_ids)
				UNION 
				SELECT fl.user1_id
				FROM friend_link AS fl
				WHERE fl.user2_id IN (:friend_ids)
			)
		ORDER BY u.id
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
	err = u.db.Select(&friendsOfFriends, query, args...)
	if err != nil {
		return nil, err
	}
	return friendsOfFriends, nil
}

func (u *UserRepository) FindFriendsOfFriendsExcludingSomeUsersByUserIDWithPagination(
	userID int, excludedUserIDs []int, page int, limit int,
) ([]*models.User, error) {
	user, err := u.findUserByUserID(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, services.ErrUserNotFound
		}
		return nil, err
	}

	var result []*models.User
	friendIDs, err := u.findFriendIDsByID(user.ID)
	if err != nil {
		return nil, err
	}
	arg := map[string]interface{}{
		"id":               user.ID,
		"friend_ids":       friendIDs,
		"exclude_user_ids": excludedUserIDs,
		"limit":            limit,
		"offset":           (page - 1) * limit,
	}
	query := `
		SELECT u.id, u.user_id, u.name
		FROM users AS u
		WHERE
		    u.id != :id AND
		    u.id NOT IN (:exclude_user_ids) AND
		    u.id IN (
		        SELECT fl.user2_id
		        FROM friend_link AS fl
		        WHERE fl.user1_id IN (:friend_ids)
		        UNION
		        SELECT fl.user1_id
		        FROM friend_link AS fl
		        WHERE fl.user2_id IN (:friend_ids)
		    )
		ORDER BY u.id
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
