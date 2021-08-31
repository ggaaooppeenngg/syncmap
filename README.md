# syncmap

A generic thread safe map. Type parameter requires go 1.18

## usage
```
package main

import (
        "fmt"

        "github.com/ggaaooppeenngg/syncmap"
)

func main() {
        m := syncmap.New[int, string](map[int]string{1: "1"})
        m.Set(2, "2")
        v, ok := m.Get(2)
        if ok {
                fmt.Println(v)
        }
        m.Delete(1)

}
```
