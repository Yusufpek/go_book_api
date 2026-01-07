package main

import (
    "go_book_api/api/repositories"
    "go_book_api/api"
)

func main() {
    repositories.InitDB()
    r := api.NewRouter()
    r.Run(":8000")
}