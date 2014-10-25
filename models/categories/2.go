package categories

import "github.com/kokeroulis/modip/models"

func init2() {
	g := models.CategoryGroup {
		Name: "Group2",
	}

	parent := newParentCategoryForTeacher(2, "bar", g)
	parent.AddChild(newCategory(2001, "barChild"))
	parent.AddChild(newCategory(2002, "barChild2"))
}

