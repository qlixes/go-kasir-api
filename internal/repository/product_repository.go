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
	query := "SELECT id, name, description FROM categories"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []model.Product
	var item model.Product

	for rows.Next() {
		err := rows.Scan(&item.ID, &item.Name, &item.Description)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *ProductRepo) FindId(id string) (*model.Product, error) {
	var item model.Product
	query := "SELECT id, name, description FROM categories WHERE id = $1"

	err := r.db.QueryRow(query, id).Scan(&item.ID, &item.Name, &item.Description)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *ProductRepo) FindName(name string) (*model.Product, error) {
	var item model.Product
	query := "SELECT id, name, description FROM categories WHERE name = $1"

	err := r.db.QueryRow(query, name).Scan(&item.ID, &item.Name, &item.Description)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *ProductRepo) Erase(id string) error {
	query := "DELETE FROM categories where id = $1"

	row, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	if _, err := row.RowsAffected(); err != nil {
		return err
	}

	return nil
}

func (r *ProductRepo) Edit(id string, payload *model.Product) (*model.Product, error) {
	var item model.Product
	query := "UPDATE categories SET name = $1, description = $2 WHERE id = $3 RETURNING id, name, description"

	err := r.db.QueryRow(query, payload.Name, payload.Description, id).Scan(&item.ID, &item.Name, &item.Description)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *ProductRepo) Store(payload *model.Product) (*model.Product, error) {
	var item model.Product
	query := "INSERT INTO categories(name, description) VALUES ($1, $2) RETURNING id, name, description"
	err := r.db.QueryRow(query, payload.Name, payload.Description).Scan(&item.ID, &item.Name, &item.Description)
	if err != nil {
		return nil, err
	}

	return &item, nil
}
