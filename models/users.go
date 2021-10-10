package models

type User struct {
	ID       string `json:"id"`
	User     string `json:"user"`
	Password string `json:"password"`
	Rol      int    `json:"rol_id"`
	State    int    `json:"user_state_id"`
	Register int    `json:"people_id"`
}

type RegisterNewUserRequest struct {
	User     string `json:"user"`
	Rol      int    `json:"rol_id"`
	Register int    `json:"people_id"`
}

type RegisterNewUserResponse struct {
	User    string `json:"user"`
	Rol     int    `json:"rol_id"`
	Message string `json:"message"`
}

type UserListView struct {
	ID       string `json:"id"`
	User     string `json:"user"`
	Rol      int    `json:"rol_id"`
	State    int    `json:"user_state_id"`
	Register int    `json:"people_id"`
}
type GetAllUserRequest struct{}

type AllUserResponse struct {
	UserListView []*UserListView `json:"userListView"`
	Count        int             `json:"count"`
}

func (user *User) ActivateAccount() error {
	user.State = 1
	return nil
}

func (user *User) ValidateUsername() error {

	return nil
}

func (user *User) ValidateNewPassword() error {

	return nil
}