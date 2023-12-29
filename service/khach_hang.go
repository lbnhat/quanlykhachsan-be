package service

import (
	"quanlykhachsan/model"
	"quanlykhachsan/repo"

	"github.com/go-playground/validator/v10"
)

type KhachHangService struct {
	repo     repo.PGInterface
	Validate *validator.Validate
}

func NewKhachHangService(repo repo.PGInterface, validate *validator.Validate) KhachHangServiceInterface {
	return &KhachHangService{
		repo:     repo,
		Validate: validate,
	}
}

type KhachHangServiceInterface interface {
	DanhSachKhachHang() []model.KhachHang
}

func (t *KhachHangService) DanhSachKhachHang() []model.KhachHang {
	result := t.repo.DanhSachKhachHang()
	return result
}
