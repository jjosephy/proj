package node

import (
    "container/list"
    "encoding/json"
    "io/ioutil"
    "fmt"
    "github.com/jjosephy/go/ds/graph"
)

type DepView struct {
    Graph *graph.DirectedGraph
    dNodes []ProjNode
    ProjectName string
    Date string
    Status string
}

func countMap(i *int, l int, lst *list.List, a []interface{}) {
    for _, v := range a {
        d := v.(map[string]interface{})
        p := newProjNode(d, l)
        lst.PushBack(p)
        *i++
        m := d["dependencies"].([]interface{})
        if len(m) > 0 {
            countMap(i, *i - 1, lst, m)
        }
    }
}

func newProjNode(m map[string]interface{}, p int) (*ProjNode){
    return &ProjNode {
        Date : m["date"].(string),
        Division : m["division"].(string),
        Name : m["name"].(string),
        Parent: p,
        Status : m["status"].(string),
        Team : m["team"].(string),
    }
}

func NewDepView(fileName string) (*DepView, error){
    b, _ := ioutil.ReadFile(fileName)
    var v interface{}
    err := json.Unmarshal(b, &v)

    if err != nil {
        return nil, err
    }

    c := 1
    m := v.(map[string]interface{})
    l := list.New()
    //root node
    p := newProjNode(m, -1)
    l.PushBack(p)
    countMap(&c, 0, l, m["dependencies"].([]interface{}))

    i := l.Len()
    g, e := graph.NewDirectedGraph(i)

    if e != nil {
        return nil, e
    }

    dNode := make([]ProjNode, i)
    ix := 0
    for e := l.Front(); e != nil; e = e.Next() {
        var ptr *ProjNode
        ptr = e.Value.(*ProjNode)
        fmt.Printf("err %v \n", p.Parent)
        /*
        if ptr.Parent != -1 {
            err := g.AddEdge(ix, p.Parent)
            if err != nil {

            }
        }
        */
        dNode[ix] = *ptr
        ix++
    }



    view := &DepView{
        Graph : g,
        dNodes : dNode,
    }

    return view, nil
}
