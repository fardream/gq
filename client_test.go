package gq_test

import (
	"context"
	"testing"
	"time"

	"github.com/fardream/gq"
)

func TestNewClient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()
	client, err := gq.NewClient(ctx, "localhost:9000", "", 3)
	if err != nil {
		t.Fatal(err)
	} else {
		defer client.Close()
	}
}
