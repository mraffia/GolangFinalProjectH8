package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	Message string `json:"message" form:"message" valid:"required~Comment message is required"`
	UserID  uint
	User    *User
	PhotoID uint
	Photo   *Photo
}

func (p *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}

func (p *Comment) BeforeDelete(tx *gorm.DB) (err error) {
	_, errDelete := govalidator.ValidateStruct(p)

	if errDelete != nil {
		err = errDelete
		return
	}

	err = nil
	return
}
