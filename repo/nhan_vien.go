package repo

import (
	"context"
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
