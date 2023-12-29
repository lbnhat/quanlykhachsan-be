package service

import (
	"quanlykhachsan/model"
	"quanlykhachsan/repo"

	"github.com/go-playground/validator/v10"
)

type PhieuDatPhongService struct {
	repo     repo.PGInterface
	Validate *validator.Validate
}

func NewPhieuDatPhongService(repo repo.PGInterface, validate *validator.Validate) PhieuDatPhongServiceInterface {
	return &PhieuDatPhongService{
		repo:     repo,
		Validate: validate,
	}
}

type PhieuDatPhongServiceInterface interface {
	DanhSachPhieuDatPhong() []model.DanhSachPhieuDatPhong
}

func (t *PhieuDatPhongService) DanhSachPhieuDatPhong() []model.DanhSachPhieuDatPhong {
	result := t.repo.DanhSachPhieuDatPhong()
	return result
}
