package models

import (
	"github.com/kokeroulis/modip/db"
	"encoding/json"
)

type CategoryAuthActions struct {
	IsVisibleToTeacher bool `json:"is_visible_to_teacher"`
	TeacherCanEdit     bool `json:"teacher_can_edit"`
}

type Category struct {
	Id          int                 `json:"id"`
	Name        string              `json:"name"`
	Parent      *Category           `json:"parent"`
	Children    []*Category         `json:"children"`
	Data        interface{}         `json:"data"`
	AuthActions CategoryAuthActions `json:"authActions"`
}

func (c *Category) Create() {
	jsonData, jsonErr := json.Marshal(c.Data)
	authActionsData, authActionErr := json.Marshal(c.AuthActions)

	if jsonErr != nil || authActionErr != nil {
		panic(jsonErr)
	}

	var parentId int = -1
	if c.Parent != nil {
		parentId = c.Parent.Id
	}

	var alreadyExists bool
	query := `SELECT alreadyExists
			  FROM category_add($1, $2, $3, $4, $5)`

	err := Db.Database.QueryRow(query, c.Id, c.Name, parentId, jsonData, authActionsData).
		Scan(&alreadyExists)

	Db.CheckQuery(err, query)

	if alreadyExists {
		panic("The Category already exists!!")
	}

	for _, child := range c.Children {
		child.Create()
	}
}

func (c *Category) AddChild(child *Category) {
	child.Parent = c
	child.AuthActions = c.AuthActions
	c.Children = append(c.Children, child)
}

func (c *Category) FindChildById(childId int) *Category {
	for _, it := range c.Children {
		if it.Id == childId {
			return it
		}
	}

	panic("Child doesn't exist")
	return nil
}

func (c *Category) Save() {

}

func (c *Category) Load() {
	var parentId int
	var query string

	query = `SELECT name, parent, data
			 FROM category WHERE id = $1`

	err := Db.Database.QueryRow(query, c.Id).
		Scan(&c.Name, &parentId, &c.Data)

	Db.CheckQuery(err, query)

	if parentId != 0 {
		panic("Only the parent categories should use this func")
	}

	query = `SELECT id, name, data
			 FROM category WHERE parent = $1`


	 rows, err := Db.Database.Query(query, c.Id)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		it := &Category{}

			if err := rows.Scan(&it.Id, &it.Name, &it.Data); err != nil {
				panic(err)
			} else {
				c.AddChild(it)
			}
		}

	if rowsErr := rows.Err(); rowsErr != nil {
		panic(err)
	}
}

