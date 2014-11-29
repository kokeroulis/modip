package models

func ListAllStuff() []Department {
	departmentList := ListAllDepartments(false)

	for index, it := range departmentList {
        it.LoadStuff();
        departmentList[index] = it;
	}
	return departmentList
}
