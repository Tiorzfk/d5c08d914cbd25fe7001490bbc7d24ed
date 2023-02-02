package helper

import (
	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/model/domain"
	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/model/web"
)

func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Kode: product.Kode,
		Nama: product.Nama,
		Stok: product.Stok,
	}
}

func ToProductCartResponse(product domain.ProductCart) web.ProductCartResponse {
	return web.ProductCartResponse{
		KodeProduk: product.KodeProduk,
		Kuantitas:  product.Kuantitas,
	}
}
