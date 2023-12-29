package model

type NhanVien struct {
	IdNhanVien  int    `json:"id_nhan_vien"`
	TenNhanVien string `json:"ten_nhan_vien"`
	GioiTinh    string `json:"gioi_tinh"`
	Sdt         string `json:"sdt"`
	Cmnd        string `json:"so_cmnd"`
	ChucVu      string `json:"chuc_vu"`
	DiaChi      string `json:"dia_chi"`
	NgayVaoLam  string `json:"ngay_vao_lam"`
}

func (NhanVien) TableName() string {
	return "nhan_vien"
}
