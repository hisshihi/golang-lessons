package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Session struct {
	ID          int
	ExitSeconds int
	IsWorking   bool
	IsSwitched  bool
	// mu для защиты работы от конкурентных вызовов
	mu sync.Mutex
	// workCh канал для сигнализации о том, что сессия начала работу
	workCh      chan struct{}
	CountSwitch int
}

func NewSession(id, seconds int) *Session {
	return &Session{
		ID:          id,
		ExitSeconds: seconds,
		workCh:      make(chan struct{}),
	}
}

func (s *Session) Work() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.IsWorking {
		return
	}
	s.IsWorking = true
	close(s.workCh)
	fmt.Printf("Сессия %d начала работу\n", s.ID)
}

func (s *Session) Switch() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.IsWorking {
		return
	}

	s.IsSwitched = true
	s.CountSwitch++
	fmt.Printf("Сессия %d переключена\n", s.ID)
}

// WorkerSession обрабатывает сессии. Если сессию не взяли в работу за период её жизни, то она переключается в режим работы
func WorkerSession(ctx context.Context, session <-chan *Session) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Воркер завершён")
			return
		case s, ok := <-session:
			if !ok {
				return
			}
			go monitorSession(s)
		}
	}
}

// monitorSession отслеживает сессию и переключает её в режим работы
func monitorSession(s *Session) {
	ticker := time.NewTicker(time.Duration(s.ExitSeconds) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-s.workCh:
			return
		case <-ticker.C:
			s.Switch()
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sessionStream := make(chan *Session)
	go WorkerSession(ctx, sessionStream)

	s1 := NewSession(1, 1)
	sessionStream <- s1

	s2 := NewSession(2, 3)
	sessionStream <- s2

	time.Sleep(3 * time.Second)
	s1.Work()

	time.Sleep(10 * time.Second)
	fmt.Printf("session 1: %#v\n", s1)
	fmt.Printf("session 2: %#v\n", s2)
}
