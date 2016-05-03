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
        fmt.Printf("main :  %v", n.Graph)
    }
}
