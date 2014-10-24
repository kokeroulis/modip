package categories

import "github.com/kokeroulis/modip/models"

var categoriesList []*models.Category

func newParentCategoryForTeacher(id int, name string) *models.Category {
	c := newCategory(id, name)

	c.AuthActions = models.CategoryAuthActions{
		IsVisibleToTeacher: true,
		TeacherCanEdit: true,
	}

	categoriesList = append(categoriesList, c)

	return c
}

func newCategory(id int, name string) *models.Category {
	c := &models.Category{
		Id: id,
		Name: name,
	}

	return c
}

func CreateCategories() {
	init1()
	init2()

	for _, c := range categoriesList {
		c.Create()
		c.Load()
	}
}

