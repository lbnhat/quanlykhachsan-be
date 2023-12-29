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
	if err = tx.Raw("select * from dich_vu dv inner join chi_tiet_dich_vu ctdv  on dv.id_dich_vu  = ctdv .id_dich_vu ").Scan(&dichvu).Error; err != nil {
		return nil
	}

	return dichvu
}
