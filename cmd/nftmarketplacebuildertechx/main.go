// cmd/nftmarketplacebuildertechx/main.go
package main

import (
"flag"
"log"
"os"

"nftmarketplacebuildertechx/internal/nftmarketplacebuildertechx"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := nftmarketplacebuildertechx.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
