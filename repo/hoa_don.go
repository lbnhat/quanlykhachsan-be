package repo

import (
	"context"
	"quanlykhachsan/model"

	"gorm.io/gorm"
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

func (r *RepoPG) ThongTinHoaDon(id int) *model.ThongTinHoaDon {

	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	var hoaDon *model.ThongTinHoaDon
	tx = tx.Where("id_hoa_don = ?", id)
	if err = tx.Debug().Select("*").Table("hoa_don hd").Joins("inner join phieu_dat_phong pdp on pdp.id_phieu_dat_phong = hd.id_phieu_dat_phong").
		// Joins("inner join chi_tiet_phieu_dat_phong ctpdp on ctpdp.id_phieu_dat_phong = pdp.id_phieu_dat_phong").
		Preload("ThongTinPhong", func(db *gorm.DB) *gorm.DB {
			return db.Debug().Select("*").Table("chi_tiet_phieu_dat_phong ctpdp").Joins("inner join phong p on p.id_phong = ctpdp.id_phong").
				Joins("inner join loai_phong lp on lp.id_loai_phong = p.id_loai_phong")
		}).
		Preload("ThongTinDichVu", func(db *gorm.DB) *gorm.DB {
			return db.Debug().Select("*").Table("phieu_dich_vu pdv").Joins("inner join chi_tiet_dich_vu ctdv on pdv.id_phieu_dich_vu = ctdv.id_phieu_dich_vu").
				Joins("inner join dich_vu dv on dv.id_dich_vu = ctdv.id_dich_vu")
		}).
		Preload("ThongTinKhachHang", func(db *gorm.DB) *gorm.DB {
			return db.Debug().Table("khach_hang")
		}).
		Preload("ThongTinNhanVien", func(db *gorm.DB) *gorm.DB {
			return db.Debug().Table("nhan_vien")
		}).
		Find(&hoaDon).Error; err != nil {
		return nil
	}
	return hoaDon
}
