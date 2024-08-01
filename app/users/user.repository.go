package users

import (
	"database/sql"

	"go.mod/app/dtos"
	"go.mod/pkg/utils"
)

type Repository struct {
	db *sql.DB
}

func (repo *Repository) LoginUser(dto *dtos.LoginUserDTO) (result *dtos.Result, err error) {

	result = &dtos.Result{}
	var userId int
	var jwtToken string

	err = repo.db.QueryRow("select userid from user_info where username = ?", dto.UserName).Scan(&userId)

	if err == nil {
		if userId != 0 {
			// user found now return a jwt token for this user
			jwtToken, err = utils.GenerateJWTAccessToken(userId)

			if err == nil {
				result.Success = true
				result.Message = "User logged in successfully"
				result.Data = jwtToken
			}
		} else {
			result.Success = false
			result.Message = "Incorrect username or password"
		}
	}

	return result, err
}

func (repo *Repository) RegisterUser(dto *dtos.RegisterUserDTO) (result *dtos.Result, err error) {

	result = &dtos.Result{}
	var userId int
	var hashedPassword string

	//check if user with entered credentials exists
	err = repo.db.QueryRow("Select userId from user_info where username = ?", dto.UserName).Scan(&userId)

	if err == nil {
		if userId == 0 {
			// hash the password
			hashedPassword, err = utils.HashPassword(*dto.Password)

			if err == nil {
				_, err = repo.db.Query("INSERT INTO user_info(firstname,middlename,lastname,email,username,password) VALUES(?,?,?,?,?,?)", dto.FirstName, dto.MiddleName, dto.LastName, dto.Email, dto.UserName, hashedPassword)

				if err == nil {
					result.Success = true
					result.Message = "User registered successfully"
				}
			}
		} else {
			result.Success = false
			result.Message = "User with the entered credentials already exists"
		}
	}

	return result, err
}
