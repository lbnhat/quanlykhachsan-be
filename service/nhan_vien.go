package service

import (
	"quanlykhachsan/model"
	"quanlykhachsan/repo"

	"github.com/go-playground/validator/v10"
)

type NhanVienService struct {
	repo     repo.PGInterface
	Validate *validator.Validate
}

func NewNhanVienService(repo repo.PGInterface, validate *validator.Validate) NhanVienServiceInterface {
	return &NhanVienService{
		repo:     repo,
		Validate: validate,
	}
}

type NhanVienServiceInterface interface {
	DanhSachNhanVien() []model.NhanVien
}

func (t *NhanVienService) DanhSachNhanVien() []model.NhanVien {
	result := t.repo.DanhSachNhanVien()
	return result
}
