package controller

import (
	"net/http"
	"strconv"

	_ "quanlykhachsan/docs"
	"quanlykhachsan/model"
	"quanlykhachsan/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type HoaDonController struct {
	hoaDonService service.HoaDonServiceInterface
}

func NewHoaDonController(service service.HoaDonServiceInterface) *HoaDonController {
	return &HoaDonController{
		hoaDonService: service,
	}
}

func (controller *HoaDonController) DanhSachHoaDon(ctx *gin.Context) {
	log.Info().Msg("findAll tags")
	hoadonResponse := controller.hoaDonService.DanhSachHoaDon()
	webResponse := model.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   hoadonResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *HoaDonController) ThongTinHoaDon(ctx *gin.Context) {
	log.Info().Msg("findAll tags")
	id := ctx.Param("id")
	idHoaDon, _ := strconv.Atoi(id)
	hoadonResponse := controller.hoaDonService.ThongTinHoaDon(idHoaDon)
	webResponse := model.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   hoadonResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
