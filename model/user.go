package model

type ShowUserList struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Term       int    `json:"term"`
	Permission string `json:"permission"`
}

type User struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Term         int    `json:"term"`
	Bio          string `json:"bio"`
	Image        string `json:"image"`
	PermissionId int    `json:"permission_id"`
}

type ShowUser struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Term       int    `json:"term"`
	Bio        string `json:"bio"`
	Permission string `json:"permission"`
}

type LoginUser struct {
	Email    int    `json:"email"`
	Password string `json:"password"`
}

type UpdateUserDetails struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Term int    `json:"term"`
	Bio  string `json:"bio"`
}

type UpdateUserPermission struct {
	Id         string `json:"id"`
	Permission int    `json:"permission_id"`
}

type Response struct {
	Id string
}
