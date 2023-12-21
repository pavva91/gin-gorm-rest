package concurrency_test

import (
	"log"
	"sync"
	"testing"

	"github.com/pavva91/gin-gorm-rest/concurrency"
)


// INFO: 3 ways of handling lockers works:
// 1) Use lock internal of scope of called function
// 3) Use lock in the scope of calling function
// 5) Use lock in the scope of calling function (pass by reference to called function)
// 6) Use from calling function the exported lock declared of scope of called function

func TestConcurrency1ThreadSafe(t *testing.T) {
	// var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)
	p := &concurrency.Point{
		X: 3,
		Y: 3,
	}

	go func() {
		// mu.Lock()
		// Mu.Lock()

		// defer mu.Unlock()
		// defer Mu.Unlock()
		defer wg.Done()

		concurrency.ReadModThreadSafe(p, 1)
	}()
	go func() {
		// mu.Lock()
		// Mu.Lock()
		// defer mu.Unlock()
		// defer Mu.Unlock()
		defer wg.Done()
		concurrency.ReadModThreadSafe(p, 0)
	}()
	wg.Wait()
	log.Printf("Value X: %d", p.X)
	expectedX := 0
	if p.X != expectedX {
		t.Errorf("got %d, want %d", p.X, expectedX)
	}

}

func TestConcurrency2NotThreadSafe(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	p := &concurrency.Point{
		X: 3,
		Y: 3,
	}

	go func() {
		defer wg.Done()

		concurrency.ReadModNotThreadSafe1(p, 1)
	}()
	go func() {
		defer wg.Done()
		concurrency.ReadModNotThreadSafe1(p, 0)
	}()
	wg.Wait()
	log.Printf("Value X: %d", p.X)
	expectedX := 0
	if p.X != expectedX {
		t.Errorf("got %d, want %d", p.X, expectedX)
	}

}

func TestConcurrency3ThreadSafe(t *testing.T) {
	var lock sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)
	p := &concurrency.Point{
		X: 3,
		Y: 3,
	}

	go func() {
		lock.Lock()
		defer lock.Unlock()

		defer wg.Done()

		concurrency.ReadModNotThreadSafe1(p, 1)
	}()
	go func() {
		lock.Lock()
		defer lock.Unlock()

		defer wg.Done()
		concurrency.ReadModNotThreadSafe1(p, 0)
	}()
	wg.Wait()
	log.Printf("Value X: %d", p.X)
	expectedX := 0
	if p.X != expectedX {
		t.Errorf("got %d, want %d", p.X, expectedX)
	}

}

// BUG: Pass mutex by value
func TestConcurrency4NotThreadSafe(t *testing.T) {
	var lock sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)
	p := &concurrency.Point{
		X: 3,
		Y: 3,
	}

	go func() {
		// Mu.Lock()

		// defer Mu.Unlock()
		defer wg.Done()

		concurrency.ReadModNotThreadSafe2(p, 1, lock)
	}()
	go func() {
		// Mu.Lock()
		// defer Mu.Unlock()
		defer wg.Done()
		concurrency.ReadModNotThreadSafe2(p, 0, lock)
	}()
	wg.Wait()
	log.Printf("Value X: %d", p.X)
	expectedX := 0
	if p.X != expectedX {
		t.Errorf("got %d, want %d", p.X, expectedX)
	}
}

// NOTE: Always pass mutex by reference
func TestConcurrency5ThreadSafe(t *testing.T) {
	var lock sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)
	p := &concurrency.Point{
		X: 3,
		Y: 3,
	}

	go func() {
		// Mu.Lock()

		// defer Mu.Unlock()
		defer wg.Done()

		concurrency.ReadModNotThreadSafe3(p, 1, &lock)
	}()
	go func() {
		// Mu.Lock()
		// defer Mu.Unlock()
		defer wg.Done()
		concurrency.ReadModNotThreadSafe3(p, 0, &lock)
	}()
	wg.Wait()
	log.Printf("Value X: %d", p.X)
	expectedX := 0
	if p.X != expectedX {
		t.Errorf("got %d, want %d", p.X, expectedX)
	}
}

func TestConcurrency6ThreadSafe(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	p := &concurrency.Point{
		X: 3,
		Y: 3,
	}

	go func() {
		concurrency.Lock.Lock()
		defer concurrency.Lock.Unlock()

		defer wg.Done()

		concurrency.ReadModNotThreadSafe4(p, 1)
	}()
	go func() {
		concurrency.Lock.Lock()
		defer concurrency.Lock.Unlock()

		defer wg.Done()
		concurrency.ReadModNotThreadSafe4(p, 0)
	}()
	wg.Wait()
	log.Printf("Value X: %d", p.X)
	expectedX := 0
	if p.X != expectedX {
		t.Errorf("got %d, want %d", p.X, expectedX)
	}
}
