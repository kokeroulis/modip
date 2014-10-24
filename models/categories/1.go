package categories

func init1() {
	parent := newParentCategoryForTeacher(1, "foo")
	parent.AddChild(newCategory(1001, "fooChild"))
	parent.AddChild(newCategory(1002, "fooChild2"))
}

