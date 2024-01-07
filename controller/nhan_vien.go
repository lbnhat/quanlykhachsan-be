package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "quanlykhachsan/docs"
	"quanlykhachsan/model"
	"quanlykhachsan/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type NhanVienController struct {
	nhanVienService service.NhanVienServiceInterface
}

func NewNhanVienController(service service.NhanVienServiceInterface) *NhanVienController {
	return &NhanVienController{
		nhanVienService: service,
	}
}

func (controller *NhanVienController) DanhSachNhanVien(ctx *gin.Context) {
	log.Info().Msg("findAll tags")
	nhanvienResponse := controller.nhanVienService.DanhSachNhanVien()
	webResponse := model.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nhanvienResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
func (controller *NhanVienController) Login(ctx *gin.Context) {

	type UserModel struct {
		Username    string    `json:"username"`
		Password    string    `json:"password"`
		FirstName   string    `json:"firstName"`
		LastName    string    `json:"lastName"`
		Gender      bool      `json:"gender"`
		PhoneNumber string    `json:"phoneNumber"`
		Email       string    `json:"email"`
		RoleId      int64     `json:"roleId"`
		Birthday    time.Time `json:"birthday"`
		Image       string    `json:"image"`
	}

	user := UserModel{
		Username:    "john_doe",
		Password:    "secret123",
		FirstName:   "John",
		LastName:    "Doe",
		Gender:      true,
		PhoneNumber: "123456789",
		Email:       "john.doe@example.com",
		RoleId:      2,
		Birthday:    time.Now(),
		Image:       "profile.jpg",
	}
	// webResponse := model.Response{
	// 	Code:   http.StatusOK,
	// 	Status: "Ok",
	// 	Data:   nhanvienResponse,
	// }
	ctx.Header("Content-Type", "application/json")

	response := make(map[string]interface{})
	response["user"] = user
	jsonResponse, _ := json.Marshal(response)
	fmt.Println(string(jsonResponse))
	// webResponse := model.Response{
	// 	Code:   http.StatusOK,
	// 	Status: "Ok",
	// 	Data:   response,
	// }
	ctx.JSON(http.StatusOK, response)

}
