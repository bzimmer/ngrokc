package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bzimmer/ngrokc"
)

func main() {
	fmt.Printf("\n** Make sure the ngrok daemon is running on localhost. **\n\n")
	ng, err := ngrokc.NewClient()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	tns, err := ng.Tunnels.Tunnels(ctx)
	if err != nil {
		panic(err)
	}
	for _, tn := range tns {
		fmt.Println(tn.Name, tn.PublicURL)
	}
	os.Exit(0)
}
