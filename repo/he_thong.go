package repo

import (
	"context"
	"fmt"
	"quanlykhachsan/model"
)

func (r *RepoPG) LayThongTinDangNhap(TenDangNhap, MatKhau string) *model.HeThong {
	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	//var err error
	heThong := &model.HeThong{}
	if err := tx.Debug().Where("ten_dang_nhap = ? and mat_khau = ?", TenDangNhap, MatKhau).First(heThong).Error; err != nil {
		fmt.Println(err)
		return nil
	}
	return heThong
}

func (r *RepoPG) KiemTraThongTinDangKy(TenDangNhap string) *model.HeThong {
	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	//var err error
	heThong := &model.HeThong{}
	if err := tx.Debug().Where("ten_dang_nhap = ?", TenDangNhap).First(heThong).Error; err != nil {
		fmt.Println(err)
		return nil
	}
	return heThong
}

func (r *RepoPG) TaoTaiKhoan(req *model.HeThong) error {
	tx, cancel := r.DBWithTimeout(context.Background())
	defer cancel()
	var err error
	if err = tx.Create(&req).Error; err != nil {
		return err
	}
	return nil
}
