package web

type ProductCartResponse struct {
	KodeProduk string `json:"kode_produk"`
	NamaProduk string `json:"nama_produk"`
	Kuantitas  int    `json:"kuantitas"`
}
