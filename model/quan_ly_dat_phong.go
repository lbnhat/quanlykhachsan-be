package model

type PhieuDatPhong struct {
	IdPhieuDatPhong int    `json:"id_phieu_dat_phong" gorm:"primaryKey"`
	IdNhanVien      int    `json:"id_nhan_vien"`
	IdKhachHang     int    `json:"id_khach_hang"`
	NgayLapPhieu    string `json:"ngay_lap_phieu"`
	NgayDen         string `json:"ngay_den"`
	TrangThai       string `json:"trang_thai"`
	TongTien        int    `json:"tong_tien"`
}

func (PhieuDatPhong) TableName() string {
	return "phieu_dat_phong"
}

type ChiTietPhieuDatPhong struct {
	IdChiTietPhieuDatPhong int    `json:"id_chi_tiet_phieu_dat_phong" gorm:"primaryKey"`
	IdPhieuDatPhong        int    `json:"id_phieu_dat_phong"`
	IdPhong                int    `json:"id_phong"`
	NgayDen                string `json:"ngay_den"`
	NgayTraPhong           string `json:"ngay_tra_phong"`
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
	GiaPhong  int `json:"gia_phong"`
	GiaDichVu int `json:"gia_dich_vu"`
	TongTien  int `json:"tong_tien"`
	// NgayDen   string `json:"ngay_den"`
	// NgayDi    string `json:"ngay_di"`
	ThoiGian  string `json:"thoi_gian"`
	TrangThai string `json:"trang_thai"`
}

type DatPhongRequest struct {
	UserID        int    `json:"user_id"`
	Email         string `json:"email"`
	Sdt           string `json:"sdt"`
	TenKhachHang  string `json:"ten_khach_hang"`
	Cmnd          string `json:"cmnd"`
	CheckinDate   string `json:"checkinDate"`
	CheckoutDate  string `json:"checkoutDate"`
	DanhSachPhong []struct {
		IDPhong     int    `json:"id_phong"`
		IDLoaiPhong string `json:"id_loai_phong"`
		SoTang      string `json:"so_tang"`
		SoPhong     string `json:"so_phong"`
		TrangThai   bool   `json:"trang_thai"`
		HinhAnh     string `json:"hinh_anh"`
		LoaiPhong   string `json:"loai_phong"`
		HangPhong   string `json:"hang_phong"`
		GiaPhong    string `json:"gia_phong"`
	} `json:"danh_sach_phong"`
	DanhSachDichVu []struct {
		IDDichVu        int    `json:"id_dich_vu"`
		TenDichVu       string `json:"ten_dich_vu"`
		GiaDichVu       int    `json:"gia_dich_vu"`
		IDChiTietDichVu int    `json:"id_chi_tiet_dich_vu"`
		IDPhieuDichVu   string `json:"id_phieu_dich_vu"`
		SoLuong         int    `json:"so_luong"`
	} `json:"danh_sach_dich_vu"`
	TongTien int `json:"tong_tien"`
}

type CapNhatPhong struct {
	IdPhieuDatPhong int    `json:"id_phieu_dat_phong"`
	TrangThai       string `json:"trang_thai"`
}
