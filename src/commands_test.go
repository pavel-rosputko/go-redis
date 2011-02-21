package redis

import "testing"
import "runtime"
import "sync"

var client Client

func init() {
	runtime.GOMAXPROCS(10)
}

func Test(t *testing.T) {
}

const c = 4

func BenchmarkPing(bm *testing.B) {
	quit := make(chan int)
	for i := 0; i < c; i++ {
		go func() {
			client := Client{}
			s := 0
			for i := 0; i < bm.N / c; i++ {
				client.Ping()
				s++
			}

			quit <- s
		}()
	}

	s := 0
	for i := 0; i < c; i++ { s += <-quit }

	println("s =", s)
}

func BenchmarkSet(bm *testing.B) {
	wg := &sync.WaitGroup{}
	for i := 0; i < c; i++ {
		wg.Add(1)
		go func() {
			client := Client{}
			// s := 0
			for i := 0; i < bm.N / c; i++ {
				client.Set("key", "value")
				// s++
			}
			wg.Done()
		}()
	}

	wg.Wait()

	// s := 0
	// for i := 0; i < c; i++ { s += <-quit }

	// println("s =", s)
}
