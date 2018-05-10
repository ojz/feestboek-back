package main

func main() {
	c := getConfig()
	app := build(c)
	app.run()
}
