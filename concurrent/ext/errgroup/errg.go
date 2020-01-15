package main

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

// The first call to return a non-nil error cancels the group
func main() {
	var g errgroup.Group

	g.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("exec #1")
		return nil
	})

	g.Go(func() error {
		time.Sleep(10 * time.Second)
		fmt.Println("exec #2")
		return errors.New("failed to exec #2")
	})

	g.Go(func() error {
		time.Sleep(15 * time.Second)
		fmt.Println("exec #3")
		return errors.New("failed to exec #3")
	})

	if err := g.Wait(); err == nil {
		fmt.Println("Successfully exec all")
	} else {
		fmt.Println("Successfully exec:", err)
	}
}
