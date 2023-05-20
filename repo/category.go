package repo

import (
	"forum/entity"
)

type ICateg interface {
	CreateCategory(c *entity.Category) error
	GetCategory(categ string) (*entity.Category, error)
	GetCategoryByPostID(id int64) ([]string, error)
	DeleteCategoryByPostID(id int64) error
}

func (r repo) CreateCategory(c *entity.Category) error {
	stmt, err := r.db.Prepare("INSERT INTO category (post_id, name) VALUES (?, ?)")
	if err != nil {
		r.log.Printf("error while to prepare datas to write into the category table: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(c.PostId, c.Name)
	if err != nil {
		r.log.Printf("error while exec prepared datas to write into category table: %s\n", err.Error())
		return err
	}
	return nil
}

// getCategoryByID retrieves a category from the Category table by ID
func (r repo) GetCategory(categ string) (*entity.Category, error) {
	stmt, err := r.db.Prepare("SELECT id, post_id, name FROM category WHERE name = ?")
	if err != nil {
		r.log.Printf("error while to prepare datas to get comment by id from category table: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	var category *entity.Category
	err = stmt.QueryRow(categ).Scan(&category.Id, &category.PostId, &category.Name)
	if err != nil {
		r.log.Printf("error while to query row and scan category to get by id: %s\n", err.Error())
		return nil, err
	}

	return category, nil
}

// getCategories retrieves a category from the Category table by ID
func (r repo) GetCategoryByPostID(id int64) ([]string, error) {
	stmt, err := r.db.Prepare("SELECT id, post_id, name FROM category WHERE id = ?")
	if err != nil {
		r.log.Printf("error while to prepare datas to get comment by id from category table: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	var resp []string

	rows, err := stmt.Query(id)
	if err != nil {
		r.log.Printf("error while to query row and scan category to get by id: %s\n", err.Error())
		return nil, err
	}

	for rows.Next() {
		var category *entity.Category
		if err := rows.Scan(&category.Id, &category.PostId, &category.Name); err != nil {
			return nil, err
		}
		resp = append(resp, category.Name)
	}

	return resp, nil
}

// deleteCategoryByID deletes a category from the Comment table by ID
func (r repo) DeleteCategoryByPostID(id int64) error {
	stmt, err := r.db.Prepare("DELETE FROM category WHERE post_id = ?")
	if err != nil {
		r.log.Printf("error while to prepare delete category by id in category table: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		r.log.Printf("error while exec prepared delete category by id in category table: %s\n", err.Error())
		return err
	}

	return nil
}
