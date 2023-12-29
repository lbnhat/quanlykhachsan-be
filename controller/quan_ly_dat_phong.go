package controller

import (
	"net/http"

	_ "quanlykhachsan/docs"
	"quanlykhachsan/model"
	"quanlykhachsan/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type PhieuDatPhongController struct {
	phieuDatPhongService service.PhieuDatPhongServiceInterface
}

func NewPhieuDatPhongController(service service.PhieuDatPhongServiceInterface) *PhieuDatPhongController {
	return &PhieuDatPhongController{
		phieuDatPhongService: service,
	}
}

func (controller *PhieuDatPhongController) DanhSachPhieuDatPhong(ctx *gin.Context) {
	log.Info().Msg("findAll tags")
	phieuDatPhongResponse := controller.phieuDatPhongService.DanhSachPhieuDatPhong()
	webResponse := model.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   phieuDatPhongResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
