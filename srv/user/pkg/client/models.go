package client

type UserRole string

const (
    TeacherRole UserRole = "TEACHER"
    StudentRole UserRole = "STUDENT"
)

type ParentInfo struct {
    ParentName  string `json:"parent_name"`
    ParentEmail string `json:"parent_email"`
    ParentCode  string `json:"parent_code"`
}

type Student struct {
    ID           uuid.UUID  `json:"id"`
    Name         string     `json:"name"`
    Role         UserRole   `json:"role"`
    Username     string     `json:"username"`
    Password     string     `json:"password"`
    Birthday     time.Time  `json:"date_of_birth,omitempty"`
    AcademicYear string     `json:"academic_year,omitempty"`
    ClassID      string     `json:"class_id,omitempty"`
    Parent       ParentInfo `json:"parent"`
}

type Teacher struct {
    ID          string   `json:"id"`
    Name        string   `json:"name"`
    Role        UserRole `json:"role"`
    Username    string   `json:"username"`
    Password    string   `json:"password"`
    SchoolName  string   `json:"school_name,omitempty"`
    Classes     []string `json:"classes,omitempty"`
}

type User struct {
    ID           string     `json:"id"`
    Name         string     `json:"name"`
    Role         UserRole   `json:"role"`
    Username     string     `json:"username"`
    Password     string     `json:"password"`
    Birthday     time.Time  `json:"date_of_birth,omitempty"`
    AcademicYear string     `json:"academic_year,omitempty"`
    ClassID      string     `json:"class_id,omitempty"`
    Parent       ParentInfo `json:"parent,omitempty"`
    SchoolName   string     `json:"school_name,omitempty"`
    Classes      []string   `json:"classes,omitempty"`
}