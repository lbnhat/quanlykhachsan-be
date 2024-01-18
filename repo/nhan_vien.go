package repo

import (
	"context"
	"fmt"
	"quanlykhachsan/model"
)

// FindAll implements TagsRepository
func (r *RepoPG) DanhSachNhanVien() []model.NhanVien {

	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	var nhanvien []model.NhanVien
	if err = tx.Find(&nhanvien).Error; err != nil {
		return nil
	}

	return nhanvien
}

func (r *RepoPG) LayThongTinDangNhap(TenDangNhap, MatKhau string) *model.HeThong {
	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	//var err error
	heThong := &model.HeThong{}
	if err := tx.Debug().Where("ten_dang_nhap = ? and mat_khau = ?", TenDangNhap, MatKhau).First(heThong).Error; err != nil {
		fmt.Println(err)
		return nil
	}
	return heThong
}

func (r *RepoPG) LayThongTinNhanVien(IdNhanVien int) *model.NhanVien {
	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	var nhanVien model.NhanVien
	if err = tx.Where("id_nhan_vien = ?", IdNhanVien).First(&nhanVien).Error; err != nil {
		return nil
	}

	return &nhanVien
}
