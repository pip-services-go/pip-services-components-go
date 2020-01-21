package info

import (
	"github.com/pip-services3-go/pip-services3-commons-go/v3/refer"
	"github.com/pip-services3-go/pip-services3-components-go/v3/build"
)

/*
Creates information components by their descriptors.
*/

var ContextInfoDescriptor = refer.NewDescriptor("pip-services", "context-info", "default", "*", "1.0")
var ContainerInfoDescriptor = refer.NewDescriptor("pip-services", "container-info", "default", "*", "1.0")
var ContainerInfoDescriptor2 = refer.NewDescriptor("pip-services-container", "container-info", "default", "*", "1.0")

// Create a new instance of the factory.
// Returns *build.Factory

func NewDefaultInfoFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(ContextInfoDescriptor, NewContextInfo)
	factory.RegisterType(ContainerInfoDescriptor, NewContextInfo)
	factory.RegisterType(ContainerInfoDescriptor2, NewContextInfo)

	return factory
}
