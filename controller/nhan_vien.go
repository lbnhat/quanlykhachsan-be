package controller

import (
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
