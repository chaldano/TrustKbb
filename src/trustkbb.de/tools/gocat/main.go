package main

import (


"fmt"


"os"


)

func main() {



fmt.Print("Ausgabe:")
fmt.Println(os.Args)

fmt.Println("Einzelargumente:")
for _, fname := range os.Args[1:]{
  fmt.Println(fname)
  }
}
