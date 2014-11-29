package models

func ListAllStuff() []Department {
	departmentList := ListAllDepartments(false)

	for _, it := range departmentList {
		it.LoadStuff()
	}

	return departmentList
}
