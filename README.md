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
	fmt.Printf("Mean = %f\n", gomathi.Mean([]int{1, 2, 3, 4}))
	fmt.Printf("Median = %f\n", gomathi.Median([]int{1, 2, 2, 3, 4}))
}
```

* Output

```bash
$ go run statistics.go
Mean = 2.500000
Median = 2.000000
```

## Testing

* Please run the following command to test this package

```bash
$ go test
```
