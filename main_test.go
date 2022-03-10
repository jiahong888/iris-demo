package main

import (
	"fmt"
	"iris-demo/unit_test"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
)

func TestStart(t *testing.T)  {
	newApp()
	t.Run("BootStart", unit_test.BootStart)
}

func TestBufferedChan(t *testing.T)  {
	defer func() {
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
		if err := recover(); err != nil {
			fmt.Println(err)
			return
		}
	}()

	var wg sync.WaitGroup
	n := 20
	type MapData struct {
		key        string
		value      string
	}
	counter := struct {
		sync.RWMutex
		accept chan MapData
		len int64
		m map[string]string
	}{
		accept: make(chan MapData, n),
		m: make(map[string]string),
	}

	defer close(counter.accept)

	// get map
	for i:=0; i<n; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("get map i:", i)
			counter.Lock()
			defer counter.Unlock()
			defer wg.Done()
			for  {
				select {
				case accept := <-counter.accept:
					fmt.Println("recipient only sent channel a:", accept)
					if v,ok := counter.m[accept.key]; ok {
						fmt.Printf("counter %s, v is %s \n", accept.key, v)
					} else {
						counter.m[accept.key] = accept.value
						atomic.AddInt64(&counter.len, 1)
					}
					return
				}
			}
		}(i)
	}

	// write map
	for i:=0; i<n; i++ {
		fmt.Println("write map i:", i)
		s := strconv.Itoa(i)
		counter.accept <- MapData{
			key:   "jiahong"+s,
			value: "v"+s,
		}
	}

	//for i:=0; i<n; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		defer wg.Done()
	//		fmt.Println("write map i:", i)
	//		counter.accept <- i
	//		//counter.receive = counter.accept
	//		s := strconv.Itoa(i)
	//		counter.m[s] = s
	//		atomic.AddInt64(&counter.len, 1)
	//		return
	//	}(i)
	//}

	wg.Wait()

	fmt.Println(counter)
	fmt.Println(counter.len)

	return

}

func TestNonBufferedChan(t *testing.T)  {
	defer func() {
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
		if err := recover(); err != nil {
			fmt.Println(err)
			return
		}
	}()

	var wg sync.WaitGroup
	n := 20
	type MapData struct {
		key        string
		value      string
	}
	counter := struct {
		sync.RWMutex
		accept chan MapData
		len int64
		m map[string]string
	}{
		accept: make(chan MapData),
		m: make(map[string]string),
	}

	defer close(counter.accept)

	// get map
	for i:=0; i<n; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("get map i:", i)
			counter.Lock()
			defer counter.Unlock()
			defer wg.Done()
			for  {
				select {
				case accept := <-counter.accept:
					fmt.Println("recipient only sent channel a:", accept)
					if v,ok := counter.m[accept.key]; ok {
						fmt.Printf("counter %s, v is %s \n", accept.key, v)
					} else {
						counter.m[accept.key] = accept.value
						atomic.AddInt64(&counter.len, 1)
					}
					return
				}
			}
		}(i)
	}

	// write map
	for i:=0; i<n; i++ {
		fmt.Println("write map i:", i)
		s := strconv.Itoa(i)
		counter.accept <- MapData{
			key:   "jiahong"+s,
			value: "v"+s,
		}
	}

	//for i:=0; i<n; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		defer wg.Done()
	//		fmt.Println("write map i:", i)
	//		counter.accept <- i
	//		//counter.receive = counter.accept
	//		s := strconv.Itoa(i)
	//		counter.m[s] = s
	//		atomic.AddInt64(&counter.len, 1)
	//		return
	//	}(i)
	//}

	wg.Wait()

	fmt.Println(counter)
	fmt.Println(counter.len)

	return

}
