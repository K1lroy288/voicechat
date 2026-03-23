package service

import (
	"user-service/model"
	"user-service/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetFriends(userID uint) ([]model.User, error) {
	return s.repo.GetFriends(userID)
}

func (s *UserService) GetFriendsRequest(userID int) ([]model.FriendRequest, error) {
	return s.repo.GetFriendsRequest(userID)
}

func (s *UserService) AddFriend(usernameOrID string, userID int) error {
	return s.repo.AddFriend(usernameOrID, userID)
}

func (s *UserService) FriendRequestResponse(friendshipID int, res string) error {
	return s.repo.FriendRequestResponse(friendshipID, res)
}
