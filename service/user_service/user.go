package user_service

import "supply-admin/models"

type User struct {
	ID       int64  `gorm:"primary_key", json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     int    `json:"role"`
	OrgID    int64    `json:"org_id"`
}

func (u *User) Add() (int64, error) {
	return models.AddUser(u.Username, u.Password, u.Phone, u.Role, int64(u.OrgID))
}
