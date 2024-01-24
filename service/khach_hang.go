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
	CapNhatThongTin(req model.KhachHangRequest) (rs model.UserModel, err error)
}

func (t *KhachHangService) DanhSachKhachHang() []model.KhachHang {
	result := t.repo.DanhSachKhachHang()
	return result
}

func (t *KhachHangService) CapNhatThongTin(req model.KhachHangRequest) (rs model.UserModel, err error) {
	kh := &model.KhachHang{
		IdKhachHang:  req.ID,
		TenKhachHang: req.Name,
		NgaySinh:     &req.NgaySinh,
		Sdt:          req.PhoneNumber,
	}
	err = t.repo.TimKhachHang(req.ID, kh)
	if err != nil {
		return rs, err
	}
	kh.TenKhachHang = req.Name
	kh.Sdt = req.PhoneNumber
	kh.NgaySinh = &req.NgaySinh

	err = t.repo.CapNhatKhachHang(kh)
	if err != nil {
		return rs, err
	}
	return model.UserModel{
		ID:          kh.IdKhachHang,
		TenDangNhap: kh.Email,
		Name:        kh.TenKhachHang,
		Gender:      true,
		PhoneNumber: kh.Sdt,
		Email:       kh.Email,
		RoleId:      1,
		NgaySinh: func() string {
			if kh.NgaySinh == nil {
				return ""
			} else {
				return *kh.NgaySinh
			}
		}(),
	}, nil
}
