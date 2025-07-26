package singleton

import "sync"

var (
	instance *Singleton
	once     sync.Once // singleton instance is created only once even when in concurrent scenarios
)

type Singleton struct {
	data string
	mu   sync.RWMutex // Mutex for thread-safe access to data
}

func (s *Singleton) GetData() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.data
}

func (s *Singleton) SetData(data string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = data
}

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{
			data: "",
		}
	})
	return instance
}
