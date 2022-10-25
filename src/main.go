package main

import (
	"context"
	"flag"
	"os"
)

func main() {
	flag.Parse()
	ctx := context.Background()
	cmd, ok := commands[flag.Arg(0)]
	if !ok {
		os.Exit(1)
	}
	cmd.do(ctx, flag.Args()[1:]...)
}

var commands = map[string]struct {
	name string
	do   func(context.Context, ...string)
}{
	"start": {
		name: "start",
		do:   start,
	},
}

func start(ctx context.Context, args ...string) {
	if len(args) != 1 {
		panic("INVALID_ARGUMENT")
	}
	r := NewRouters()
	r.ConnectController()
	r.Gin.Run(":" + args[0])
}
