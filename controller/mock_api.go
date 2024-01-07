package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "quanlykhachsan/docs"

	"github.com/gin-gonic/gin"
)

type MockController struct {
}

func NewMockController() *MockController {
	return &MockController{}
}

func (controller *MockController) Login(ctx *gin.Context) {

	type UserModel struct {
		Username    string    `json:"username"`
		Password    string    `json:"password"`
		FirstName   string    `json:"firstName"`
		LastName    string    `json:"lastName"`
		Gender      bool      `json:"gender"`
		PhoneNumber string    `json:"phoneNumber"`
		Email       string    `json:"email"`
		RoleId      int64     `json:"roleId"`
		Birthday    time.Time `json:"birthday"`
		Image       string    `json:"image"`
	}

	user := UserModel{
		Username:    "john_doe",
		Password:    "secret123",
		FirstName:   "John",
		LastName:    "Doe",
		Gender:      true,
		PhoneNumber: "123456789",
		Email:       "john.doe@example.com",
		RoleId:      2,
		Birthday:    time.Now(),
		Image:       "profile.jpg",
	}
	// webResponse := model.Response{
	// 	Code:   http.StatusOK,
	// 	Status: "Ok",
	// 	Data:   nhanvienResponse,
	// }
	ctx.Header("Content-Type", "application/json")

	response := make(map[string]interface{})
	response["user"] = user
	jsonResponse, _ := json.Marshal(response)
	fmt.Println(string(jsonResponse))
	// webResponse := model.Response{
	// 	Code:   http.StatusOK,
	// 	Status: "Ok",
	// 	Data:   response,
	// }
	ctx.JSON(http.StatusOK, response)

}

func (controller *MockController) HotelSearch(ctx *gin.Context) {

	type Hotel struct {
		ID               string `json:"id"`
		HotelName        string `json:"hotelName"`
		HotelDescription string `json:"hotelDescription"`
		HotelAddress     string `json:"hotelAddress"`
	}

	hotel := append([]Hotel{}, Hotel{
		ID:               "1",
		HotelName:        "số 2",
		HotelDescription: "fsadsadsa",
		HotelAddress:     "đâsdasdadas",
	})
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, hotel)
}

func (controller *MockController) Room(ctx *gin.Context) {

	type Room struct {
		ID          int    `json:"id"`
		Image       string `json:"image"`
		TypeID      int    `json:"type_id"`
		BedQuantity int    `json:"bed_quantity"`
		Description string `json:"description"`
		Price       int64  `json:"price"`
	}

	hotel := append([]Room{}, Room{
		ID:          1,
		Image:       "https://bizweb.dktcdn.net/100/153/764/products/giuong-ngu-hien-dai-72t.jpg?v=1691638130990",
		TypeID:      2,
		BedQuantity: 2,
		Description: "Phòng đẹp thơm",
		Price:       100000,
	}, Room{
		ID:          1,
		Image:       "https://bizweb.dktcdn.net/100/153/764/products/giuong-ngu-hien-dai-72t.jpg?v=1691638130990",
		TypeID:      2,
		BedQuantity: 2,
		Description: "Phòng đẹp thơm",
		Price:       100000,
	}, Room{
		ID:          1,
		Image:       "https://bizweb.dktcdn.net/100/153/764/products/giuong-ngu-hien-dai-72t.jpg?v=1691638130990",
		TypeID:      2,
		BedQuantity: 2,
		Description: "Phòng đẹp thơm",
		Price:       100000,
	}, Room{
		ID:          1,
		Image:       "https://bizweb.dktcdn.net/100/153/764/products/giuong-ngu-hien-dai-72t.jpg?v=1691638130990",
		TypeID:      2,
		BedQuantity: 2,
		Description: "Phòng đẹp thơm",
		Price:       100000,
	}, Room{
		ID:          1,
		Image:       "https://bizweb.dktcdn.net/100/153/764/products/giuong-ngu-hien-dai-72t.jpg?v=1691638130990",
		TypeID:      2,
		BedQuantity: 2,
		Description: "Phòng đẹp thơm",
		Price:       100000,
	}, Room{
		ID:          1,
		Image:       "https://bizweb.dktcdn.net/100/153/764/products/giuong-ngu-hien-dai-72t.jpg?v=1691638130990",
		TypeID:      2,
		BedQuantity: 2,
		Description: "Phòng đẹp thơm",
		Price:       100000,
	})
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, hotel)
}
