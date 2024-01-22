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

type ThongTinHoaDon struct {
	IdHoaDon          int    `json:"id_hoa_don"`
	IDPhieuDatPhong   string `json:"id_phieu_dat_phong"`
	IdKhachHang       string `json:"id_khach_hang"`
	IdNhanVien        string `json:"id_nhan_vien"`
	IdPhong           string `json:"id_phong"`
	IdDichVu          string `json:"id_dich_vu"`
	ThongTinKhachHang struct {
		IdKhachHang  string `json:"id_khach_hang"`
		TenKhachHang string `json:"ten_khach_hang"`
		Sdt          string `json:"sdt"`
	} `json:"thong_tin_khach_hang" gorm:"foreignKey:id_khach_hang;references:id_khach_hang"`
	ThongTinPhong []struct {
		IdPhieuDatPhong string `json:"id_phieu_dat_phong"`
		IdPhong         string `json:"id_phong"`
		SoPhong         string `json:"so_phong"`
		Tang            string `json:"tang"`
		TenLoaiPhong    string `json:"ten_loai_phong"`
		HangPhong       string `json:"hang_phong"`
		Gia             int    `json:"gia"`
		NgayDen         string `json:"ngay_den"`
		NgayTraPhong    string `json:"ngay_tra_phong"`
	} `json:"thong_tin_phong" gorm:"foreignKey:id_phieu_dat_phong;references:id_phieu_dat_phong"`
	ThongTinDichVu []struct {
		IdPhieuDatPhong string `json:"id_phieu_dat_phong"`
		IdDichVu        string `json:"id_dich_vu"`
		TenDichVu       string `json:"ten_dich_vu"`
		SoLuong         int    `json:"so_luong"`
		GiaDichVu       int    `json:"gia_dich_vu"`
	} `json:"thong_tin_dich_vu"  gorm:"foreignKey:id_phieu_dat_phong;references:id_phieu_dat_phong"`
	ThongTinNhanVien struct {
		IdNhanVien  string `json:"id_nhan_vien"`
		TenNhanVien string `json:"ten_nhan_vien"`
		Sdt         string `json:"sdt"`
	} `json:"thong_tin_nhan_vien" gorm:"foreignKey:id_nhan_vien;references:id_nhan_vien"`
	GiaPhong     int    `json:"gia_phong"`
	GiaDichVu    int    `json:"gia_dich_vu"`
	TongTien     int    `json:"tong_tien"`
	NgayDen      string `json:"ngay_den"`
	NgayTraPhong string `json:"ngay_tra_phong"`
	SoNgay       int    `json:"so_ngay"`
	TrangThai    string `json:"trang_thai"`
}
