package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/helper"
	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/model/domain"
	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/model/web"
)

type ProductCartRepositoryImpl struct {
}

func NewProductCartRepository() *ProductCartRepositoryImpl {
	return &ProductCartRepositoryImpl{}
}

func (repository *ProductCartRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.ProductCart) domain.ProductCart {
	SQL := "INSERT INTO produk_cart(kode_produk,kuantitas) VALUES(?, ?)"
	result, err := tx.ExecContext(ctx, SQL, product.KodeProduk, product.Kuantitas)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = int(id)

	return product
}

func (repository *ProductCartRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.ProductCart) domain.ProductCart {
	SQL := "UPDATE produk_cart set kode_produk = ?, kuantitas = ? WHERE kode_produk = ?"
	_, err := tx.ExecContext(ctx, SQL, product.KodeProduk, product.Kuantitas, product.KodeProduk)
	helper.PanicIfError(err)

	return product
}

func (repository *ProductCartRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, kode string) (domain.ProductCart, error) {
	SQL := "SELECT id,kode_produk,kuantitas FROM produk_cart WHERE kode_produk = ?"
	rows, err := tx.QueryContext(ctx, SQL, kode)
	helper.PanicIfError(err)
	defer rows.Close()

	product := domain.ProductCart{}
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.KodeProduk, &product.Kuantitas)
		helper.PanicIfError(err)

		return product, nil
	} else {
		return product, errors.New("Cart is not found")
	}
}

func (repository *ProductCartRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.ProductCart) domain.ProductCart {
	SQL := "DELETE FROM produk_cart WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.PanicIfError(err)

	return product
}

func (repository *ProductCartRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []web.ProductCartResponse {
	SQL := "SELECT kode_produk,produk.nama as nama_produk, kuantitas FROM produk_cart INNER JOIN produk on produk.kode == produk_cart.kode_produk"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []web.ProductCartResponse

	for rows.Next() {
		product := web.ProductCartResponse{}
		err := rows.Scan(&product.KodeProduk, &product.NamaProduk, &product.Kuantitas)
		helper.PanicIfError(err)

		products = append(products, product)
	}

	return products
}
