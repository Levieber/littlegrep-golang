package main

import "littlegrep/app"

func main() {
	app := app.NewConfig()
	app.BuildConfig()
	app.Run()
}
