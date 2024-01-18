package model

type KhachHang struct {
	IdKhachHang  int    `json:"id_khach_hang" gorm:"primaryKey"`
	TenKhachHang string `json:"ten_khach_hang"`
	GioiTinh     string `json:"gioi_tinh"`
	Sdt          string `json:"sdt"`
	Cmnd         string `json:"so_cmnd"`
	Email        string `json:"email"`
	DiaChi       string `json:"dia_chi"`
	NgaySinh     string `json:"ngay_sinh"`
}

func (KhachHang) TableName() string {
	return "khach_hang"
}
