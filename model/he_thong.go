package model

type HeThong struct {
	IDDangNhap  int    `json:"id_dang_nhap" gorm:"primaryKey"`
	TenDangNhap string `json:"ten_dang_nhap"`
	MatKhau     string `json:"mat_khau"`
	RoleID      int    `json:"role_id"`
}

func (HeThong) TableName() string {
	return "he_thong"
}

type Quyen struct {
	RoleID   int    `json:"role_id"  gorm:"primaryKey"`
	TenQuyen string `json:"ten_quyen"`
}

func (Quyen) TableName() string {
	return "quyen"
}

type DangNhap struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DangKy struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type UserModel struct {
	ID          int    `json:"id"`
	TenDangNhap string `json:"ten_dang_nhap"`
	// Password    string `json:"password"`
	Name string `json:"name"`
	// LastName    string `json:"lastName"`
	Gender      bool   `json:"gender"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	RoleId      int64  `json:"roleId"`
	NgaySinh    string `json:"ngay_sinh"`
	Image       string `json:"image"`
}
