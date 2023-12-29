package controller

import (
	"net/http"

	_ "quanlykhachsan/docs"
	"quanlykhachsan/model"
	"quanlykhachsan/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type PhongController struct {
	phongService service.PhongServiceInterface
}

func NewPhongController(service service.PhongServiceInterface) *PhongController {
	return &PhongController{
		phongService: service,
	}
}

func (controller *PhongController) DanhSachPhong(ctx *gin.Context) {
	log.Info().Msg("findAll tags")
	phongResponse := controller.phongService.DanhSachPhong()
	webResponse := model.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   phongResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
