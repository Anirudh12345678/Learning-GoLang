package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var dbData = []string{"1", "2", "3", "4", "5"}
var res = []string{}
var m = sync.Mutex{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go getData(i)
	}
	wg.Wait()
	fmt.Printf("Time taken %v", time.Since(t0))
	fmt.Printf("Results are %v", res)
}

func getData(i int) {
	del := rand.Float32() * 2000
	time.Sleep(time.Duration(del) * time.Millisecond)
	fmt.Printf("The value is %s\n", dbData[i])
	m.Lock()
	res = append(res, dbData[i])
	m.Unlock()
	wg.Done()
}
