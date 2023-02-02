package web

type ProductResponse struct {
	Kode string `json:"kode"`
	Nama string `json:"nama"`
	Stok int    `json:"stok"`
}
