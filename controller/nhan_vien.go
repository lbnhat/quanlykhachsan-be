package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

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
		ID          int    `json:"id"`
		TenDangNhap string `json:"ten_dang_nhap"`
		Password    string `json:"password"`
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		Gender      bool   `json:"gender"`
		PhoneNumber string `json:"phoneNumber"`
		Email       string `json:"email"`
		RoleId      int64  `json:"roleId"`
		NgaySinh    string `json:"ngay_sinh"`
		Image       string `json:"image"`
	}

	user := UserModel{
		ID:          5,
		TenDangNhap: "nhat_le@gmail.com.com",
		Password:    "secret123",
		FirstName:   "Le",
		LastName:    "Nhat",
		Gender:      true,
		PhoneNumber: "123456789",
		Email:       "nhat_le@gmail.com.com",
		RoleId:      1,
		NgaySinh:    "2001-10-10",
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
	response["roleId"] = 1
	jsonResponse, _ := json.Marshal(response)
	fmt.Println(string(jsonResponse))
	// webResponse := model.Response{
	// 	Code:   http.StatusOK,
	// 	Status: "Ok",
	// 	Data:   response,
	// }
	ctx.JSON(http.StatusOK, response)

}
