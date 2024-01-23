package service

import (
	"fmt"
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
	DanhSachPhieuDatPhong(trangThai string) []model.DanhSachPhieuDatPhong
	DanhSachPhieuDatPhongTuKhach(idKhach int) []model.DanhSachPhieuDatPhong
	DatPhong(req model.DatPhongRequest) (err error)
	CapNhatPhong(req model.CapNhatPhong) (err error)
	BaoCao(option string) (rs model.BaoCao)
}

func (t *DatPhongService) DanhSachPhieuDatPhong(trangThai string) []model.DanhSachPhieuDatPhong {
	result := t.repo.DanhSachPhieuDatPhong(0, trangThai)
	return result
}

func (t *DatPhongService) DanhSachPhieuDatPhongTuKhach(idKhach int) []model.DanhSachPhieuDatPhong {
	result := t.repo.DanhSachPhieuDatPhong(idKhach, "")
	return result
}

func (t *DatPhongService) DatPhong(req model.DatPhongRequest) (err error) {
	kh := model.KhachHang{
		TenKhachHang: req.TenKhachHang,
		Sdt:          req.Sdt,
		Cmnd:         req.Cmnd,
		Email:        req.Email,
	}
	if len(req.DanhSachPhong) == 0 {
		return fmt.Errorf("Thiếu thông tin phòng")
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
		if v.SoLuong != 0 {
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
	}

	return nil
}

func (t *DatPhongService) CapNhatPhong(req model.CapNhatPhong) (err error) {

	pdp, err := t.repo.LayPhieuDatPhong(req.IdPhieuDatPhong)
	pdp.TrangThai = req.TrangThai
	err = t.repo.CapNhatPhieuDatPhong(&pdp)
	if err != nil {
		return err
	}

	if req.TrangThai == "Đã thanh toán" {
		hd := model.HoaDon{
			IdPhieuDatPhong: pdp.IdPhieuDatPhong,
			NgayLapPhieu:    time.Now().Format("2006-01-02"),
			TongTien:        pdp.TongTien,
			IdNhanVien:      222,
		}
		err = t.repo.TaoHoaDon(&hd)
		if err != nil {
			return err
		}
	}
	return nil
}

// CREATE SEQUENCE phieu_dat_phong_id_phieu_dat_phong_seq START 1;

// ALTER TABLE public.phieu_dat_phong
// ALTER COLUMN id_phieu_dat_phong SET DEFAULT nextval('phieu_dat_phong_id_phieu_dat_phong_seq'::regclass);

func (t *DatPhongService) BaoCao(option string) (rs model.BaoCao) {
	switch option {
	case "ngay":
		result := t.repo.BaoCaoTheoNgay()
		if len(result) == 0 {
			now := time.Now()
			result = append(result, model.BaoCaoTheoNgay{
				Ngay:     now.Format("2006-01-02"),
				TongTien: 0,
			})
		}
		tongTien := 0
		for _, v := range result {
			tongTien += v.TongTien
		}
		return model.BaoCao{
			TongTien: tongTien,
			BieuDo:   result,
		}
	case "tuan":
		result := t.repo.BaoCaoTheoTuan()
		tongTien := 0
		for _, v := range result {
			tongTien += v.TongTien
		}
		return model.BaoCao{
			TongTien: tongTien,
			BieuDo:   result,
		}
	case "nam":
		result := t.repo.BaoCaoTheoNam()
		tongTien := 0
		for _, v := range result {
			tongTien += v.TongTien
		}
		return model.BaoCao{
			TongTien: tongTien,
			BieuDo:   result,
		}
	default:
		result := t.repo.BaoCaoTheoThang()
		tongTien := 0
		for _, v := range result {
			tongTien += v.TongTien
		}
		return model.BaoCao{
			TongTien: tongTien,
			BieuDo:   result,
		}
	}
}
