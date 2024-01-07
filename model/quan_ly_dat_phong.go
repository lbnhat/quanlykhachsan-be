package model

type PhieuDatPhong struct {
	IdPhieuDatPhong int    `json:"id_phieu_dat_phong"`
	IdNhanVien      string `json:"id_nhan_vien"`
	IdKhachHang     string `json:"id_khach_hang"`
	NgayLapPhieu    string `json:"ngay_lap_phieu"`
	NgayDen         string `json:"ngay_den"`
	TinhTrang       string `json:"tinh_trang"`
}

func (PhieuDatPhong) TableName() string {
	return "phieu_dat_phong"
}

type ChiTietPhieuDatPhong struct {
	IdChiTietPhieuDatPhong int    `json:"id_chi_tiet_phieu_dat_phong"`
	IdPhieuDatPhong        string `json:"id_phieu_dat_phong"`
	IdPhong                string `json:"id_phong"`
	NgayDen                string `json:"ngay_den"`
	NgayTra                string `json:"ngay_tra_phong"`
}

func (ChiTietPhieuDatPhong) TableName() string {
	return "chi_tiet_phieu_dat_phong"
}

// type DanhSachPhieuDatPhong struct {
// 	IdPhieuDatPhong        int    `json:"id_phieu_dat_phong"`
// 	IdNhanVien             string `json:"id_nhan_vien"`
// 	IdKhachHang            string `json:"id_khach_hang"`
// 	TenNhanVien            string `json:"ten_nhan_vien"`
// 	TenKhachHang           string `json:"ten_khach_hang"`
// 	Sdt                    string `json:"sdt"`
// 	MaDatPhong             string `json:"ma_dat_phong"`
// 	NgayLapPhieu           string `json:"ngay_lap"`
// 	NgayDen                string `json:"ngay_den"`
// 	ThoiGian               string `json:"thoi_gian"`
// 	TinhTrang              string `json:"tinh_trang"`
// 	IdChiTietPhieuDatPhong int    `json:"id_chi_tiet_phieu_dat_phong"`
// 	IdPhong                string `json:"id_phong"`
// 	NgayTra                string `json:"ngay_tra_phong"`
// 	TenDichVu              string `json:"ten_dich_vu"`
// 	GiaDichVu              string `json:"gia_dich_vu"`
// 	SoLuong                string `json:"so_luong"`
// 	GiaPhong               string `json:"gia_phong"`
// 	TongTien               string `json:"tong_tien"`
// 	TongGia                string `json:"tong_gia"`
// 	SoTang                 string `json:"so_tang"`
// 	LoaiPhong              string `json:"loai_phong"`
// 	HangPhong              string `json:"hang_phong"`
// }

type DanhSachPhieuDatPhong struct {
	IDPhieuDatPhong   string `json:"id_phieu_dat_phong"`
	IdKhachHang       string `json:"id_khach_hang"`
	IdPhong           string `json:"id_phong"`
	IdDichVu          string `json:"id_dich_vu"`
	ThongTinKhachHang struct {
		IdKhachHang  string `json:"id_khach_hang"`
		TenKhachHang string `json:"ten_khach_hang"`
		Sdt          string `json:"sdt"`
	} `json:"thong_tin_khach_hang" gorm:"foreignKey:id_khach_hang;references:id_khach_hang"`
	ThongTinPhong []struct {
		IdPhong   string `json:"id_phong"`
		MaPhong   string `json:"ma_phong"`
		Tang      string `json:"tang"`
		LoaiPhong string `json:"loai_phong"`
		HangPhong string `json:"hang_phong"`
		Gia       int    `json:"gia"`
	} `json:"thong_tin_phong" gorm:"foreignKey:id_phong;references:id_phong"`
	ThongTinDichVu []struct {
		IdDichVu  string `json:"id_dich_vu"`
		TenDichVu string `json:"ten_dich_vu"`
		SoLuong   int    `json:"so_luong"`
		Gia       int    `json:"gia"`
	} `json:"thong_tin_dich_vu"  gorm:"foreignKey:id_dich_vu;references:id_dich_vu"`
	GiaPhong  int    `json:"gia_phong"`
	GiaDichVu int    `json:"gia_dich_vu"`
	TongTien  int    `json:"tong_tien"`
	NgayDen   string `json:"ngay_den"`
	NgayDi    string `json:"ngay_di"`
	ThoiGian  string `json:"thoi_gian"`
}
