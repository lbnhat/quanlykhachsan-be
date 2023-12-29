package router

import (
	"quanlykhachsan/config"
	"quanlykhachsan/controller"
	"quanlykhachsan/repo"
	"quanlykhachsan/service"
	"time"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	db := config.DatabaseConnection()
	validate := validator.New()

	// Repository
	repoPG := repo.NewPGRepo(db)

	// Service
	khachhangService := service.NewKhachHangService(repoPG, validate)
	nhanvienService := service.NewNhanVienService(repoPG, validate)
	phongService := service.NewPhongService(repoPG, validate)
	dichvuService := service.NewDichVuService(repoPG, validate)
	phieudatphongService := service.NewPhieuDatPhongService(repoPG, validate)
	hoadonService := service.NewHoaDonService(repoPG, validate)

	// Controller
	khacHangController := controller.NewKhachHangController(khachhangService)
	nhanVienController := controller.NewNhanVienController(nhanvienService)
	phongController := controller.NewPhongController(phongService)
	dichvuController := controller.NewDichVuController(dichvuService)
	phieudatphongController := controller.NewPhieuDatPhongController(phieudatphongService)
	hoadonController := controller.NewHoaDonController(hoadonService)
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")

	//Khach Hang
	khachhangRouter := baseRouter.Group("/khach-hang")
	khachhangRouter.GET("", khacHangController.DanhSachKhachHang)
	// Nhan vien
	nhanvienRouter := baseRouter.Group("/nhan-vien")
	nhanvienRouter.GET("", nhanVienController.DanhSachNhanVien)
	// Phong

	phongRouter := baseRouter.Group("/phong")
	phongRouter.GET("", phongController.DanhSachPhong)
	// Dich vụ
	dichvuRouter := baseRouter.Group("/dich-vu")
	dichvuRouter.GET("", dichvuController.DanhSachDichVu)

	phieudatphongRouter := baseRouter.Group("/phieu-dat-phong")
	phieudatphongRouter.GET("", phieudatphongController.DanhSachPhieuDatPhong)

	hoadonRouter := baseRouter.Group("/hoa-don")
	hoadonRouter.GET("", hoadonController.DanhSachHoaDon)

	return router
}
