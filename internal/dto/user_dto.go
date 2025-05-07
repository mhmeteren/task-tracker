package dto

import "task-tracker/internal/model"

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (r *CreateUserRequest) ToModel() model.User {
	return model.User{
		Name:  r.Name,
		Email: r.Email,
	}
}

type UserListItem struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

func ToUserListItem(u model.User) UserListItem {
	return UserListItem{
		ID:        u.ID,
		Name:      u.Name,
		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToUserList(u []model.User) []UserListItem {
	var list []UserListItem
	for _, user := range u {
		list = append(list, ToUserListItem(user))
	}
	return list
}

type UserDetail struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	UserKey   string `json:"user_key"`
	CreatedAt string `json:"created_at"`
}

func ToUserDetail(u model.User) UserDetail {
	return UserDetail{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		UserKey:   u.UserKey,
		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
