package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

var urls = []string{"https://example.com", "https://example.org", "https://example.net"}

func main() {
	ctx := context.Background()
	g, qctx := errgroup.WithContext(ctx)
	g.SetLimit(2)
	for _, url := range urls {
		g.Go(func() error {
			return isAvailable(qctx, url)
		})
	}

	if err := g.Wait(); err != nil {
		log.Fatalf("Some resource is not available: %v", err)
	} else {
		log.Println("All resource available")
	}
}

func isAvailable(ctx context.Context, url string) error {
	c := http.Client{}
	req, err := http.NewRequestWithContext(ctx, "OPTIONS", url, nil)
	if err != nil {
		return err
	}

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("wrong status code %d for url %v", resp.StatusCode, url)
	}

	return nil
}