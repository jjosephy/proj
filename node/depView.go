package node

import (
    "encoding/json"
    "fmt"
    "io/ioutil"

    "github.com/jjosephy/go/ds/graph"
)

type DepView struct {
    graph *graph.DirectedGraph
    dNodes []ProjNode
    ProjectName string
    Date string
    Status string
}

func parseMap(m map[string]interface{}) {

    //var l string
    for k := range m {
        fmt.Println("key name", k)
    }

    /*
    p := &ProjNode {
        Date  : "01/22/2016"
        Division : ""
        Name string
        Status string
        Team string
    }
    */
}


func countMap(i *int, l int, a []interface{}) {
    //p := *i
    for _, v := range a {
        d := v.(map[string]interface{})
        fmt.Println("i: ", *i , " p: ", l, " d: ", d)
        *i++
        m := d["dependencies"].([]interface{})
        if len(m) > 0 {
            countMap(i, *i - 1, m)
        }
    }
}

/*
func countM(i *int, a map[string]interface{}) {
    fmt.Println("name : ", a["name"])
    e := a["dependencies"].([]interface{})
    for _, v := range e {
        d := v.(map[string]interface{})
        fmt.Println("map it : i: ", *i , " dx: ", d)
        *i++
        m := d["dependencies"].([]interface{})
        if len(m) > 0 {
            countM(i, d)
        }
    }
}
*/

func NewDepView(fileName string) (*DepView, error){
    b, _ := ioutil.ReadFile(fileName)
    var v interface{}
    err := json.Unmarshal(b, &v)

    if err != nil {
        return nil, err
    }

    c := 1
    m := v.(map[string]interface{})
    //root node
    /*
    p := &ProjNode {
        Date : m["date"].(string),
        Division : m["division"].(string),
        Name : m["name"].(string),
        Status : m["status"].(string),
        Team : m["team"].(string),
    }
    */
    //fmt.Printf("p : %v \n", p)

    d := m["dependencies"].([]interface{})
    countMap(&c, 0, d)

    //t := c + 1
    //fmt.Println("t : ", c)

    /*
    fmt.Printf("JSON %T \n", v)
    m := v.(map[string]interface{})
    d := m["dependecies"].([]interface{})

    i := 0
    countMap(&i, d)

    fmt.Println("count : ", i + 1)
    */
    /*
    i := len(m)
    g, e := graph.NewDirectedGraph(i)

    if e != nil {
        return nil, e
    }

    d := make([]ProjNode, i)

    parseMap(m)

    return &DepView{
        graph : g,
        dNodes : d,
    }, nil
    */

    return nil, nil
}
