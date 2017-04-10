package libvirt

import (
	"fmt"

	libvirt "github.com/libvirt/libvirt-go"
)

// Handler is useless
type Handler struct {
	Connect *libvirt.Connect
}

// ListDomains list active domains
func (h *Handler) ListDomains() error {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		return fmt.Errorf("unable to connect:%v", err)
	}
	defer conn.Close()

	doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	if err != nil {
		return err
	}

	for _, dom := range doms {
		name, err := dom.GetName()
		if err == nil {
			fmt.Printf("  %s\n", name)
		}
		dom.Free() // see the doc: don't ask
	}
	return nil
}
