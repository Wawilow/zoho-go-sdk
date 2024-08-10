package kommo_sdk

import (
	"fmt"
)

// Environment represents available environments.
type Environment string

const (
	PRODUCTION Environment = "production"
)

var (
	headers = map[string]string{"Content-Type": "application/json", "Accept": "application/json"}
)

type baseController struct {
	config *Configuration
}

func NewBaseController(config *Configuration) *baseController {
	NewBaseController := baseController{config: config}
	return &NewBaseController
}

func (bc *baseController) FormUrl(path string) string {
	return fmt.Sprintf("https://%s%s", bc.config.DomainName, path)
}
