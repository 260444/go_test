package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports chan int, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("127.0.0.1:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// fmt.Printf("%s端口关闭了\n", address)
			results <- 0
			continue

		}
		conn.Close()
		// fmt.Printf("%s端口打开了\n", address)
		results <- p
	}
}
func main() {

	ports := make(chan int, 100)
	results := make(chan int)

	var openports []int
	var closeports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i < 1024; i++ {

			ports <- i
		}
	}()

	for i := 1; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		} else {
			closeports = append(closeports, port)
		}

	}

	close(ports)
	close(results)

	sort.Ints(openports)
	sort.Ints(closeports)

	for _, v := range openports {
		fmt.Printf("v: %d\n", v)
	}

	for _, v := range closeports {
		fmt.Printf("v: %d\n", v)
	}
	// 非并发
	/* for i := 21; i < 30; i++ {
		address := fmt.Sprintf("127.0.0.1:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {

			fmt.Printf("%s端口关闭了\n", address)
			continue
		}
		conn.Close()
		fmt.Printf("%s端口打开了\n", address)
	} */
	// 并发
	/* start := time.Now()
	var wg sync.WaitGroup
	for i := 21; i < 1000; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("127.0.0.1:%d", WaitGroup\
			conn, err := net.Dial("tcp", address)
			if err != nil {

				fmt.Printf("%s端口关闭了\n", address)
				return
			}
			conn.Close()
			fmt.Printf("%s端口打开了\n", address)
		}(i)
	}
	wg.Wait()
	d := time.Since(start) / 1e9
	fmt.Printf("%d秒\n", d) */

}
