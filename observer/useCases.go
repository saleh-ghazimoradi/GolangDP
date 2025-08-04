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

// Stock Market Price Ticker System

type StockMarketSubject interface {
	Register(observer StockMarketObserver)
	Deregister(id string)
	Notify(stock string, price float64)
}

type StockMarket struct {
	Observers map[string]StockMarketObserver
	mu        sync.Mutex
}

func (s *StockMarket) Register(observer StockMarketObserver) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Observers[observer.TradingId()] = observer
}

func (s *StockMarket) Deregister(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.Observers, id)
}

func (s *StockMarket) Notify(stock string, price float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, observer := range s.Observers {
		observer.ReceiveUpdate(stock, price)
	}
}

func (s *StockMarket) UpdatePrice(stock string, price float64) {
	fmt.Printf("Stock %s price updated to %.2f\n", stock, price)
	s.Notify(stock, price)
}

type StockMarketObserver interface {
	TradingId() string
	ReceiveUpdate(stock string, price float64)
}

type TradingApp struct {
	Id string
}

func (t *TradingApp) TradingId() string {
	return t.Id
}

func (t *TradingApp) ReceiveUpdate(stock string, price float64) {
	fmt.Printf("TradingApp %s received update: %s price is %.2f\n", t.Id, stock, price)
}
