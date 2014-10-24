package categories

func init2() {
	parent := newParentCategoryForTeacher(2, "bar")
	parent.AddChild(newCategory(2001, "barChild"))
	parent.AddChild(newCategory(2002, "barChild2"))
}

