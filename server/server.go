package server

func Init() {
	r := router()
	r.Run()
}
