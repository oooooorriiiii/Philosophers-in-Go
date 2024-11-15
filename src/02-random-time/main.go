package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Philosopher struct {
	id                  int
	leftFork, rightFork *sync.Mutex
}

func (p *Philosopher) dine(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		p.think()
		p.eat()
		p.sleep()
	}
}

func (p *Philosopher) sleep() {
	logAction(p.id, "sleeping")
	sleepUnitTime := time.Duration(10 * rand.Intn(10))
	time.Sleep(time.Millisecond * sleepUnitTime)
}

func (p *Philosopher) think() {
	logAction(p.id, "thinking (wait for forks)")
	waitStart := time.Now()

	if p.id%2 == 0 {
		// Odd philosophers pick the left fork first
		p.rightFork.Lock()
		p.leftFork.Lock()
	} else {
		// Even philosophers pick the right fork first
		p.leftFork.Lock()
		p.rightFork.Lock()
	}

	waitDuration := time.Since(waitStart)
	logAction(p.id, fmt.Sprintf("acquired forks after waiting %.3fs", waitDuration.Seconds()))
}

func (p *Philosopher) eat() {
	logAction(p.id, "eating")
	eatUnitTime := time.Duration(10 * rand.Intn(10))
	time.Sleep(time.Millisecond * eatUnitTime)

	p.leftFork.Unlock()
	p.rightFork.Unlock()
	logAction(p.id, "finished eating")
}

func logAction(id int, action string) {
	fmt.Printf("[%s] Philosopher %d is %s\n", time.Now().Format("15:04:05.000"), id, action)
}

func main() {
	numPhilosophers := 20
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
