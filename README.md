
# Simple Ip Address Generator from subnet

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

### Example Code
```go
package main

import (
	"fmt"
	ipgen "github.com/wahyuhadi/go-ipgen"

)


func main(){

	ip := ipgen.IpAddressGen("10.30.0.0/16")

	for i:=0; i<len(ip); i++ {
		fmt.Println(ip[i])
	}

	fmt.Println(len(ip))
	// Do stuff with ip address

}
```

### Example Code with goroutine and mutex
```go
package main

import (
	"context"
	"flag"
	"fmt"
	"sync"

	ipgen "github.com/wahyuhadi/go-ipgen"
)

var (
	workerCount = flag.Int("wk", 1, "Worker count")
	subnet      = flag.String("subnet", "192.168.1.1/24", "Subnet Network")
)

type Complete struct {
	N      int
	Lock   sync.Mutex
	Cancel context.CancelFunc
}

func main() {
	flag.Parse()

	ip := ipgen.IpAddressGen(*subnet)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	ctr := &Complete{
		N:      len(ip),
		Cancel: cancel,
	}
	retChan := make(chan string, 4)
	for i := 0; i < *workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work(ctx, retChan, ctr)
		}()
	}

	for _, i := range ip {
		retChan <- i
	}

	wg.Wait()
}

func work(ctx context.Context, retChan <-chan string, ctr *Complete) {
	for {
		select {
		case <-ctx.Done():
			return
		case ip := <-retChan:
			fmt.Println(ip)
			// Do stuff with ipaddress
			ctr.Lock.Lock()
			ctr.N--
			if ctr.N <= 0 {
				ctr.Lock.Unlock()
				ctr.Cancel()
				return
			}
			ctr.Lock.Unlock()
		}
	}
}

```

```sh
$ go run apps.go --wk=4 --subnet=192.168.1.1/16
```

* --wk = worker
* --subnet = network subnet
