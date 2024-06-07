package servers

type Server interface {
	Start()
	InitAuthHttpHandler()
}
