package main

import "btl/api/router"

func main() {
	r, err := router.NewRouter()
	if err != nil {
		panic(err)
	}
	r.Run()
}
