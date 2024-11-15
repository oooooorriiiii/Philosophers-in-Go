package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	id                  int
	leftFork, rightFork *sync.Mutex
}

func (p *Philosopher) dine(wg *sync.WaitGroup) {
	defer wg.Done()
	// 3回 "考える" -> "食べる" が終了したら終了
	for i := 0; i < 3; i++ {
		p.think()
		p.eat()
	}
}

func (p *Philosopher) think() {
	fmt.Printf("Philosopher %d is thinking...\n", p.id)
	time.Sleep(time.Millisecond * 100)
}

func (p *Philosopher) eat() {
	if p.id%2 == 0 {
		// Odd philosophers pick the left fork first
		p.rightFork.Lock()
		p.leftFork.Lock()
	} else {
		// Even philosophers pick the right fork first
		p.leftFork.Lock()
		p.rightFork.Lock()
	}

	fmt.Printf("Philosopher %d is eating...\n", p.id)
	time.Sleep(time.Millisecond * 100)

	p.leftFork.Unlock()
	p.rightFork.Unlock()
	fmt.Printf("Philosopher %d has finished eating.\n", p.id)
}

func main() {
	numPhilosophers := 5
	var wg sync.WaitGroup

	forks := make([]*sync.Mutex, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		forks[i] = &sync.Mutex{}
	}

	philosophers := make([]*Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &Philosopher{
			id:        i + 1,
			leftFork:  forks[i],
			rightFork: forks[(i+1)%numPhilosophers],
		}
		wg.Add(1)

		// create new goroutine
		go philosophers[i].dine(&wg)
	}

	wg.Wait()
	fmt.Println("Dinner is over.")
}
