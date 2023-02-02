package web

type ProductUpdateRequest struct {
	Kode string `validate:"required" json:"kode"`
	Nama string `validate:"required,max=100,min=1" json:"nama"`
	Stok int    `validate:"required" json:"stok"`
}
