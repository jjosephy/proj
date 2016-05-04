package main

import (
    "fmt"

    "github.com/jjosephy/projview/node"
    //"github.com/jjosephy/go/ds/graph"
    "os"
)


func main() {
    fmt.Println("projview")

    n, err := node.NewDepView("data.json")
    if err != nil {
        fmt.Errorf("error %s", err)
        os.Exit(1)
    }

    /*
    l, _ := n.Graph.Adjacent(0)
    p := n.Nodes[0]
    if err != nil {
        fmt.Errorf("error %s", err)
    } else {
        fmt.Printf("main :  %v \n n %v \n", l, p)
    }
    */
    fmt.Println("------------------")
    for i := 0; i < len(n.Nodes); i++ {
        fmt.Println("Project", n.Nodes[i])
        fmt.Println("-- Dependencies --")
        l, _ := n.Graph.Adjacent(i)
        for e := l.Front(); e != nil; e = e.Next() {
            fmt.Println("Dep: ", n.Nodes[e.Value.(int)])
        }
        fmt.Println("")
	}
}
