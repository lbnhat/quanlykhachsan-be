package model

type HoaDon struct {
	IdHoaDon        int    `json:"id_hoa_don" gorm:"primaryKey"`
	IdNhanVien      int    `json:"id_nhan_vien"`
	IdPhieuDatPhong int    `json:"id_phieu_dat_phong"`
	NgayLapPhieu    string `json:"ngay_lap_phieu"`
	TongTien        int    `json:"tong_tien"`
}

func (HoaDon) TableName() string {
	return "hoa_don"
}

type DanhSachHoaDon struct {
	IdHoaDon        int    `json:"id_hoa_don"`
	IdNhanVien      string `json:"id_nhan_vien"`
	IdPhieuDatPhong string `json:"id_phieu_dat_phong"`
	NgayLapPhieu    string `json:"ngay_lap_phieu"`
	TongTien        string `json:"tong_tien"`
	NgayTao         string `json:"ngay_tao"`
	TenKhachHang    string `json:"ten_khach_hang"`
	TenNhanVien     string `json:"ten_nhan_vien"`
	Sdt             string `json:"sdt"`
	NguoiLapPhieu   string `json:"nguoi_lap_phieu"`
}
