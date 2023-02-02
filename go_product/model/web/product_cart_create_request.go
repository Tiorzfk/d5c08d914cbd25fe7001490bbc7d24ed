package web

type ProductCartCreateRequest struct {
	KodeProduk string `validate:"required,max=10,min=1" json:"kode_produk"`
	Kuantitas  int    `validate:"required,max=100,min=1" json:"kuantitas"`
}
