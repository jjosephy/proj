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
    Nodes []ProjNode
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
    if err := json.Unmarshal(b, &v); err != nil {
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
    dNode := make([]ProjNode, i)
    g, e := graph.NewDirectedGraph(i)
    if e != nil {
        return nil, e
    }

    ix := 0
    for e := l.Front(); e != nil; e = e.Next() {
        var ptr *ProjNode
        ptr = e.Value.(*ProjNode)
        if ptr.Parent != -1 {
            fmt.Println("x ", ptr.Parent, " y ", ix)
            if err := g.AddEdge(ptr.Parent, ix); err != nil {
                return nil, err
            }
        }
        dNode[ix] = *ptr
        ix++
    }

    view := &DepView{
        Graph : g,
        Nodes : dNode,
    }

    return view, nil
}
