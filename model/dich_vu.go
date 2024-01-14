package model

type DichVu struct {
	IdDichVu  int    `json:"id_dich_vu" gorm:"primaryKey"`
	TenDichVu string `json:"ten_dich_vu"`
	GiaDichVu string `json:"gia_dich_vu"`
}

func (DichVu) TableName() string {
	return "dich_vu"
}

type ChiTietDichVu struct {
	IdChiTietDichVu int `json:"id_chi_tiet_dich_vu" gorm:"primaryKey"`
	IdDichVu        int `json:"id_dich_vu"`
	IdPhieuDichVu   int `json:"id_phieu_dich_vu"`
	SoLuong         int `json:"so_luong"`
	GiaDichVu       int `json:"gia_dich_vu"`
}

func (ChiTietDichVu) TableName() string {
	return "chi_tiet_dich_vu"
}

type PhieuDichVu struct {
	IdPhieuDichVu    int    `json:"id_phieu_dich_vu" gorm:"primaryKey"`
	IdPhieuDatPhong  int    `json:"id_phieu_dat_phong"`
	TongTien         int    `json:"tong_tien"`
	NgaySuDungDichVu string `json:"ngay_su_dung_dich_vu"`
}

func (PhieuDichVu) TableName() string {
	return "phieu_dich_vu"
}

type DanhSachDichVu struct {
	IdDichVu        int    `json:"id_dich_vu"`
	TenDichVu       string `json:"ten_dich_vu"`
	GiaDichVu       int    `json:"gia_dich_vu"`
	IdChiTietDichVu int    `json:"id_chi_tiet_dich_vu"`
	IdPhieuDichVu   string `json:"id_phieu_dich_vu"`
	SoLuong         int    `json:"so_luong"`
}
