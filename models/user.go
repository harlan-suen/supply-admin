package models

type User struct {
	ID       int64  `gorm:"primary_key", json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     int    `json:"role"`
	OrgID    int64  `json:"orgId"`
}

func AddUser(username, password, phone string, role int, orgID int64) (id int64, err error) {
	var user = User{
		Username: username,
		Password: password,
		Phone:    phone,
		Role:     role,
		OrgID:    orgID,
	}
	err = db.Table("user").Select("Username", "Password", "Phone", "Role", "OrgID").Create(&user).Error
	if err != nil || user.ID == 0 {
		return 0, err
	}
	return user.ID, nil
}
