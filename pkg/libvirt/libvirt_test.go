package libvirt

import (
	"testing"
)

func TestListDomains(t *testing.T) {
	handler := &Handler{}
	err := handler.ListDomains()
	if err != nil {
		t.Fatalf("cannot list domains: %v", err)
	}
	t.Errorf("Hai!")

}
