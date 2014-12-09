package main

import (
	"database/sql"
      _ "github.com/go-sql-driver/mysql"
//	"flag"
	"fmt"
//	"io/ioutil"
)

type dep_pps_mathima struct {
    Bundle                          string
    EntityId                        int
    RevisionId                      int    //not sure? should be identical? sure looks like that for most...
    Language                        string //needed? all should be greek(el) anyway
    Field_dep_pps_mathima_value     string //the important part
    Field_dep_pps_mathima_format    string //matters?
}

type dep_pms_mathima struct {
    Bundle                          string //
    EntityId                        int    //  these are the same as above
    RevisionId                      int    //  abstract and simplify?
    Language                        string //
    Field_dep_pms_mathima_value     string
    Field_dep_pms_mathima_format    string
}

type dep_prosopiko struct {
}

type node struct {
}

type taxonomy struct {
    Tid                     int
    Vid                     int
    Name                    string
    Description             string
//    Format                //not needed? 
//    Weight                //not needed
    Field__pms_tmima_tm     []pms_tmima
    Language                string
}

func main() {
    con, err := sql.Open("mysql","user:pass@/dbname")
    defer con.Close()

//    rows, err := con.Query("select distinct dep_node.nid from dep_node where dep_node.type = '$1'")
//    if err != nil { /* error handling */}
//    dep_node_ids := make([]*Node) // Node traversal after this point collecting nodes & ids in an array

    fmt.Println("TODO: Actually implemement relations from the tested queries by Kokeroulis")
    fmt.Println("TODO: Once we have the relations mapped, use the described tables above to get the data in 'consumable' form")
    fmt.Println("NEED: at this point this tool successfully authenticated to the db (if given the right creds")
    fmt.Println("      However to operate we need to make sense of the db first for this tool to really be useful")
    fmt.Println("      like the commented array above with simple queries (ref Mail 'DB Schema') we can get a dump of dumb data")

}

