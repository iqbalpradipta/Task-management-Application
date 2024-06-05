package controllers

import (
	"net/http"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/go-playground/validator/v10"
	"github.com/iqbalpradipta/Task-management-Application/BE/src/helpers"
	"github.com/iqbalpradipta/Task-management-Application/BE/src/middleware"
	"github.com/iqbalpradipta/Task-management-Application/BE/src/model"
	"github.com/iqbalpradipta/Task-management-Application/BE/src/services"
	"github.com/labstack/echo/v4"
)

type userService struct {
	UserRepository services.Users
}

func UserController(userRepository services.Users) *userService {
	return &userService{userRepository}
}

func (u *userService) GetUsers(c echo.Context) error {
	users, err := u.UserRepository.GetUsers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Error when get users !"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success Get Users", users))
}

func (u *userService) GetUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := u.UserRepository.GetUserById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Error when get data by Id"))
	}
	
	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success get data By Id" ,convResponse(user)))
}

func (u *userService) CreateUser(c echo.Context) error {
    userRequest := new(model.UsersRequest)

    if err := c.Bind(&userRequest); err != nil {
        return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed to create cause your request time out"))
    }

    validation := validator.New()
    err := validation.Struct(userRequest)

    if err != nil {
        return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Error when Validation !"))
    }

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 14)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Error when hashing password !"))
	}
	
    user := model.Users{
        Name: userRequest.Name,
		Email: userRequest.Email,
		Password: string(hashPassword),
		Created_at: time.Now(),
    }


    createdUser, err := u.UserRepository.CreateUsers(user)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Failed to create user in the database"))
    }

    return c.JSON(http.StatusOK, helpers.SuccessResponse("Success Create Users", createdUser))
}

func (u *userService) Login(c echo.Context) error {
	req := new(model.Login)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Request Time Out!"))
	}

	user, err := u.UserRepository.GetUserByEmail(req.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Login failed ! please check again your data !"))
	}

	errCompare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if errCompare != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Your email/password is wrong!"))
	}

	if req.Email != user.Email {
		return c.JSON(http.StatusNotAcceptable, helpers.FailedResponse("Your email/password is wrong!"))
	}

	token, err := middleware.CreateToken(user.ID, user.Email, user.Name)
	if err != nil {
		c.JSON(http.StatusUnauthorized, helpers.FailedResponse("Unauthorized"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success Login", token))

}


func (u *userService) UpdateUsers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := u.UserRepository.GetUserById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Update data failed ! please check again your data !"))
	}

	req := new(model.UsersRequestUpdate)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Your request is Timeout !"))
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Error when hashing password !"))
	}

	hashConv := string(hashPassword)	

	if req.Name != "" {
		user.Name = req.Name
	}

	if req.Email != "" {
		user.Email = req.Email
	}

	if req.Password != "" {
		user.Password = hashConv
	}

	user.Updated_at = time.Now()

	response, err := u.UserRepository.UpdatedUsers(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Failed to update data"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success Update Data", response))
}

func (u *userService) DeleteUsers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := u.UserRepository.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Error Find Users Id"))
	}

	data, err := u.UserRepository.DeleteUsers(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Error delete data cause data not found !"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success delete Users", convResponse(data)))
}

func convResponse(u model.Users) model.UserResponse {
	return model.UserResponse{
		ID: u.ID,
		Name: u.Name,
		Email: u.Email,
	}
}