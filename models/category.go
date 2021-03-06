package models

import (
	"github.com/kokeroulis/modip/db"
	"encoding/json"
)

type CategoryGroup struct {
	Id         int         `json:"id"`
	Name       string      `json:"name"`
	Categories []*Category `json:"categories"`
}

func (g *CategoryGroup) AddCategory(c *Category) {
	g.Categories = append(g.Categories, c)
}

func (g *CategoryGroup) Create() {
	var alreadyExists bool
	query := `SELECT alreadyExists
			  FROM category_group_add($1, $2)`

	for _, category := range g.Categories {
		category.Create()

		err := Db.Database.QueryRow(query, category.Id, g.Name).
		Scan(&alreadyExists)

		Db.CheckQuery(err, query)

		if alreadyExists {
			panic("The Category is already in group!!")
		}
	}
}

func (g *CategoryGroup) Load() {
	query := `SELECT name, category
			  FROM categorygroup`

	rows, err := Db.Database.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var categoryId int

		if err := rows.Scan(&g.Name, &categoryId); err != nil {
			panic(err)
		} else {
			c := &Category{
				Id: categoryId,
			}
			c.Load()
			g.AddCategory(c)
		}
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		panic(err)
	}
}

type CategoryAuthActions struct {
	IsVisibleToTeacher bool `json:"is_visible_to_teacher"`
	TeacherCanEdit     bool `json:"teacher_can_edit"`
}

type Category struct {
	Id          int                 `json:"id"`
	Name        string              `json:"name"`
	Parent      *Category           `json:"-"`
	Children    []*Category         `json:"children"`
	Data        string              `json:"data,omitempty"`
	AuthActions CategoryAuthActions `json:"authActions"`
}

func (c *Category) Create() {
	authActionsData, authActionErr := json.Marshal(c.AuthActions)

	if authActionErr != nil {
		panic(authActionErr)
	}

	var parentId int = -1
	if c.Parent != nil {
		parentId = c.Parent.Id
	}

	var alreadyExists bool
	query := `SELECT alreadyExists
			  FROM category_add($1, $2, $3, $4)`

	err := Db.Database.QueryRow(query, c.Id, c.Name, parentId, authActionsData).
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

func (c *Category) Save(teacherId int) {
	var notExists bool
	query := `SELECT notExists
			  FROM category_save($1, $2, $3)`

	err := Db.Database.QueryRow(query, c.Id, c.Data, teacherId).
	Scan(&notExists)

	Db.CheckQuery(err, query)

	if notExists {
		panic("The Category is already in group!!")
	}
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

	query = `SELECT id, name, cd.data
			 FROM category AS c
			 LEFT JOIN CategoryData AS cd
			 ON cd.category = c.id
			 WHERE parent = $1
			 ORDER BY id`

	rows, err := Db.Database.Query(query, c.Id)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		it := &Category{}
		var dataInterface interface{}
		if err := rows.Scan(&it.Id, &it.Name, &dataInterface); err != nil {
			panic(err)
		} else {
			data, ok:= dataInterface.(string)
			if ok {
				it.Data = data
			}

			c.AddChild(it)
		}
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		panic(err)
	}
}

func ListAllCategories() []CategoryGroup {
	var groupList []CategoryGroup

	g := CategoryGroup {
		Id: 1,
	}

	g.Load()

	groupList = append(groupList, g)

	return groupList
}

