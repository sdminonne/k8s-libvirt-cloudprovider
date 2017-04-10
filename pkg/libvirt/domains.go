package libvirt

import (
	"fmt"

	libvirt "github.com/libvirt/libvirt-go"
	"github.com/sdminonne/k8s-libvirt-cloudprovider/models"
)

// ListDomains list active domains
func ListDomains() ([]*models.Domain, error) {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		return nil, fmt.Errorf("unable to connect:%v", err)
	}
	defer conn.Close()

	domains := []*models.Domain{}

	doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	if err != nil {
		return nil, err
	}
	for _, dom := range doms {
		m := &models.Domain{}
		id, err := dom.GetID()
		if err != nil {
			return nil, fmt.Errorf("unable to get ID for domain: %#v", dom)
		}
		m.ID = int32(id)

		xml, err := dom.GetXMLDesc(0)
		if err != nil {
			return nil, fmt.Errorf("unable to get XML for domain: %#v", dom)
		}

		m.XML = xml
		dom.Free() // see the doc: don't ask
		domains = append(domains, m)
	}
	return domains, nil
}
