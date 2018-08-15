package info

import (
	"github.com/pip-services-go/pip-services-commons-go/refer"
	"github.com/pip-services-go/pip-services-components-go/build"
)

var ContextInfoDescriptor = refer.NewDescriptor("pip-services", "context-info", "default", "*", "1.0")
var ContainerInfoDescriptor = refer.NewDescriptor("pip-services", "container-info", "default", "*", "1.0")
var ContainerInfoDescriptor2 = refer.NewDescriptor("pip-services-container", "container-info", "default", "*", "1.0")

func NewDefaultInfoFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(ContextInfoDescriptor, NewContextInfo)
	factory.RegisterType(ContainerInfoDescriptor, NewContextInfo)
	factory.RegisterType(ContainerInfoDescriptor2, NewContextInfo)

	return factory
}
