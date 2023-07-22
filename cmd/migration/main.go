package main

func main() {

	app, cleanup, err := newApp()
	if err != nil {
		panic(err)
	}
	app.Run()
	defer cleanup()
}
