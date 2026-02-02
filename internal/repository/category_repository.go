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

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepo{db: db}
}

func (r *categoryRepo) FindAll() ([]model.Category, error) {
	var items []model.Category
	query := "SELECT id, name, description FROM categories"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var item model.Category
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
	var item model.Category
	query := "SELECT id, name, description FROM categories WHERE id = $1"

	err := r.db.QueryRow(query, id).Scan(&item.ID, &item.Name, &item.Description)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *categoryRepo) FindName(name string) (*model.Category, error) {
	var item model.Category
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
	var item model.Category
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := "UPDATE categories SET name = $1, description = $2 WHERE id = $3 RETURNING id, name, description"
	err = tx.QueryRow(query, payload.Name, payload.Description, id).Scan(&item.ID, &item.Name, &item.Description)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *categoryRepo) Store(payload *model.Category) (*model.Category, error) {
	var item model.Category
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := "INSERT INTO categories(name, description) VALUES ($1, $2) RETURNING id, name, description"
	err = tx.QueryRow(query, payload.Name, payload.Description).Scan(&item.ID, &item.Name, &item.Description)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &item, nil
}
