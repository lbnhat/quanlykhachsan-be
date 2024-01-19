package model

type Phong struct {
	IdPhong     int    `json:"id_phong"`
	IdLoaiPhong string `json:"id_loai_phong"`
	SoTang      string `json:"so_tang"`
	SoPhong     string `json:"so_phong"`
	TrangThai   string `json:"trang_thai"`
	HinhAnh     string `json:"hinh_anh"`
}

func (Phong) TableName() string {
	return "phong"
}

type LoaiPhong struct {
	IdLoaiPhong  int    `json:"id_loai_phong"`
	TenLoaiPhong string `json:"loai_phong"`
	HangPhong    string `json:"hang_phong"`
	Gia          string `json:"gia_phong"`
}

func (LoaiPhong) TableName() string {
	return "loai_phong"
}

type DanhSachPhong struct {
	IdPhong           int    `json:"id_phong"`
	IdLoaiPhong       string `json:"id_loai_phong"`
	SoTang            string `json:"so_tang"`
	SoPhong           string `json:"so_phong"`
	TrangThaiDatPhong string `json:"trang_thai_dat_phong"`
	HinhAnh           string `json:"hinh_anh"`
	TenLoaiPhong      string `json:"loai_phong"`
	HangPhong         string `json:"hang_phong"`
	Gia               string `json:"gia_phong"`
}
