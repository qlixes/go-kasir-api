package repository

import (
	"database/sql"

	"kasir-api/internal/model"
)

type ProductRepository interface {
	FindAll() ([]model.Product, error)
	FindId(id string) (*model.Product, error)
	FindName(name string) (*model.Product, error)
	Erase(id string) error
	Edit(id string, payload *model.Product) (*model.Product, error)
	Store(payload *model.Product) (*model.Product, error)
}

type ProductRepo struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) FindAll() ([]model.Product, error) {
	var items []model.Product
	query := "SELECT id, name, price, quantity, category_id FROM products"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var item model.Product
	for rows.Next() {
		err := rows.Scan(&item.ID, &item.Name, &item.Price, &item.Quantity, &item.Category.ID)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *ProductRepo) FindId(id string) (*model.Product, error) {
	var item model.Product
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := "SELECT id, name, price, quantity, category_id FROM products WHERE id = $1"

	err = tx.QueryRow(query, id).Scan(&item.ID, &item.Name, &item.Price, &item.Quantity, &item.Category.ID)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *ProductRepo) FindName(name string) (*model.Product, error) {
	var item model.Product
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := "SELECT id, name, price, quantity, category_id FROM products WHERE name = $1"

	err = tx.QueryRow(query, name).Scan(&item.ID, &item.Name, &item.Price, &item.Quantity, &item.Category.ID)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *ProductRepo) Erase(id string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := "DELETE FROM products where id = $1"

	_, err = tx.Exec(query, id)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *ProductRepo) Edit(id string, payload *model.Product) (*model.Product, error) {
	var item model.Product
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := "UPDATE products SET name = $1, price = $2, quantity = $3, category_id = $4 WHERE id = $5 RETURNING id, name, price, quantity, category_id"
	err = tx.QueryRow(query, payload.Name, payload.Price, payload.Quantity, payload.Category.ID, id).Scan(&item.ID, &item.Name, &item.Price, &item.Quantity, &item.Category.ID)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *ProductRepo) Store(payload *model.Product) (*model.Product, error) {
	var item model.Product
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := "INSERT INTO products(name, description, price, quantity, category_id) VALUES ($1, $2) RETURNING id, name, price, quantity, category_id"
	err = tx.QueryRow(query, payload.Name, payload.Price, payload.Quantity, payload.Category.ID).Scan(&item.ID, &item.Name, &item.Price, &item.Quantity, &item.Category.ID)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &item, nil
}
