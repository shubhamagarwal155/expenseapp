package dtos

type RegisterUserDTO struct {
	FirstName  *string `json:"firstname"`
	MiddleName *string `json:"middlename"`
	LastName   *string `json:"lastname"`
	Email      *string `json:"email"`
	UserName   *string `json:"username"`
	Password   *string `json:"password"`
}
