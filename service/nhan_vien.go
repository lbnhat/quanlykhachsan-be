package service

import (
	"fmt"
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
	DangNhap(TenDangNhap, MatKhau string) (model.UserModel, error)
}

func (t *NhanVienService) DanhSachNhanVien() []model.NhanVien {
	result := t.repo.DanhSachNhanVien()
	return result
}

func (t *NhanVienService) DangNhap(TenDangNhap, MatKhau string) (rs model.UserModel, err error) {
	heThong := t.repo.LayThongTinDangNhap(TenDangNhap, MatKhau)
	fmt.Println(heThong)
	if heThong == nil {
		return rs, fmt.Errorf("Sai thông tin đăng nhập")
	}
	switch heThong.RoleID {
	case 1:
		khachHang := t.repo.LayThongTinKhachHang(heThong.IDDangNhap)
		return model.UserModel{
			ID:          khachHang.IdKhachHang,
			TenDangNhap: heThong.TenDangNhap,
			Name:        khachHang.TenKhachHang,
			Gender:      true,
			PhoneNumber: khachHang.Sdt,
			Email:       khachHang.Email,
			RoleId:      int64(heThong.RoleID),
			NgaySinh:    khachHang.NgaySinh,
		}, nil
	default:
		nhanVien := t.repo.LayThongTinNhanVien(heThong.IDDangNhap)
		return model.UserModel{
			ID:          nhanVien.IdNhanVien,
			TenDangNhap: heThong.TenDangNhap,
			Name:        nhanVien.TenNhanVien,
			Gender:      true,
			PhoneNumber: nhanVien.Sdt,
			// Email:       nhanVien.Email,
			RoleId:   int64(heThong.RoleID),
			NgaySinh: nhanVien.NgaySinh,
		}, nil
	}
	return model.UserModel{}, nil
	//return result
}
