package main

type User struct {
	ID    string  `json:"id" gorm:"primaryKey;column:id"`
	Name  string  `json:"name" gorm:"column:name"`
	Roles []*Role `json:"roles" gorm:"many2many:UserRole;foreignKey:id;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Role struct {
	ID          string  `json:"id" gorm:"primaryKey;column:id"`
	Description string  `json:"description" gorm:"column:description"`
	Users       []*User `json:"roles" gorm:"many2many:UserRole;foreignKey:id;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type UserRole struct {
	UserID string `json:"UserID" gorm:"primaryKey;column:UserID"`
	RoleID string `json:"RoleID" gorm:"primaryKey;column:RoleID"`
}
