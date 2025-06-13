package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/tour/tree"
)

func BackgroundSay(s string) {
	for i := range 5 {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("BG-%d: %s\n", i, s)
	}
}

func Say(s string) {
	for i := range 5 {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("N-%d: %s\n", i, s)
	}
}

func GoRoutinesExample() {
	go BackgroundSay("Foo")
	Say("Bar")
}

func Sum(s []int, ch chan int) {
	sum := 0
	for i, v := range s {
		sum += v
		fmt.Printf("S[%d]: %d\n", i, v)
		time.Sleep(5 * time.Millisecond) // Simulate work
	}
	ch <- sum //? send sum to c
}

func SumExample() {
	s1 := []int{90, 80, 70, 60, 50, 40, 30}
	s2 := []int{100, 200, 300, 400, 500, 600, 700}

	ch := make(chan int)
	go Sum(s1, ch) //? start a goroutine to sum s1
	go Sum(s2, ch) //? start a goroutine to sum s2
	sum1 := <-ch   //? receive sum from c
	sum2 := <-ch   //? receive sum from c
	fmt.Printf("Sum1: %d, Sum2: %d\n", sum1, sum2)
}

func DeadlockExample() {
	ch := make(chan int) //? Create an unbuffered channel

	ch <- 1 //? This will cause a deadlock because no goroutine is reading from the channel
	fmt.Println("This will never be printed due to deadlock")
}

func DeadlockExample2() {
	ch := make(chan int, 2) //? Create a buffered channel with capacity 2

	ch <- 1
	ch <- 2
	ch <- 3 //? This will cause a deadlock because the channel is full and no goroutine is reading from it
	fmt.Println("This will never be printed due to deadlock")
}

func DeadlockFixed() {
	ch := make(chan int, 2)

	//? This goroutine runs in the background and sends values to the channel
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
	}()

	fmt.Println(<-ch) //? Read from the channel and free up space
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func SendNumbers(ch chan int) {
	for i := range 5 {
		ch <- i
	}
	close(ch) //? Close the channel to signal that no more values will be sent
}

func ReceiveWithOkExample() {
	ch := make(chan int)

	go SendNumbers(ch)

	for {
		//? Use the comma-ok idiom to check if the channel is closed
		if _, ok := <-ch; !ok { //? Channel has been closed and emptied
			fmt.Println("Channel closed, done receiving.")
			break
		}
		fmt.Println("Received!")
	}
}

func FibonacciWithChannel(n int, ch chan int) {
	a, b := 0, 1
	for range n {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}

func RangeExample() {
	ch := make(chan int, 10)
	go FibonacciWithChannel(cap(ch), ch)

	//? This does the same as v, ok := <-ch and if the channel is closed, it will exit the loop
	for v := range ch {
		fmt.Println(v)
	}
}

func SendFoo(foo chan int) {
	for i := 0; ; i += 10 {
		foo <- i
		time.Sleep(100 * time.Millisecond)
	}
}

func SendBar(bar chan int) {
	for j := 0; ; j += 100 {
		bar <- j
		time.Sleep(250 * time.Millisecond)
	}
}

func SelectExample() {
	foo := make(chan int)
	bar := make(chan int)

	go SendFoo(foo)
	go SendBar(bar)

	for range 10 {
		select {
		case msg1 := <-foo:
			fmt.Println("Foo:", msg1)
		case msg2 := <-bar:
			fmt.Println("Bar:", msg2)
		}
	}
}

func TickTackBoomExample() {
	tick, tack := time.Tick(100*time.Millisecond), time.Tick(150*time.Millisecond)
	boom := time.After(time.Second)

	start := time.Now()
	elapsed := func() time.Duration { return time.Since(start).Round(time.Millisecond) }

	for {
		select {
		case <-tick:
			fmt.Printf("[%s] tick\n", elapsed())
		case <-tack:
			fmt.Printf("[%s] tack\n", elapsed())
		case <-boom:
			fmt.Printf("[%s] BOOM!\n", elapsed())
			return
		default:
			fmt.Printf("[%s] ...\n", elapsed())
			time.Sleep(50 * time.Millisecond) //? Sleep to avoid busy waiting
		}
	}
}

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func StartWalk(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

func WalkExample() {
	ch := make(chan int)
	go StartWalk(tree.New(1), ch)

	for v := range ch {
		fmt.Println(v)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go StartWalk(t1, ch1)
	go StartWalk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if ok1 != ok2 {
			return false
		}

		if !ok1 {
			break
		}

		if v1 != v2 {
			return false
		}
	}
	return true
}

func SameExample() {
	t1 := tree.New(1)
	t2 := tree.New(1)
	t3 := tree.New(2)

	fmt.Println("t1 and t2 are same:", Same(t1, t2)) // Should be true
	fmt.Println("t1 and t3 are same:", Same(t1, t3)) // Should be false
}

type SafeCounter struct {
	mu sync.Mutex //? Mutual Exclusion
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func SafeCounterExample() {
	c := SafeCounter{v: make(map[string]int)}
	for range 1000 {
		go c.Inc("foo")
	}

	time.Sleep(time.Second)
	fmt.Println("Counter value for 'foo':", c.Value("foo")) // Should print 1000
}

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

var visited = make(map[string]bool)
var mu sync.Mutex

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	mu.Lock()
	if visited[url] {
		mu.Unlock()
		return
	}
	visited[url] = true
	mu.Unlock()

	// Fetch the page
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	// Crawl child URLs
	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, wg)
	}
}

type fakeResult struct {
	body string
	urls []string
}

type fakeFetcher map[string]*fakeResult

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

func CrawlExample() {
	var wg sync.WaitGroup
	wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher, &wg)
	wg.Wait()
}
