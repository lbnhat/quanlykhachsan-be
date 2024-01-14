package repo

import (
	"context"
	"quanlykhachsan/model"
)

// FindAll implements TagsRepository
func (r *RepoPG) DanhSachKhachHang() []model.KhachHang {
	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	var khachhang []model.KhachHang
	if err = tx.Find(&khachhang).Error; err != nil {
		return nil
	}

	return khachhang
}

func (r *RepoPG) TaoKhachHang(req *model.KhachHang) error {
	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	if err = tx.Create(&req).Error; err != nil {
		return err
	}
	return nil
}

func (r *RepoPG) TimKhachHang(id int, req *model.KhachHang) error {
	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	if err = tx.Where("id_khach_hang = ?", id).Find(&req).Error; err != nil {
		return err
	}
	return nil
}
