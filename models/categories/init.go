package categories

import "github.com/kokeroulis/modip/models"

var groupList []models.CategoryGroup

func newParentCategoryForTeacher(id int, name string, group models.CategoryGroup) *models.Category {
	c := newCategory(id, name)

	c.AuthActions = models.CategoryAuthActions{
		IsVisibleToTeacher: true,
		TeacherCanEdit: true,
	}

	group.AddCategory(c)
	groupList = append(groupList, group)

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

	for _, g := range groupList {
		g.Create()
	}
}

