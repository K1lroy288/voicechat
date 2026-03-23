package model

import "gorm.io/gorm"

type FriendRequest struct {
	gorm.Model
	SenderID   uint   `gorm:"uniqueIndex:idx_sender_receiver"`
	ReceiverID uint   `gorm:"uniqueIndex:idx_sender_receiver"`
	Status     string `gorm:"default:pending"`

	Sender   User `gorm:"foreignKey:SenderID"`
	Receiver User `gorm:"foreignKey:ReceiverID"`
}
