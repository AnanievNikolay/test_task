package configuration

//IServiceConfig ...
type IServiceConfig interface {
	ExternalAPI() string
	Host() string
	Port() int
	Duration() int
	WebsocketPort() int
	ConnectionString() string
}
