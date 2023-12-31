package service

import (
	"quanlykhachsan/model"
	"quanlykhachsan/repo"

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
	DanhSachHoaDon() []model.HoaDon
}

func (t *HoaDonService) DanhSachHoaDon() []model.HoaDon {
	result := t.repo.DanhSachHoaDon()
	return result
}
