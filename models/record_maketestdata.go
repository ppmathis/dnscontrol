package models

import (
	"fmt"
	"testing"

	"github.com/DNSControl/dnscontrol/v4/pkg/mustbe"
)

func (dc *DomainConfig) AddTestRC(t *testing.T, label string, ttl uint32, typeNum uint16, args ...any) *RecordConfig {
	mustbe.ValidArgs(args)
	rc, err := dc.NewRecordConfig(label, ttl, typeNum, args...)
	if err != nil {
		fmt.Printf("dc.NewRecordConfig() returned %v", err)
		t.FailNow()
	}
	dc.AddRecordConfig(rc)
	return rc
}
