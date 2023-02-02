package web

type ProductCreateRequest struct {
	Kode string `validate:"required,max=10,min=1" json:"kode"`
	Nama string `validate:"required,max=100,min=1" json:"name"`
	Stok int    `validate:"required" json:"stok"`
}
