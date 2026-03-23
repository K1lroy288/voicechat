package repository

import (
	"user-service/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *model.User) (bool, error) {
	var exist model.User
	res := r.DB.Where("username = ?", user.Username).First(&exist).Error
	if res != nil {
		err := r.DB.Create(user).Error
		return false, err
	}
	return res == nil, nil
}

func (r *UserRepository) GetUserByUsername(username string) (model.User, error) {
	var user model.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *UserRepository) GetUserById(userID int) (model.User, error) {
	var user model.User
	err := r.DB.Where("id = ?", userID).First(&user).Error
	return user, err
}

func (r *UserRepository) GetFriends(userID uint) ([]model.User, error) {
	var friends []model.User

	err := r.DB.Raw(`
        SELECT u.* FROM users u
        JOIN friend_requests fr ON (
            (fr.sender_id = ? AND fr.receiver_id = u.id) OR
            (fr.receiver_id = ? AND fr.sender_id = u.id)
        )
        WHERE fr.status = 'accepted'
    `, userID, userID).Scan(&friends).Error

	return friends, err
}

func (r *UserRepository) GetFriendsRequest(userID int) ([]model.FriendRequest, error) {
	var requests []model.FriendRequest
	err := r.DB.Preload("Sender").Where("received_id = ? AND status = ?", userID, "pending").Error
	return requests, err
}

func (r *UserRepository) AddFriend(usernameOrID string, userID int) error {
	var exist model.User
	err := r.DB.Where("username = ? or id = ?", usernameOrID, usernameOrID).First(&exist).Error
	if err != nil {
		return err
	}

	return r.DB.Create(&model.FriendRequest{SenderID: uint(userID), ReceiverID: exist.ID}).Error
}

func (r *UserRepository) FriendRequestResponse(friendshipID int, res string) error {
	return r.DB.Where("id = ?", friendshipID).Model(&model.FriendRequest{}).Update("status", res).Error
}
