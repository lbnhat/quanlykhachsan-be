package controller

import (
	"net/http"

	_ "quanlykhachsan/docs"
	"quanlykhachsan/model"
	"quanlykhachsan/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type DichVuController struct {
	dichVuService service.DichVuServiceInterface
}

func NewDichVuController(service service.DichVuServiceInterface) *DichVuController {
	return &DichVuController{
		dichVuService: service,
	}
}

func (controller *DichVuController) DanhSachDichVu(ctx *gin.Context) {
	log.Info().Msg("findAll tags")
	dichvuResponse := controller.dichVuService.DanhSachDichVu()
	webResponse := model.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   dichvuResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
