package kommo_sdk

type Configuration struct {
	Environment Environment
	DomainName  string
	Token       string
}

func DefaultConfiguration(Token string) Configuration {
	return Configuration{
		PRODUCTION,
		"yourdomain.salesdrive.me",
		Token,
	}
}

func newConfiguration(options Configuration) Configuration {
	config := Configuration{}
	return config
}
