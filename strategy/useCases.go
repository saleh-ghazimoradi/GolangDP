package strategy

import (
	"archive/zip"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

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
