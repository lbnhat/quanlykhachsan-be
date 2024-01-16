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

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
}

func handleOptions(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
}

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
	datphongService := service.NewDatPhongService(repoPG, validate)
	hoadonService := service.NewHoaDonService(repoPG, validate)

	// Controller
	khacHangController := controller.NewKhachHangController(khachhangService)
	nhanVienController := controller.NewNhanVienController(nhanvienService)
	phongController := controller.NewPhongController(phongService)
	dichvuController := controller.NewDichVuController(dichvuService)
	datphongController := controller.NewDatPhongController(datphongService)
	hoadonController := controller.NewHoaDonController(hoadonService)

	mockController := controller.NewMockController()

	router := gin.Default()
	http.HandleFunc("/", handleOptions)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST, GET, OPTIONS, PUT, DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
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
	nhanvienRouter.POST("/dang-nhap", nhanVienController.Login)
	nhanvienRouter.POST("/dang-ki", nhanVienController.Login)
	// Phong

	phongRouter := baseRouter.Group("/phong")
	phongRouter.GET("", phongController.DanhSachPhong)
	phongRouter.GET("/danh-sach-phong-trong", phongController.DanhSachPhongTrong)
	// Dich vụ
	dichvuRouter := baseRouter.Group("/dich-vu")
	dichvuRouter.GET("", dichvuController.DanhSachDichVu)

	datphongRouter := baseRouter.Group("/dat-phong")
	datphongRouter.GET("", datphongController.DanhSachPhieuDatPhong)
	datphongRouter.GET("/khach-hang", datphongController.DanhSachPhieuDatPhongTuKhachHang)
	datphongRouter.POST("/dat", datphongController.DatPhong)
	datphongRouter.POST("/cap-nhat-trang-thai", datphongController.CapNhatPhong)

	hoadonRouter := baseRouter.Group("/hoa-don")
	hoadonRouter.GET("", hoadonController.DanhSachHoaDon)

	//mock api
	mockRouter := baseRouter.Group("/mock")
	mockRouter.GET("hotel/search", mockController.HotelSearch)
	mockRouter.GET("hotel/room/search", mockController.Room)

	return router
}
