package repo

import (
	"context"
	"quanlykhachsan/model"
)

// FindAll implements TagsRepository
func (r *RepoPG) DanhSachHoaDon() []model.HoaDon {

	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	var hoadon []model.HoaDon
	if err = tx.Raw(`select * 
	from hoa_don hd 
	inner join  phieu_dat_phong pdp  on
	pdp .id_phieu_dat_phong =hd .id_phieu_dat_phong `).Scan(&hoadon).Error; err != nil {
		return nil
	}
	return hoadon
}
