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
