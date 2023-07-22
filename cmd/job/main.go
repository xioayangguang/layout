package main

import "layout/cmd/job/wireinject"

func main() {
	app, cleanup, err := wireinject.NewApp()
	if err != nil {
		panic(err)
	}
	app.Run()
	defer cleanup()
}
