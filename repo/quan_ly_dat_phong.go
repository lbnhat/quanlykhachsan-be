package repo

import (
	"context"
	"quanlykhachsan/model"
	"time"

	"gorm.io/gorm"
)

// FindAll implements TagsRepository
func (r *RepoPG) DanhSachPhieuDatPhong(idKhach int, trangThai string) []model.DanhSachPhieuDatPhong {

	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	var phieudatphong []model.DanhSachPhieuDatPhong
	if idKhach != 0 {
		tx = tx.Where("id_khach_hang = ?", idKhach)
	}
	if trangThai != "" {
		tx = tx.Where("trang_thai = ?", trangThai)
	}
	if err = tx.Debug().Select("*").Table("phieu_dat_phong pdp").
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
		}).Order("id_phieu_dat_phong desc").
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

func (r *RepoPG) LayPhieuDatPhong(IdPhieuDatPhong int) (model.PhieuDatPhong, error) {
	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	var phieudatphong model.PhieuDatPhong
	if err = tx.Where("id_phieu_dat_phong = ?", IdPhieuDatPhong).Find(&phieudatphong).Error; err != nil {
		return phieudatphong, err
	}
	return phieudatphong, nil
}

func (r *RepoPG) CapNhatPhieuDatPhong(req *model.PhieuDatPhong) error {
	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	if err = tx.Updates(&req).Error; err != nil {
		return err
	}
	return nil
}

func (r *RepoPG) BaoCaoTheoThang() []model.BaoCaoTheoNgay {
	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	now := time.Now()

	// Ngày đầu tháng
	firstDay := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	// Ngày cuối cùng của tháng (lấy ngày đầu tháng của tháng tiếp theo, rồi trừ đi 1 ngày)
	lastDay := firstDay.AddDate(0, 1, 0).Add(-time.Second)
	var baoCaoTheoNgay []model.BaoCaoTheoNgay
	if err = tx.Raw(`WITH all_days AS (
						SELECT generate_series(
							?::date, 
							?::date, 
							'1 day'::interval
						)::date AS ngay
					)
					SELECT 
						ad.ngay::text,
						COALESCE(SUM(hd.tong_tien), 0) AS tong_tien
					FROM 
						all_days ad
					LEFT JOIN 
						hoa_don hd ON ad.ngay = DATE_TRUNC('day', hd.ngay_lap_phieu)::date
					GROUP BY 
						ad.ngay
					ORDER BY 
						ad.ngay;`,
		firstDay, lastDay).Scan(&baoCaoTheoNgay).Error; err != nil {
		return nil
	}
	return baoCaoTheoNgay
}

func (r *RepoPG) BaoCaoTheoNgay() []model.BaoCaoTheoNgay {
	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	var baoCaoTheoNgay []model.BaoCaoTheoNgay
	now := time.Now()

	// Lấy ngày hiện tại
	currentDate := now.Format("2006-01-02")
	if err = tx.Raw(`SELECT 
						(DATE_TRUNC('day', ngay_lap_phieu)::date)::text ngay,
						SUM(tong_tien) AS tong_tien
						FROM 
						hoa_don
						WHERE 
						ngay_lap_phieu::date = ?
						GROUP BY 
						DATE_TRUNC('day', ngay_lap_phieu)::date;`, currentDate,
	).Scan(&baoCaoTheoNgay).Error; err != nil {
		return nil
	}
	return baoCaoTheoNgay
}

func (r *RepoPG) BaoCaoTheoTuan() []model.BaoCaoTheoNgay {
	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	var baoCaoTheoNgay []model.BaoCaoTheoNgay
	if err = tx.Raw(`WITH all_weekdays AS (
						SELECT generate_series(
							date_trunc('week', CURRENT_DATE)::date, 
							date_trunc('week', CURRENT_DATE + INTERVAL '6 days')::date, 
							'1 day'::interval
						)::date AS ngay
					)
					SELECT 
						ad.ngay::text,
						COALESCE(SUM(hd.tong_tien), 0) AS tong_tien
					FROM 
						all_weekdays ad
					LEFT JOIN 
						hoa_don hd ON ad.ngay = hd.ngay_lap_phieu::date
					GROUP BY 
						ad.ngay
					ORDER BY 
						ad.ngay;`,
	).Scan(&baoCaoTheoNgay).Error; err != nil {
		return nil
	}
	return baoCaoTheoNgay
}

func (r *RepoPG) BaoCaoTheoNam() []model.BaoCaoTheoNgay {
	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	// Get the current date
	now := time.Now()
	// First day of the year
	firstDayOfYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
	lastMonth := time.Date(now.Year(), 12, 1, 0, 0, 0, 0, now.Location())
	var baoCaoTheoNgay []model.BaoCaoTheoNgay
	if err = tx.Raw(`WITH all_months AS (
						SELECT generate_series(
							?::date, 
							?::date, 
							'1 month'::interval
						)::date AS ngay
					)
					SELECT 
						to_char(am.ngay, 'YYYY-MM') AS ngay,
						COALESCE(SUM(hd.tong_tien), 0) AS tong_tien
					FROM 
						all_months am
					LEFT JOIN 
						hoa_don hd ON am.ngay <= hd.ngay_lap_phieu::date 
								AND hd.ngay_lap_phieu::date < (am.ngay + INTERVAL '1 month')
					GROUP BY 
						ngay
					ORDER BY 
						ngay;`,
		firstDayOfYear.Format("2006-01-02"), lastMonth.Format("2006-01-02")).Scan(&baoCaoTheoNgay).Error; err != nil {
		return nil
	}
	return baoCaoTheoNgay
}
