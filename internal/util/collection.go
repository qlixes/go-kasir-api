package util

import "database/sql"

func ScanAll[T any](rows *sql.Rows, scanFn func(*sql.Rows) (T, error)) ([]T, error) {
	defer rows.Close()
	var results []T

	for rows.Next() {
		item, err := scanFn(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
