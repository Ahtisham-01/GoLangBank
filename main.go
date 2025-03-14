package main

import "GoLangBank/api"

func main() {
    server := api.NewAPIServer(":3000")
    server.Run()
}
