package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sdminonne/k8s-libvirt-cloudprovider/pkg/libvirt"
	"github.com/sdminonne/k8s-libvirt-cloudprovider/restapi/operations/domain"
)

//GetAllActiveDomains know what? It returns all the active domains
var GetAllActiveDomains domain.GetAllActiveDomainsHandlerFunc = func(domain.GetAllActiveDomainsParams) middleware.Responder {
	domains, err := libvirt.ListDomains()
	if err != nil {
		return domain.NewGetAllActiveDomainsInternalServerError().WithPayload(err.Error())
	}
	return domain.NewGetAllActiveDomainsOK().WithPayload(domains)
}
