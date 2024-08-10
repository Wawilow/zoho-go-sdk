package kommo_sdk

type ClientInterface interface {
	Configuration() *Configuration
	//SCCController() *SCCController
}

type client struct {
	configuration *Configuration
	//sCCController                  *SCCController
}

func NewClient(configuration Configuration) ClientInterface {
	newClient := &client{
		configuration: &configuration,
	}

	newBaseController := NewBaseController(&configuration)
	//newClient.sCCController = NewSCCController(*newBaseController)
	return newClient
}

func (c *client) Configuration() *Configuration {
	return c.configuration
}

//func (c *client) SCCController() *SCCController {
//	return c.sCCController
//}
