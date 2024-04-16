package rest

func CreateUser(c echo.Context) error {
	log.Println("Creating User")
	// create a new object of User model
	u := new(db.User)
	// bind request body to the model object
	if err := c.Bind(u); err != nil {
		panic(err)
	}
	// call the service
	err := service.CreateUser(u)
	if err != nil {
		log.Println(err)
	}
	//return success response
	return c.String(http.StatusCreated, "User created successfully")
}