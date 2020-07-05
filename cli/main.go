package main

import (
	"fmt"
	"os"
)

func main() {
	app := NewApp(os.Stdin)
	result := app.execute()
	fmt.Println(result)
}