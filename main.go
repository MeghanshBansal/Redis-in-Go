package main

func main() {
	s := NewServer(Config{ListenAddr: ":8000"})
	s.Start()
}
