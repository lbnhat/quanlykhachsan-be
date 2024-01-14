package repo

import (
	"context"
	"quanlykhachsan/model"
	"time"

	"gorm.io/gorm"
)

const (
	StateNew byte = iota + 1 // starts from 1
	StateDoing
	StateDone

	generalQueryTimeout = 60 * time.Second
	defaultPageSize     = 30
	maxPageSize         = 1000
)

func NewPGRepo(db *gorm.DB) PGInterface {
	return &RepoPG{DB: db}
}

type PGInterface interface {

	// DB
	GetRepo() *gorm.DB
	DBWithTimeout(ctx context.Context) (*gorm.DB, context.CancelFunc)

	//KhachHang
	DanhSachKhachHang() []model.KhachHang
	TaoKhachHang(req *model.KhachHang) error
	TimKhachHang(id int, req *model.KhachHang) error

	//NhanVien
	DanhSachNhanVien() []model.NhanVien

	//Phong
	DanhSachPhong() []model.DanhSachPhong
	DanhSachPhongTrong(heckinDate, checkoutDate, hangPhong, loaiPhong string) []model.DanhSachPhong // Dịch vụ
	DanhSachDichVu() []model.DanhSachDichVu

	DanhSachPhieuDatPhong(idKhach int) []model.DanhSachPhieuDatPhong
	TaoPhieuDatPhong(req *model.PhieuDatPhong) error
	TaoChiTietPhieuDatPhong(req *model.ChiTietPhieuDatPhong) error

	DanhSachHoaDon() []model.HoaDon

	TaoPhieuDichvu(req *model.PhieuDichVu) error
	TaoChiTietDichVu(req *model.ChiTietDichVu) error
}

type RepoPG struct {
	DB    *gorm.DB
	debug bool
}

func (r *RepoPG) GetRepo() *gorm.DB {
	return r.DB
}

func (r *RepoPG) DBWithTimeout(ctx context.Context) (*gorm.DB, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(ctx, generalQueryTimeout)
	return r.DB.WithContext(ctx), cancel
}
