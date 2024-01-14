package repo

import (
	"context"
	"quanlykhachsan/model"

	"gorm.io/gorm"
)

// FindAll implements TagsRepository
func (r *RepoPG) DanhSachPhieuDatPhong(idKhach int) []model.DanhSachPhieuDatPhong {

	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	var phieudatphong []model.DanhSachPhieuDatPhong
	if idKhach != 0 {
		tx = tx.Where("id_khach_hang = ?", idKhach)
	}
	if err = tx.Debug().Select("*").Table("phieu_dat_phong pdp").
		// Joins("inner join chi_tiet_phieu_dat_phong ctpdp on ctpdp.id_phieu_dat_phong = pdp.id_phieu_dat_phong").
		Preload("ThongTinPhong", func(db *gorm.DB) *gorm.DB {
			return db.Debug().Select("*").Table("chi_tiet_phieu_dat_phong ctpdp").Joins("inner join phong p on p.id_phong = ctpdp.id_phong").
				Joins("inner join loai_phong lp on lp.id_loai_phong = p.id_loai_phong")
		}).
		Preload("ThongTinKhachHang", func(db *gorm.DB) *gorm.DB {
			return db.Debug().Table("khach_hang")
		}).
		Find(&phieudatphong).Error; err != nil {
		return nil
	}
	return phieudatphong
}

func (r *RepoPG) TaoPhieuDatPhong(req *model.PhieuDatPhong) error {
	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	if err = tx.Create(&req).Error; err != nil {
		return err
	}
	return nil
}

func (r *RepoPG) TaoChiTietPhieuDatPhong(req *model.ChiTietPhieuDatPhong) error {
	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	if err = tx.Create(&req).Error; err != nil {
		return err
	}
	return nil
}
