package models

import (
	"github.com/cockroachdb/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserRole string

var (
	UserRoleManager  = UserRole("manager")
	UserRoleEmployee = UserRole("employee")
	UserRoleMember   = UserRole("member")
)

type UserStatus string

var (
	UserStatusActive  = UserStatus("active")
	UserStatusPending = UserStatus("pending")
	UserStatusBanned  = UserStatus("banned")
)

type User struct {
	Model

	UID      string     `gorm:"type:character varying(13);not null;unique"`
	Phone    string     `gorm:"type:character varying(11);not null;unique"`
	Password string     `gorm:"type:character varying(50);not null"`
	Role     UserRole   `gorm:"type:character varying(10);not null"`
	Status   UserStatus `gorm:"type:character varying(10);not null"`
	Cart     []int64    `gorm:"type:bigint[];not null"`
	// Point: point reward when user done a order
}

func (u *User) TableName() string {
	return "users"
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.Wrap(err, "generate from password")
	}

	return string(bytes), nil
}

func (u *User) Authenticate(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
