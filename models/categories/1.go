package categories

import "github.com/kokeroulis/modip/models"

func init1() {
	g := models.CategoryGroup {
		Name: "Group1",
	}

	parent := newParentCategoryForTeacher(1, "foo", g)
	parent.AddChild(newCategory(1001, "fooChild"))
	parent.AddChild(newCategory(1002, "fooChild2"))
}

