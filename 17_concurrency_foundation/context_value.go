package main

import (
	"context"
	"fmt"
	"time"
)

type ctxKey string // Best practice: use a private, typed key to avoid collisions with other packages.

const requestIDKey ctxKey = "requestID"

func reqLog(ctx context.Context, msg string) {
	id := "no ID"
	value := ctx.Value(requestIDKey)
	if value != nil {
		id = value.(string)
	}
	fmt.Printf("- [%s] %s\n", id, msg)
}

func contextValueExample() {
	ctx := context.WithValue(context.Background(), requestIDKey, "req-001")
	fakeCtx := context.WithValue(context.Background(), "otherKey", "value")

	reqLog(ctx, "start")
	time.Sleep(50 * time.Millisecond) // Simulate work
	reqLog(ctx, "processing...")
	reqLog(fakeCtx, "invalid context")
	time.Sleep(200 * time.Millisecond) // Simulate more work
	reqLog(ctx, "end")

}
