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

type KhachHangController struct {
	khachHangService service.KhachHangServiceInterface
}

func NewKhachHangController(service service.KhachHangServiceInterface) *KhachHangController {
	return &KhachHangController{
		khachHangService: service,
	}
}

func (controller *KhachHangController) DanhSachKhachHang(ctx *gin.Context) {
	log.Info().Msg("findAll tags")
	khachhangResponse := controller.khachHangService.DanhSachKhachHang()
	webResponse := model.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   khachhangResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *KhachHangController) CapNhatKhachHang(ctx *gin.Context) {
	var requestBody model.KhachHangRequest
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	rs, err := controller.khachHangService.CapNhatThongTin(requestBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.Header("Content-Type", "application/json")

	response := make(map[string]interface{})
	response["user"] = rs
	jsonResponse, _ := json.Marshal(response)
	fmt.Println(string(jsonResponse))
	ctx.JSON(http.StatusOK, response)

}
