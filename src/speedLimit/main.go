package main

import (
	"context"
	"log"
	"os"
	"sync"

	"golang.org/x/time/rate"
)

func main() {
	defer log.Printf("Done.")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	apiConnection := Open()

	var wg sync.WaitGroup

	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ReadFile(context.Background())
			if err != nil {
				log.Printf("cannot ReadFile: %v", err)
			}
			log.Printf("ReadFile")
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ResolveAddress(context.Background())
			if err != nil {
				log.Printf("cannot ResolveAddress: %v", err)
			}
			log.Printf("ResolveAddress")
		}()
	}

	wg.Wait()
}

// APIConnection _
type APIConnection struct {
	rateLimiter *rate.Limiter
}

// Open _
func Open() *APIConnection {
	return &APIConnection{
		rateLimiter: rate.NewLimiter(rate.Limit(1), 1),
	}
}

// ReadFile _
func (api *APIConnection) ReadFile(ctx context.Context) error {
	if err := api.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	// <-time.After(500 * time.Millisecond)
	return nil
}

// ResolveAddress _
func (api *APIConnection) ResolveAddress(ctx context.Context) error {
	if err := api.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	// <-time.After(500 * time.Millisecond)
	return nil
}
