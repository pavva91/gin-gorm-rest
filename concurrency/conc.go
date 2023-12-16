package concurrency

import (
	"log"
	"sync"
	"time"
)

type Point struct {
	X int
	Y int
}

var(
	mu sync.Mutex
	Lock sync.Mutex
)
func ReadModThreadSafe(p *Point, mode int) {
	mu.Lock()
	defer mu.Unlock()

	log.Printf("Value Before: %d, mode %d", p.X, mode)
	if mode == 1 {
		time.Sleep(1 * time.Second)
		p.X = 1
	}
	if mode == 0 {
		p.X = 0
	}
	log.Printf("Value After: %d, mode %d", p.X, mode)
}

func ReadModNotThreadSafe1(p *Point, mode int) {
	var muLocal sync.Mutex
	muLocal.Lock()
	defer muLocal.Unlock()

	log.Printf("Value Before: %d, mode %d", p.X, mode)
	if mode == 1 {
		time.Sleep(1 * time.Second)
		p.X = 1
	}
	if mode == 0 {
		p.X = 0
	}
	log.Printf("Value After: %d, mode %d", p.X, mode)
}

// BUG: Pass mutex by value
func ReadModNotThreadSafe2(p *Point, mode int, lock sync.Mutex) {
	lock.Lock()
	defer lock.Unlock()

	log.Printf("Value Before: %d, mode %d", p.X, mode)
	if mode == 1 {
		time.Sleep(1 * time.Second)
		p.X = 1
	}
	if mode == 0 {
		p.X = 0
	}
	log.Printf("Value After: %d, mode %d", p.X, mode)
}

func ReadModNotThreadSafe3(p *Point, mode int, lock *sync.Mutex) {
	lock.Lock()
	defer lock.Unlock()

	log.Printf("Value Before: %d, mode %d", p.X, mode)
	if mode == 1 {
		time.Sleep(1 * time.Second)
		p.X = 1
	}
	if mode == 0 {
		p.X = 0
	}
	log.Printf("Value After: %d, mode %d", p.X, mode)
}

func ReadModNotThreadSafe4(p *Point, mode int) {
	log.Printf("Value Before: %d, mode %d", p.X, mode)
	if mode == 1 {
		time.Sleep(1 * time.Second)
		p.X = 1
	}
	if mode == 0 {
		p.X = 0
	}
	log.Printf("Value After: %d, mode %d", p.X, mode)
}
