package main

func main() {
	cl := NewClient(Config{ServerAddr: ":8000"})
	cl.Start()
}
