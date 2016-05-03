package main

import (
    "fmt"

    "github.com/jjosephy/projview/node"
    //"github.com/jjosephy/go/ds/graph"
    //"os"
)


func main() {
    fmt.Println("projview")

    n, err := node.NewDepView("data.json")

    if err != nil {
        fmt.Errorf("error %s", err)
    } else {
        fmt.Printf("main :  %v", n)
    }

    fmt.Println()
    /*
    n, err := node.NewDepView(10)

    if err != nil {
        fmt.Errorf("error %s", err)
    }

    fmt.Printf("o : %v", n)
    fmt.Println("end")
    */
    /*
    g, e := graph.NewDirectedGraph(5)
    if e != nil {
        fmt.Errorf("error %s", e)
        os.Exit(1)
    }
    */


}
