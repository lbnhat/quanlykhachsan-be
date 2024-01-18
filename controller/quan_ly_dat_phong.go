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
	trangThai := ctx.Query("trang_thai")
	phieuDatPhongResponse := controller.datPhongService.DanhSachPhieuDatPhong(trangThai)
	webResponse := model.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   phieuDatPhongResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *DatPhongController) CapNhatPhong(ctx *gin.Context) {
	log.Info().Msg("findAll tags")
	var requestBody model.CapNhatPhong
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	phieuDatPhongResponse := controller.datPhongService.CapNhatPhong(requestBody)
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
	err := controller.datPhongService.DatPhong(requestBody)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Đặt phòng thất bại"})
		return
	}
	webResponse := model.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		//Data:   res,
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

func (controller *DatPhongController) BaoCaoTheoNgay(ctx *gin.Context) {
	log.Info().Msg("findAll tags")
	option := ctx.Query("option")
	phieuDatPhongResponse := controller.datPhongService.BaoCao(option)
	webResponse := model.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   phieuDatPhongResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
