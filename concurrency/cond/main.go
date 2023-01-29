package main

func main() {
// 	c := sync.NewCond(&sync.Mutex{})
// 
// 	var ready int
// 
// 	for i := 0; i < 10; i++ {
// 		go func(i int) {
// 			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)
// 
// 			c.L.Lock()
// 			ready++
// 			c.L.Unlock()
// 
// 			fmt.Printf("Athletes #%d is ready.\n", i)
// 			
// 			c.Broadcast()
// 		}(i)
// 	}
// 
// 	c.L.Lock()
// 	for ready != 10 {
// 		c.Wait()
// 		fmt.Println("Judge is called once.")
// 	}
// 	c.L.Unlock()
// 
// 	fmt.Println("All athletes are ready now,  kick off...")
// 
	Run()
}
