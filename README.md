# goMATHi : GOlang MATHs Interface

This is an effort to learn statistics with golang and create more optimized
maths functions than python for learning data science. 
I dedicate this package to my GrandMother Smt. Gomathi Devi Jat.


## Example Usage

* Sample code to calculate mean

```go
package main

import (
    "fmt"
    "github.com/miteshbsjat/gomathi"
)

func main() {
    fmt.Printf("Mean = %f\n", gomathi.Mean(1, 2, 3, 4))
}
```

* Output

```bash
$ go run statistics.go
Mean = 2.500000
```
