package configuration

var settings *ServiceSettings

func Settings() *ServiceSettings {
	return settings
}

//Settings ...
type ServiceSettings struct {
	Fsym string
	Tsym string
}
