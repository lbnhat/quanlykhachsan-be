package repo

import (
	"context"
	"quanlykhachsan/model"
)

// FindAll implements TagsRepository
func (r *RepoPG) DanhSachHoaDon() []model.DanhSachHoaDon {
	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	var hoadon []model.DanhSachHoaDon
	if err = tx.Raw(`select
						*
					from
						hoa_don hd
					inner join phieu_dat_phong pdp on
						pdp.id_phieu_dat_phong = hd .id_phieu_dat_phong
					inner join khach_hang kh on
						kh.id_khach_hang = pdp.id_khach_hang
					inner join nhan_vien nv on
						pdp.id_nhan_vien = nv.id_nhan_vien
					`).Scan(&hoadon).Error; err != nil {
		return nil
	}
	return hoadon
}

func (r *RepoPG) TaoHoaDon(req *model.HoaDon) error {
	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	if err = tx.Create(&req).Error; err != nil {
		return err
	}
	return nil
}
