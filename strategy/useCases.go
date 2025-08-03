package strategy

import (
	"archive/zip"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

// USE CASE NUM 1

type CompressionStrategy interface {
	Compress(data []byte) ([]byte, error)
	Decompress(data []byte) ([]byte, error)
}

type GzipCompression struct{}

func (g *GzipCompression) Compress(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("input data is empty")
	}

	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	if _, err := w.Write(data); err != nil {
		return nil, fmt.Errorf("gzip write failed: %w", err)
	}
	if err := w.Close(); err != nil {
		return nil, fmt.Errorf("gzip close failed: %w", err)
	}
	return b.Bytes(), nil
}

func (g *GzipCompression) Decompress(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("input data is empty")
	}
	b := bytes.NewReader(data)
	r, err := gzip.NewReader(b)
	if err != nil {
		return nil, fmt.Errorf("gzip reader creation failed: %w", err)
	}
	defer r.Close()
	decompressed, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("gzip read failed: %w", err)
	}
	return decompressed, nil
}

func NewGzipCompression() (*GzipCompression, error) {
	return &GzipCompression{}, nil
}

type ZipCompression struct{}

func (z *ZipCompression) Compress(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("input data is empty")
	}
	var b bytes.Buffer
	w := zip.NewWriter(&b)
	f, err := w.Create("file")
	if err != nil {
		return nil, fmt.Errorf("zip file creation failed: %w", err)
	}
	if _, err := f.Write(data); err != nil {
		return nil, fmt.Errorf("zip write failed: %w", err)
	}
	if err := w.Close(); err != nil {
		return nil, fmt.Errorf("zip close failed: %w", err)
	}
	return b.Bytes(), nil
}

func (z *ZipCompression) Decompress(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("input data is empty")
	}
	r, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, fmt.Errorf("zip reader creation failed: %w", err)
	}
	if len(r.File) == 0 {
		return nil, fmt.Errorf("no files found in zip archive")
	}

	f := r.File[0]
	rc, err := f.Open()
	if err != nil {
		return nil, fmt.Errorf("zip file open failed: %w", err)
	}
	defer rc.Close()
	decompressed, err := io.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("zip read failed: %w", err)
	}
	return decompressed, nil
}

func NewZipCompression() (*ZipCompression, error) {
	return &ZipCompression{}, nil
}

type CompressorContext struct {
	CompressionStrategy CompressionStrategy
}

func (c *CompressorContext) SetCompressionStrategy(strategy CompressionStrategy) {
	c.CompressionStrategy = strategy
}

func (c *CompressorContext) CompressData(data []byte) ([]byte, error) {
	if c.CompressionStrategy == nil {
		return nil, fmt.Errorf("compression strategy is nil")
	}
	return c.CompressionStrategy.Compress(data)
}

func (c *CompressorContext) Decompress(data []byte) ([]byte, error) {
	if c.CompressionStrategy == nil {
		return nil, fmt.Errorf("compression strategy is nil")
	}
	return c.CompressionStrategy.Decompress(data)
}

func NewCompressorContext(strategy CompressionStrategy) (*CompressorContext, error) {
	return &CompressorContext{
		CompressionStrategy: strategy,
	}, nil
}

// USE CASE NUM 2

type ShippingStrategy interface {
	CalculateCost(weight, distance float64) float64
}

type StandardShipping struct{}

func (s *StandardShipping) CalculateCost(weight, distance float64) float64 {
	return weight*.5 + distance*.1
}

func NewStandardShipping() *StandardShipping {
	return &StandardShipping{}
}

type ExpressShipping struct{}

func (s *ExpressShipping) CalculateCost(weight, distance float64) float64 {
	return weight*1 + distance*.2
}

func NewExpressShipping() *ExpressShipping {
	return &ExpressShipping{}
}

type OvernightShipping struct{}

func (o *OvernightShipping) CalculateCost(weight, distance float64) float64 {
	return weight*2 + distance*.5
}

func NewOvernightShipping() *OvernightShipping {
	return &OvernightShipping{}
}

type ShippingContext struct {
	ShippingStrategy ShippingStrategy
}

func (s *ShippingContext) SetShippingStrategy(strategy ShippingStrategy) {
	s.ShippingStrategy = strategy
}

func (s *ShippingContext) CalculateShippingCost(weight, distance float64) float64 {
	return s.ShippingStrategy.CalculateCost(weight, distance)
}

func NewShippingContext(strategy ShippingStrategy) *ShippingContext {
	return &ShippingContext{
		ShippingStrategy: strategy,
	}
}

// USE CASE NUM 3

type EvictionStrategy interface {
	Evict(c *Cache)
}

type LRU struct{}

func (l *LRU) Evict(c *Cache) {
	fmt.Println("Evicting by LRU strategy")
}

func NewLRU() *LRU {
	return &LRU{}
}

type FIFO struct{}

func (f *FIFO) Evict(c *Cache) {
	fmt.Println("Evicting by FIFO strategy")
}

func NewFIFO() *FIFO {
	return &FIFO{}
}

type LFU struct{}

func (l *LFU) Evict(c *Cache) {
	fmt.Println("Evicting by LFU strategy")
}

func NewLFU() *LFU {
	return &LFU{}
}

type Cache struct {
	Storage          map[string]string
	EvictionStrategy EvictionStrategy
	Capacity         int
	MaxCapacity      int
}

func (c *Cache) SetEvictionStrategy(strategy EvictionStrategy) {
	c.EvictionStrategy = strategy
}

func (c *Cache) Add(key, value string) {
	if c.Capacity == c.MaxCapacity {
		c.Evict()
	}
	c.Capacity++
	c.Storage[key] = value
}

func (c *Cache) Evict() {
	c.EvictionStrategy.Evict(c)
	c.Capacity--
}

func (c *Cache) Get(key string) {
	delete(c.Storage, key)
}

func NewCache(e EvictionStrategy) *Cache {
	return &Cache{
		Storage:          make(map[string]string),
		EvictionStrategy: e,
		Capacity:         0,
		MaxCapacity:      2,
	}
}
