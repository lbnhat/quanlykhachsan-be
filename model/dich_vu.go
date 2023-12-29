package model

type DichVu struct {
	IdDichVu  int    `json:"id_dich_vu"`
	TenDichVu string `json:"ten_dich_vu"`
	GiaDichVu string `json:"gia_dich_vu"`
}

func (DichVu) TableName() string {
	return "dich_vu"
}

type ChiTietDichVu struct {
	IdChiTietDichVu int    `json:"id_chi_tiet_dich_vu"`
	IdDichVu        string `json:"id_dich_vu"`
	IdPhieuDichVu   string `json:"id_phieu_dich_vu"`
	SoLuong         string `json:"so_luong"`
	GiaDichVu       string `json:"gia_dich_vu"`
}

func (ChiTietDichVu) TableName() string {
	return "chi_tiet_dich_vu"
}

type DanhSachDichVu struct {
	IdDichVu        int    `json:"id_dich_vu"`
	TenDichVu       string `json:"ten_dich_vu"`
	GiaDichVu       string `json:"gia_dich_vu"`
	IdChiTietDichVu int    `json:"id_chi_tiet_dich_vu"`
	IdPhieuDichVu   string `json:"id_phieu_dich_vu"`
	SoLuong         string `json:"so_luong"`
}
