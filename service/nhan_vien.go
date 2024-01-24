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
	DangKyKhachHang(req model.DangKy) (rs model.UserModel, err error)
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
			NgaySinh: func() string {
				if khachHang.NgaySinh == nil {
					return ""
				} else {
					return *khachHang.NgaySinh
				}
			}(),
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
	//return model.UserModel{}, nil
	//return result
}

func (t *NhanVienService) DangKyKhachHang(req model.DangKy) (rs model.UserModel, err error) {
	heThong := t.repo.KiemTraThongTinDangKy(req.Email)
	fmt.Println(heThong)
	if heThong == nil {
		kh := model.KhachHang{
			TenKhachHang: req.Name,
			Sdt:          req.PhoneNumber,
			Email:        req.Email,
		}
		err = t.repo.TaoKhachHang(&kh)
		if err != nil {
			return rs, fmt.Errorf("Sai thông tin đăng kí")
		}
		tk := model.HeThong{
			IDDangNhap:  kh.IdKhachHang,
			TenDangNhap: req.Email,
			MatKhau:     req.Password,
			RoleID:      1,
		}
		err = t.repo.TaoTaiKhoan(&tk)
		if err != nil {
			return rs, fmt.Errorf("Sai thông tin đăng kí")
		}
		return model.UserModel{
			ID:          tk.IDDangNhap,
			TenDangNhap: tk.TenDangNhap,
			Name:        kh.TenKhachHang,
			PhoneNumber: kh.Sdt,
			Email:       tk.TenDangNhap,
			RoleId:      1,
		}, nil
	} else {
		return rs, fmt.Errorf("Email đăng kí đã tồn tại!")
	}

}
