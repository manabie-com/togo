package registry

import "togo/interface/controllers"

type RegistryInterface interface {
	NewAppController() controllers.AppController
}
