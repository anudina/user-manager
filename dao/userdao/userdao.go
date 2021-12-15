package userdao

type User struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
}
type Object interface{}

type UserRequest struct {
	Request User
}
type UserResponse struct {
	Status   Object `json:"Status"`
	Error    Object `json:"Error"`
	Response Object `json:"Response"`
}

var UserList = []User{
	{UserId: "Badigan", UserName: "Anand"}, {UserId: "JD", UserName: "Jayadev"}, {UserId: "Ravi", UserName: "Ravi"},
}
