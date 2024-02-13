package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `json:"name" form:"name" valid:"required~Name of social media is required"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url" valid:"required~Social media url is required"`
	UserID         uint
	User           *User
}

func (p *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}

func (p *SocialMedia) BeforeDelete(tx *gorm.DB) (err error) {
	_, errDelete := govalidator.ValidateStruct(p)

	if errDelete != nil {
		err = errDelete
		return
	}

	err = nil
	return
}
