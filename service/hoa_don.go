package service

import (
	"fmt"
	"quanlykhachsan/model"
	"quanlykhachsan/repo"
	"time"

	"github.com/go-playground/validator/v10"
)

type HoaDonService struct {
	repo     repo.PGInterface
	Validate *validator.Validate
}

func NewHoaDonService(repo repo.PGInterface, validate *validator.Validate) HoaDonServiceInterface {
	return &HoaDonService{
		repo:     repo,
		Validate: validate,
	}
}

type HoaDonServiceInterface interface {
	DanhSachHoaDon() []model.DanhSachHoaDon
	ThongTinHoaDon(idHoaDon int) *model.ThongTinHoaDon
}

func (t *HoaDonService) DanhSachHoaDon() []model.DanhSachHoaDon {
	result := t.repo.DanhSachHoaDon()
	return result
}

func (t *HoaDonService) ThongTinHoaDon(idHoaDon int) *model.ThongTinHoaDon {
	result := t.repo.ThongTinHoaDon(idHoaDon)
	if result != nil && len(result.ThongTinPhong) > 0 {
		result.NgayTraPhong = result.ThongTinPhong[0].NgayTraPhong
		// Chuyển đổi chuỗi thành đối tượng time.Time
		date1, err1 := time.Parse(time.RFC3339, result.NgayTraPhong)
		date2, err2 := time.Parse(time.RFC3339, result.NgayDen)
		fmt.Println(result.NgayDen)
		fmt.Println(result.NgayTraPhong)
		// Kiểm tra lỗi khi chuyển đổi
		if err1 != nil || err2 != nil {
			fmt.Println("Lỗi khi chuyển đổi chuỗi thành đối tượng time.Time")
		}
		result.SoNgay = int((date1.Sub(date2).Hours() / 24))
	}

	if result.SoNgay != 0 {
		for i, v := range result.ThongTinPhong {
			result.ThongTinPhong[i].Gia = v.Gia * result.SoNgay
		}
	}

	return result
}
