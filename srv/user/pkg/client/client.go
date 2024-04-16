package client

const (
    createUserPath         = "/create_user"
    loginPath              = "/login"
    logoutPath             = "/logout"
    editUserPath           = "/edit_user"
    editParentDetailsPath  = "/edit_parent_details"
)

type UserService interface {
	CreateUser(req CreateUserRequest)(*User, error)
	EditUser(req EditUserRequest)(*User, error)
	EditParentDetails(req EditParentDetailsRequest)(*User, error)
	Login(req LoginRequest)(*User, error)
	Logout(req LogoutRequest)(error)
}

type CreateUserRequest struct {
    Name         string `json:"name"`
    Username     string `json:"username"`
    Password     string `json:"password"`
    Role         string `json:"role"`
    ParentName   string `json:"parent_name"`
    ParentEmail  string `json:"parent_email"`
    ParentCode   string `json:"parent_code"`
    Birthday     string `json:"birthday"`
    AcademicYear string `json:"academic_year"`
    ClassID      string `json:"class_id"`
}

type EditUserRequest struct {
    Name         string `json:"name"`
    Username     string `json:"username"`
    Password     string `json:"password"`
    Role         string `json:"role"`
    ParentName   string `json:"parent_name"`
    ParentEmail  string `json:"parent_email"`
    ParentCode   string `json:"parent_code"`
    Birthday     string `json:"birthday"`
    AcademicYear string `json:"academic_year"`
    ClassID      string `json:"class_id"`
}

type EditParentDetailsRequest struct {
	StudentId uuid.UUID `json:"student_id,required"`
    ParentName  string `json:"parent_name"`
    ParentEmail string `json:"parent_email"`
    ParentCode  string `json:"parent_code"`
}

type LoginRequest struct {
	Username     string `json:"username"`
    Password     string `json:"password"`
    Role         string `json:"role"`
}