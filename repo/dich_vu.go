package repo

import (
	"context"
	"quanlykhachsan/model"
)

// FindAll implements TagsRepository
func (r *RepoPG) DanhSachDichVu() []model.DanhSachDichVu {

	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	var dichvu []model.DanhSachDichVu
	if err = tx.Raw("select * from dich_vu").Scan(&dichvu).Error; err != nil {
		return nil
	}

	return dichvu
}

func (r *RepoPG) TaoPhieuDichvu(req *model.PhieuDichVu) error {
	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	if err = tx.Create(&req).Error; err != nil {
		return err
	}
	return nil
}

func (r *RepoPG) TaoChiTietDichVu(req *model.ChiTietDichVu) error {
	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	if err = tx.Create(&req).Error; err != nil {
		return err
	}
	return nil
}
