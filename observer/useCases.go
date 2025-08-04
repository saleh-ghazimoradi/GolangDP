package observer

import "C"
import (
	"fmt"
	"sync"
)

// News Feed Subscription System

type NewsAgencySubject interface {
	Register(newsFeedObserver NewsFeedObserver)
	Deregister(newsFeedObserver NewsFeedObserver)
	NotifyAll(article string)
}

type NewsAgency struct {
	NewsFeedObservers []NewsFeedObserver
	mu                sync.Mutex
}

func (n *NewsAgency) Register(newsFeedObserver NewsFeedObserver) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.NewsFeedObservers = append(n.NewsFeedObservers, newsFeedObserver)
}

func (n *NewsAgency) Deregister(newsFeedObserver NewsFeedObserver) {
	n.mu.Lock()
	defer n.mu.Unlock()
	for i, observer := range n.NewsFeedObservers {
		if observer == newsFeedObserver {
			n.NewsFeedObservers = append(n.NewsFeedObservers[:i], n.NewsFeedObservers[i+1:]...)
		}
	}
}

func (n *NewsAgency) NotifyAll(article string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	for _, observer := range n.NewsFeedObservers {
		observer.Update(article)
	}
}

func (n *NewsAgency) Publish(article string) {
	fmt.Printf("Publishing new article: %s\n", article)
	n.NotifyAll(article)
}

type NewsFeedObserver interface {
	Update(article string)
}

type MobileApp struct {
	Name string
}

func (m *MobileApp) Update(article string) {
	fmt.Printf("%s received article: %s\n", m.Name, article)
}

type Website struct {
	Url string
}

func (w *Website) Update(article string) {
	fmt.Printf("Website %s updated with article: %s\n", w.Url, article)
}
