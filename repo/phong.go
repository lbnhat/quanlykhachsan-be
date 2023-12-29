package repo

import (
	"context"
	"quanlykhachsan/model"
)

// FindAll implements TagsRepository
func (r *RepoPG) DanhSachPhong() []model.DanhSachPhong {

	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	var phong []model.DanhSachPhong
	if err = tx.Raw("select *  from phong p  inner join loai_phong lp on p.id_loai_phong = lp .id_loai_phong ").Scan(&phong).Error; err != nil {
		return nil
	}

	return phong
}
