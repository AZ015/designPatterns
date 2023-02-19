package main

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

var (
	ErrPoolClosed   = errors.New("pool has been closed")
	ErrSizeTooSmall = errors.New("size value too small")
)

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size < 0 {
		return nil, ErrSizeTooSmall
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("Acquire:", "Shared resources")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:", "New Resources")
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		log.Println("Release:", "In Queue")
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true
	close(p.resources)

	for resource := range p.resources {
		resource.Close()
	}
}
