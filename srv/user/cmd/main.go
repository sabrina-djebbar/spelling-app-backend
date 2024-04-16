package cmd

func main() {
    var err error
    db, err = sql.Open("postgres", "postgres://username:password@localhost/database_name?sslmode=disable")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    r := gin.Default()
    
    r.POST(createUserPath, CreateUser)
    r.POST(loginPath, Login)
    r.POST(logoutPath, Logout)
    r.PUT(editUserPath, EditUser)
    r.PUT(editParentDetailsPath, EditParentDetails)
    
    r.Run(":8080")
}