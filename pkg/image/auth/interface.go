package auth

type InfoInterface interface {
	RegistryUrl() string
	Id() string
	Pwd() string
}
