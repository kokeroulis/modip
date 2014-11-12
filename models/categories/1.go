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
}

