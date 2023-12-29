package service

import (
	"quanlykhachsan/model"
	"quanlykhachsan/repo"

	"github.com/go-playground/validator/v10"
)

type DichVuService struct {
	repo     repo.PGInterface
	Validate *validator.Validate
}

func NewDichVuService(repo repo.PGInterface, validate *validator.Validate) DichVuServiceInterface {
	return &DichVuService{
		repo:     repo,
		Validate: validate,
	}
}

type DichVuServiceInterface interface {
	DanhSachDichVu() []model.DanhSachDichVu
}

func (t *DichVuService) DanhSachDichVu() []model.DanhSachDichVu {
	result := t.repo.DanhSachDichVu()
	return result
}
