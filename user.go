package model

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type User struct {
	Base      `valid:"required"`
	Name      string    `gorm:"column:owner_name;type:varchar(255);not null" valid:"notnull"`
	UserID 	  string    `gorm:"column:bank_id;type:uuid;not null" valid:"notnull"`
	Email     string    `json:"number" gorm:"type:varchar(50)" valid:"notnull"`
	Status    string    `json:"status" gorm:"type:varchar(20)" valid:"notnull"`
}

func (u *User) isValid() error {
	_, err := govalidator.ValidateStruct(u)

	if len(u.Email) <= 0 || u.Email == " " {
		return errors.New("email required")
	}

	if len(u.Name) <= 0  || u.Name == " "{
		return errors.New("name required")
	}

	if u.Status != "active" && u.Status != "inactive" {
		return errors.New("invalid status")
	}

	if err != nil {
		return err
	}
	return nil
}

func NewUser(name string, email string) (*User, error) {
	user := User{
		Name:      name,
		Email:     email,
		Status:    "active",
	}
	user.UserID = uuid.NewV4().String()
	err := user.isValid()
	if err != nil {
		return nil, err
	}
	return &user, nil
}