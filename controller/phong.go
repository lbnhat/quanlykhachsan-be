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

func (controller *PhongController) DanhSachPhongTrong(ctx *gin.Context) {
	log.Info().Msg("findAll tags")
	checkinDate := ctx.Query("checkin_date")
	checkoutDate := ctx.Query("checkout_date")
	hangPhong := ctx.Query("hang_phong")
	loaiPhong := ctx.Query("loai_phong")
	phongResponse := controller.phongService.DanhSachPhongTrong(checkinDate, checkoutDate, hangPhong, loaiPhong)
	// webResponse := model.Response{
	// 	Code:   http.StatusOK,
	// 	Status: "Ok",
	// 	Data:   phongResponse,
	// }
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, phongResponse)

}
