package repo

import (
	"context"
	"quanlykhachsan/model"

	"gorm.io/gorm"
)

// FindAll implements TagsRepository
func (r *RepoPG) DanhSachPhieuDatPhong() []model.DanhSachPhieuDatPhong {

	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	var phieudatphong []model.DanhSachPhieuDatPhong
	if err = tx.Select("*").Table("phieu_dat_phong pdp").
		Joins("inner join chi_tiet_phieu_dat_phong ctpdp on ctpdp.id_phieu_dat_phong = pdp.id_phieu_dat_phong").Preload("ThongTinKhachHang", func(db *gorm.DB) *gorm.DB {
		return db.Table("khach_hang")
	}).
		Find(&phieudatphong).Error; err != nil {
		return nil
	}
	return phieudatphong
}
