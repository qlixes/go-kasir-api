package repository

import (
	"database/sql"

	"kasir-api/internal/model"
)

type CategoryRepository interface {
	FindAll() ([]model.Category, error)
	FindId(id string) (*model.Category, error)
	FindName(name string) (*model.Category, error)
	Erase(id string) error
	Edit(id string, payload *model.Category) (*model.Category, error)
	Store(payload *model.Category) (*model.Category, error)
}

type categoryRepo struct {
	db *sql.DB
}

var (
	items []model.Category
	item  model.Category
)

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepo{db: db}
}

func (r *categoryRepo) FindAll() ([]model.Category, error) {
	query := "SELECT id, name, description FROM categories"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&item.ID, &item.Name, &item.Description)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *categoryRepo) FindId(id string) (*model.Category, error) {
	query := "SELECT id, name, description FROM categories WHERE id = $1"

	err := r.db.QueryRow(query, id).Scan(&item.ID, &item.Name, &item.Description)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *categoryRepo) FindName(name string) (*model.Category, error) {
	query := "SELECT id, name, description FROM categories WHERE name = $1"

	err := r.db.QueryRow(query, name).Scan(&item.ID, &item.Name, &item.Description)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *categoryRepo) Erase(id string) error {
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

func (r *categoryRepo) Edit(id string, payload *model.Category) (*model.Category, error) {
	query := "UPDATE categories SET name = $1, description = $2 WHERE id = $3"
	err := r.db.QueryRow(query, payload.Name, payload.Description, id).Scan(&item.ID, &item.Name, &item.Description)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *categoryRepo) Store(payload *model.Category) (*model.Category, error) {
	query := "INSERT INTO categories(name, description) VALUES ($1, $2)"
	err := r.db.QueryRow(query, payload.Name, payload.Description).Scan(&item.ID, &item.Name, &item.Description)
	if err != nil {
		return nil, err
	}

	return &item, nil
}
