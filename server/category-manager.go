package main

import (
	"database/sql"
	"fmt"
)

// CategoryNotFoundError is
type CategoryNotFoundError struct {
	Name string
	ID   int64
}

func (z *CategoryNotFoundError) Error() string {
	return fmt.Sprintf("CategoryNotFoundError: ID=%d, Name=%s", z.ID, z.Name)
}

// CategoryManager manages categories.
type CategoryManager struct {
}

// NewCategoryManager news a CategoryManager.
func NewCategoryManager() *CategoryManager {
	return &CategoryManager{}
}

// GetVars gets vars.
func (z *CategoryManager) GetVars(tx Querier, fields string, wheres string, outs ...interface{}) error {
	q := make(map[string]interface{})
	q["select"] = fields
	q["from"] = "taxonomies"
	q["where"] = []string{
		wheres,
	}
	q["limit"] = 1

	query := BuildQueryString(q)
	row := tx.QueryRow(query)
	return row.Scan(outs...)
}

// GetCategoryByID gets a category by its ID.
func (z *CategoryManager) GetCategoryByID(tx Querier, id int64) (*Category, error) {
	query := `SELECT * FROM taxonomies WHERE id = ? LIMIT 1`
	row := tx.QueryRow(query, id)
	return z.scanOne(row)
}

func (z *CategoryManager) scanOne(scn RowScanner) (*Category, error) {
	var cat Category
	if err := scn.Scan(&cat.ID, &cat.Name, &cat.Slug, &cat.Parent, &cat.Ancestor); err != nil {
		if err == sql.ErrNoRows {
			return nil, &CategoryNotFoundError{}
		}
		return nil, err
	}
	return &cat, nil
}

func (z *CategoryManager) scanMulti(rows *sql.Rows) ([]*Category, error) {
	defer rows.Close()
	cats := make([]*Category, 0)
	for rows.Next() {
		cat, err := z.scanOne(rows)
		if err != nil {
			return nil, err
		}
		cats = append(cats, cat)
	}
	return cats, nil
}

// ListCategories lists categories.
func (z *CategoryManager) ListCategories(tx Querier) ([]*Category, error) {
	query := `SELECT * FROM taxonomies`
	rows, err := tx.Query(query)
	if err != nil {
		return nil, err
	}
	return z.scanMulti(rows)
}

// GetTree creates a category tree.
func (z *CategoryManager) GetTree(tx Querier) ([]*Category, error) {
	cats, err := z.ListCategories(tx)
	if err != nil {
		return nil, err
	}

	var makeChildren func(parent *Category)

	makeChildren = func(parent *Category) {
		for i, c := range cats {
			if c != nil && c.Parent == parent.ID {
				parent.Children = append(parent.Children, c)
				cats[i] = nil
				makeChildren(c)
			}
		}
	}

	var dummy Category
	makeChildren(&dummy)
	return dummy.Children, nil
}

// GetChildren gets direct descendant children.
func (z *CategoryManager) GetChildren(tx Querier, parent int64) ([]*Category, error) {
	query := `SELECT * FROM taxonomies WHERE parent=?`
	rows, err := tx.Query(query, parent)
	if err != nil {
		return nil, err
	}
	return z.scanMulti(rows)
}

func (z *CategoryManager) UpdateCategory(tx Querier, cat *Category) error {
	return nil
}

func (z *CategoryManager) CreateCategory(tx Querier, cat *Category) error {
	return nil
}
