package service

import (
	"quanlykhachsan/model"
	"quanlykhachsan/repo"
	"time"

	"github.com/go-playground/validator/v10"
)

type DatPhongService struct {
	repo     repo.PGInterface
	Validate *validator.Validate
}

func NewDatPhongService(repo repo.PGInterface, validate *validator.Validate) DatPhongServiceInterface {
	return &DatPhongService{
		repo:     repo,
		Validate: validate,
	}
}

type DatPhongServiceInterface interface {
	DanhSachPhieuDatPhong() []model.DanhSachPhieuDatPhong
	DanhSachPhieuDatPhongTuKhach(idKhach int) []model.DanhSachPhieuDatPhong
	DatPhong(req model.DatPhongRequest) (err error)
}

func (t *DatPhongService) DanhSachPhieuDatPhong() []model.DanhSachPhieuDatPhong {
	result := t.repo.DanhSachPhieuDatPhong(0)
	return result
}

func (t *DatPhongService) DanhSachPhieuDatPhongTuKhach(idKhach int) []model.DanhSachPhieuDatPhong {
	result := t.repo.DanhSachPhieuDatPhong(idKhach)
	return result
}

func (t *DatPhongService) DatPhong(req model.DatPhongRequest) (err error) {
	kh := model.KhachHang{
		TenKhachHang: req.TenKhachHang,
		Sdt:          req.Sdt,
		Cmnd:         req.Cmnd,
		Email:        req.Email,
	}

	if req.UserID != 0 {
		err = t.repo.TimKhachHang(req.UserID, &kh)
		if err != nil {
			return err
		}
	} else {
		err = t.repo.TaoKhachHang(&kh)
		if err != nil {
			return err
		}
	}

	pdp := model.PhieuDatPhong{
		IdKhachHang:  kh.IdKhachHang,
		NgayLapPhieu: time.Now().Format("2006-01-02"),
		NgayDen:      req.CheckinDate,
		TrangThai:    "Đang chờ xác nhận",
		IdNhanVien:   223,
		TongTien:     req.TongTien,
	}
	err = t.repo.TaoPhieuDatPhong(&pdp)
	if err != nil {
		return err
	}
	for _, v := range req.DanhSachPhong {
		ctpdp := model.ChiTietPhieuDatPhong{
			IdPhieuDatPhong: pdp.IdPhieuDatPhong,
			IdPhong:         v.IDPhong,
			NgayDen:         req.CheckinDate,
			NgayTraPhong:    req.CheckoutDate,
		}
		err = t.repo.TaoChiTietPhieuDatPhong(&ctpdp)
		if err != nil {
			return err
		}
	}

	tienDichVu := 0
	for _, v := range req.DanhSachDichVu {
		tienDichVu = tienDichVu + (v.GiaDichVu * v.SoLuong)
	}
	pdv := model.PhieuDichVu{
		IdPhieuDatPhong:  pdp.IdPhieuDatPhong,
		TongTien:         tienDichVu,
		NgaySuDungDichVu: req.CheckinDate,
	}
	err = t.repo.TaoPhieuDichvu(&pdv)
	if err != nil {
		return err
	}

	for _, v := range req.DanhSachDichVu {
		ctdv := model.ChiTietDichVu{
			IdDichVu:      v.IDDichVu,
			IdPhieuDichVu: pdv.IdPhieuDichVu,
			SoLuong:       v.SoLuong,
			GiaDichVu:     v.GiaDichVu,
		}
		err = t.repo.TaoChiTietDichVu(&ctdv)
		if err != nil {
			return err
		}
	}

	return nil
}

// CREATE SEQUENCE phieu_dat_phong_id_phieu_dat_phong_seq START 1;

// ALTER TABLE public.phieu_dat_phong
// ALTER COLUMN id_phieu_dat_phong SET DEFAULT nextval('phieu_dat_phong_id_phieu_dat_phong_seq'::regclass);
