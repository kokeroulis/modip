package main

import (
	"fmt"
	"regexp"
	"strings"
	"io/ioutil"
	"strconv"
)

const funcParameter = "lessonId int" //""
const whereSqlCommand = "lesson = $" //""
const sqlTableReference = "lesson int references lesson(id) on delete cascade,"//""

var structName string
var contents string

func init() {
	b, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	if funcParameter == "" {
		panic("Please change the funcParameter")
	}

	if whereSqlCommand == "" {
		panic("Please change the whereSqlCommand")
	}

	contents = string(b)
}

func findStruct(line string) {
	pattern := "^type ([A-Z]|[a-z]|[0-9])+"
	matched, _ := regexp.MatchString(pattern, line)

	if !matched {
		return
	}

	re := regexp.MustCompile(pattern)
	structName = strings.Split(re.FindString(line), " ")[1]
	fmt.Println(structName)
}

type Entry struct {
	Name string
	Type string
	Key  string
}

var entries []Entry

func (e *Entry) findName(line string) {
	pattern := "([A-Z]|[a-z]|[1-9]+|_|[0-9])+"
	matched, _ := regexp.MatchString(pattern, line)

	if !matched {
		return
	}

	re := regexp.MustCompile(pattern)
	e.Name = re.FindString(line)
}

func (e *Entry) findSchema(line string) {
	pattern := "schema:\"([a-z]|[A-Z]|_|[0-9])+"
	matched, _ := regexp.MatchString(pattern, line)

	if !matched {
		return
	}

	re := regexp.MustCompile(pattern)
	e.Key = strings.Split(re.FindString(line), "schema:\"")[1]
}

func (e *Entry) findType(line string) {
	result := strings.Replace(line, e.Name, "", -1)
	result = strings.Replace(result, e.Key, "", -1)
	result = strings.Replace(result, "`schema:\"", "", -1)
	result = strings.Replace(result, "\"`", "", -1)

	result = strings.Trim(result, " ")
	result = strings.Trim(result, "\t")

	e.Type = result
}

func createPostgresTable() string {
	result := "CREATE TABLE " + structName + " ("
	result += "\n"

	result += "\t"
	result += sqlTableReference
	result += "\n"

	for index, e := range entries {
		result += "\t"
		result += e.Key
		result += " "

		result += findCorrectType(e)

		if len(entries) -1 != index {
			result += ","
		}
		result += "\n"
	}

	result += ")"

	return result
}

func findCorrectType(e Entry) string {
	e.Type = strings.Trim(e.Type, " ")
	e.Type = strings.Trim(e.Type, "\t")

	switch e.Type {
	case "string":
		return "text"
	case "int":
		return "int"
	case "bool":
		return "boolean"
	default:
		err := "Unknown type: " + e.Type
		panic(err)
	}
}

func createUpdateFunc() string {
	result := "func (f *" + structName + ") "
	result += "Update(" + funcParameter + ") {"
	result += "\n"
	result += "query := `UPDATE " + structName + " SET \n"
	for index, e := range entries {
		result += "\t"
		result += e.Key
		result += " = "
		result += "$" + strconv.Itoa(index + 1)

		if len(entries) -1 != index {
			result += ","
		}

		result += "\n"
	}

	result += "WHERE " + whereSqlCommand + "`"

	result += "\n\n"
	result += "_, err := Db.Database.Exec(query, \n"

	for index, e := range entries {
		result += "\t"
		result += e.Name

		if len(entries) -1 != index {
			result += ", "
		}

		result += "\n"
	}

	result += ")"
	result += "\n\n"
	result += "Db.CheckQueryWithNoRows(err, query)"
	result += "\n"
	result += "}"

	return result
}

func createSelectFunc() string {
	result := "func (f *" + structName + ") "
	result += "Load(" + funcParameter + ") {"
	result += "\n"

	result += "query := `SELECT \n"
	for index, e := range entries {
		result += "\t"
		result += e.Key

		if len(entries) -1 != index {
			result += ","
		}

		result += "\n"
	}

	result += "FROM " + structName + "\n"
	result += "WHERE " + whereSqlCommand + "`"
	result += "\n\n"

	result += "err := Db.Database.QueryRow(query, TODO). \n"
	result += "\tScan("
	for index, e := range entries {
		if index != 0 {
			result += "\t\t"
		}
		result += "&f."
		result += e.Name

		if len(entries) -1 != index {
			result += ","
		}

		result += "\n"
	}

	result += ")"
	result += "\n\n"
	result += "Db.CheckQueryWithNoRows(err, query)"

	result += "\n"
	result += "}"

	return result
}

func main() {
	l := strings.Split(contents, "\n")

	for index, it := range l {
		if index == 0 {
			findStruct(it)
		}

		if strings.Contains(it, "schema:\"") {
			e := Entry{}
			e.findName(it)
			e.findSchema(it)
			e.findType(it)

			entries = append(entries, e)
		}
	}

	var out string
	out = "=============== Start create table\n"
	out += createPostgresTable() + "\n"
	out += "=============== End Create table\n"

	out += "=============== Start create update func\n"
	out += createUpdateFunc() + "\n"
	out += "=============== End Create update func\n"

	out += "=============== Start create select func\n"
	out += createSelectFunc() + "\n"
	out += "=============== End Create select func\n"

	ioutil.WriteFile("out.txt", []byte(out), 0644)
	fmt.Println(out)
}
