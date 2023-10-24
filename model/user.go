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
	Email        string `json:"email"`
	Term         int    `json:"term"`
	Bio          string `json:"bio"`
	PermissionId int    `json:"permission"`
}

type ShowUser struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Term       int    `json:"term"`
	Bio        string `json:"bio"`
	Permission string `json:"permission"`
}

type LoginUser struct {
	Email    int    `json:"email"`
	Password string `json:"password"`
}

type UpdateUserDetails struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Term  int    `json:"term"`
	Bio   string `json:"bio"`
}

type UpdateUserPermission struct {
	Id         string `json:"id"`
	Permission int    `json:"permission"`
}

//type UserImage struct {
//	Id    string `json:"id"`
//	Image string `json:"image"`
//}

type Response struct {
	Id string
}
