package main

import (
	"log"
	"sync"
	"time"
)

type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(maxG int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxG)
	for i := 0; i < maxG; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

func (p *Pool) Run(w Worker) {
	p.work <- w
}

func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}

var names = []string{"steve", "bob", "mary", "therese", "jason"}

type namePrinter struct {
	name string
}

func (n *namePrinter) Task() {
	log.Println(n.name)
	time.Sleep(time.Second)
}

func main() {
	p := New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}

			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	p.Shutdown()
}
