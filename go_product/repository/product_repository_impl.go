package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/helper"
	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/model/domain"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() *ProductRepositoryImpl {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "INSERT INTO produk(kode,nama,strok) VALUES(?,?,?)"
	_, err := tx.ExecContext(ctx, SQL, product.Kode, product.Nama, product.Stok)
	helper.PanicIfError(err)

	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "UPDATE produk set kode = ?, nama = ?, stok = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Kode, product.Nama, product.Stok, product.Kode)
	helper.PanicIfError(err)

	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "DELETE FROM produk WHERE kode = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Kode)
	helper.PanicIfError(err)

	return product
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, kode string) (domain.Product, error) {
	SQL := "SELECT kode,nama,stok FROM product WHERE kode = ?"
	rows, err := tx.QueryContext(ctx, SQL, kode)
	helper.PanicIfError(err)
	defer rows.Close()

	product := domain.Product{}
	if rows.Next() {
		err := rows.Scan(&product.Kode, &product.Nama, &product.Stok)
		helper.PanicIfError(err)

		return product, nil
	} else {
		return product, errors.New("Product is not found")
	}
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	SQL := "SELECT kode,nama,stok FROM produk"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []domain.Product

	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.Kode, &product.Nama, &product.Kode)
		helper.PanicIfError(err)

		products = append(products, product)
	}

	return products
}
