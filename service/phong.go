package service

import (
	"quanlykhachsan/model"
	"quanlykhachsan/repo"

	"github.com/go-playground/validator/v10"
)

type PhongService struct {
	repo     repo.PGInterface
	Validate *validator.Validate
}

func NewPhongService(repo repo.PGInterface, validate *validator.Validate) PhongServiceInterface {
	return &PhongService{
		repo:     repo,
		Validate: validate,
	}
}

type PhongServiceInterface interface {
	DanhSachPhong() []model.DanhSachPhong
	DanhSachPhongTrong(checkinDate, checkoutDate, hangPhong, loaiPhong string) []model.DanhSachPhong
}

func (t *PhongService) DanhSachPhong() []model.DanhSachPhong {
	result := t.repo.DanhSachPhong()
	return result
}

func (t *PhongService) DanhSachPhongTrong(checkinDate, checkoutDate, hangPhong, loaiPhong string) []model.DanhSachPhong {
	result := t.repo.DanhSachPhongTrong(checkinDate, checkoutDate, hangPhong, loaiPhong)
	return result
}
