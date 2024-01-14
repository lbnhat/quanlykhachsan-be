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

type DatPhongController struct {
	datPhongService service.DatPhongServiceInterface
}

func NewDatPhongController(service service.DatPhongServiceInterface) *DatPhongController {
	return &DatPhongController{
		datPhongService: service,
	}
}

func (controller *DatPhongController) DanhSachPhieuDatPhong(ctx *gin.Context) {
	log.Info().Msg("findAll tags")
	phieuDatPhongResponse := controller.datPhongService.DanhSachPhieuDatPhong()
	webResponse := model.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   phieuDatPhongResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *DatPhongController) DatPhong(ctx *gin.Context) {
	log.Info().Msg("findAll tags")

	var requestBody model.DatPhongRequest
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	res := controller.datPhongService.DatPhong(requestBody)
	webResponse := model.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   res,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *DatPhongController) DanhSachPhieuDatPhongTuKhachHang(ctx *gin.Context) {
	log.Info().Msg("findAll tags")
	idKhachHang := ctx.Query("user_id")
	iduser, err := strconv.Atoi(idKhachHang)
	if err != nil {
		iduser = 0
	}
	phieuDatPhongResponse := controller.datPhongService.DanhSachPhieuDatPhongTuKhach(iduser)
	webResponse := model.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   phieuDatPhongResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
