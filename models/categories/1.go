package categories

import "github.com/kokeroulis/modip/models"

func init1() {
	g := models.CategoryGroup {
		Name: "Διδάσκων",
	}

	parent := newParentCategoryForTeacher(1, "Α.Δ. Εκπ. Προσωπικού", g)
	parent.AddChild(newCategory(1001, "Αριθμός Δημοσιεύσεων"))
	parent.AddChild(newCategory(1002, "Αναγνώριση του επιστημονικού και άλλου έργου"))
	parent.AddChild(newCategory(1003, "Ερευνητικά προγράμματα και έργα"))
	parent.AddChild(newCategory(1004, "Ερευνητικές Υποδομές"))
	parent.AddChild(newCategory(1005, "Σύνδεση με την κοινωνία"))
	parent.AddChild(newCategory(1999, "For testing"))

	parent2 := newParentCategoryForTeacher(2, " Α.Δ. Μαθημάτων Π.Π.Σ. | Doesn't work on modip", g)
	parent2.AddChild(newCategory(2001, "Doesn't work on modip"))

	parent3 := newParentCategoryForTeacher(3, " Α.Δ.Μαθημάτων Π.M.Σ. | Doesn't work on modip", g)
	parent3.AddChild(newCategory(3001, "Doesn't work on modip"))

	parent4 := newParentCategoryForTeacher(4, " Α.Δ. Ερευν. Προγραμ.", g)
	parent4.AddChild(newCategory(4001, "Αναλυτικά Στοιχεία Προγράμματος"))
}

