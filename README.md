# libanvil


Install `anvil` (from here https://github.com/foundry-rs/foundry) and then use `libanvil` like this:

```go
package main

import (
	"fmt"
	"github.com/xrash/libanvil"
)

func main() {
	a, err := libanvil.RunAnvil(nil)
	if err != nil {
		panic(err)
	}

	defer a.Stop()
}
```

