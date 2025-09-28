package main

import (
	"context"
	"fmt"
)

func process(ctx context.Context) {
	id := ctx.Value("ID")
	fmt.Println("process id: ", id)
}

func main() {
	parentCtx := context.Background()
	ctx := context.WithValue(parentCtx, "ID", "121rdv-4df")
	process(ctx)
}
