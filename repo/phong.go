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

func (r *RepoPG) DanhSachPhongTrong(checkinDate, checkoutDate, hangPhong, loaiPhong string) []model.DanhSachPhong {

	// tx, cancel := r.DBWithTimeout(context.Background())
	// defer cancel()
	// var err error
	// var phong []model.DanhSachPhong
	// if err = tx.Raw(`select
	// 					*
	// 				from
	// 					phong p inner join loai_phong lp on  p.id_loai_phong =lp.id_loai_phong
	// 				left join chi_tiet_phieu_dat_phong ctpdp on
	// 					p.id_phong  = ctpdp.id_phong
	// 					and ((? between ctpdp.ngay_den  and ctpdp.ngay_tra_phong)
	// 						or (? between ctpdp.ngay_den and ctpdp.ngay_tra_phong)
	// 							or (? <= ctpdp.ngay_den
	// 								and ? >= ctpdp.ngay_tra_phong)
	// 								) where ctpdp.id_chi_tiet_phieu_dat_phong  is null and lp.hang_phong = ? and lp.ten_loai_phong = ?`,
	// 	checkinDate, checkoutDate, checkinDate, checkoutDate, hangPhong, loaiPhong).Scan(&phong).Error; err != nil {
	// 	return nil
	// }

	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	var phong []model.DanhSachPhong
	if hangPhong != "all" {
		tx = tx.Where("lp.hang_phong = ?", hangPhong)
	}
	if loaiPhong != "all" {
		tx = tx.Where("lp.ten_loai_phong = ?", loaiPhong)
	}
	// if err = tx.Debug().Select("*").Table("phong p").
	// 	Joins("inner join loai_phong lp on  p.id_loai_phong =lp.id_loai_phong").
	// 	Joins(`left join (select ctpdp.*,pdp.trang_thai  from chi_tiet_phieu_dat_phong ctpdp inner join phieu_dat_phong pdp on
	// 		ctpdp.id_phieu_dat_phong = pdp.id_phieu_dat_phong where pdp.trang_thai != 'Hủy') ctpdp on
	// 		p.id_phong = ctpdp.id_phong
	// 		and ((? between ctpdp.ngay_den and ctpdp.ngay_tra_phong)
	// 			or (? between ctpdp.ngay_den and ctpdp.ngay_tra_phong)
	// 				or (? <= ctpdp.ngay_den
	// 					and ? >= ctpdp.ngay_tra_phong)
	// 											)
	// 	where
	// 		ctpdp.id_chi_tiet_phieu_dat_phong is null`, checkinDate, checkoutDate, checkinDate, checkoutDate).
	// 	Find(&phong).Error; err != nil {
	// 	return nil
	// }

	if err = tx.Debug().Select(`p.*,
								lp.*,
								CASE
									WHEN COUNT(ctpdp.id_chi_tiet_phieu_dat_phong) > 0 THEN 'Đã đặt'
									ELSE 'Chưa đặt'
								END AS trang_thai_dat_phong`).Table("phong p").
		Joins("inner join loai_phong lp on  p.id_loai_phong = lp.id_loai_phong").
		Joins(`left join (
					select
						ctpdp.*,
						pdp.trang_thai
					from
						chi_tiet_phieu_dat_phong ctpdp
					inner join phieu_dat_phong pdp on
						ctpdp.id_phieu_dat_phong = pdp.id_phieu_dat_phong
					where
						pdp.trang_thai != 'Hủy'
				) ctpdp on
					p.id_phong = ctpdp.id_phong
					and (
				(? between ctpdp.ngay_den and ctpdp.ngay_tra_phong)
						or (? between ctpdp.ngay_den and ctpdp.ngay_tra_phong)
							or (? <= ctpdp.ngay_den
								and ? >= ctpdp.ngay_tra_phong)) `, checkinDate, checkoutDate, checkinDate, checkoutDate).Group("p.id_phong, lp.id_loai_phong").Order("p.so_phong").
		Find(&phong).Error; err != nil {
		return nil
	}

	return phong
}
