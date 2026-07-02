package privatetypes

import (
	"testing"

	dnsv2 "codeberg.org/miekg/dns"
	privatetypesrdata "github.com/DNSControl/dnscontrol/v4/pkg/privatetypes/rdata"
)

func TestPorkbunUrlfwd_Plain(t *testing.T) {
	y := &PORKBUNURLFWD{
		Hdr: dnsv2.Header{Name: "example.org.", Class: dnsv2.ClassINET},
		PORKBUNURLFWD: privatetypesrdata.PORKBUNURLFWD{
			Target:      "http://example.com",
			TypeName:    "urlfwd",
			IncludePath: "no",
			Wildcard:    "no",
		},
	}
	rry, err := dnsv2.New(y.String())
	if err != nil {
		t.Fatal(err)
	}
	if rry.String() != y.String() {
		t.Fatalf("PORKBUN_URLFWD string presentations should be identical:\n%s\n%s", rry.String(), y.String())
	}
}

func TestPorkbunUrlfwd_WithMetadata(t *testing.T) {
	y := &PORKBUNURLFWD{
		Hdr: dnsv2.Header{Name: "example.org.", Class: dnsv2.ClassINET},
		PORKBUNURLFWD: privatetypesrdata.PORKBUNURLFWD{
			Target:      "http://example.com",
			TypeName:    "urlfwd",
			IncludePath: "permanent",
			Wildcard:    "no",
		},
	}
	rry, err := dnsv2.New(y.String())
	if err != nil {
		t.Fatal(err)
	}
	if rry.String() != y.String() {
		t.Fatalf("PORKBUN_URLFWD string presentations should be identical:\n%s\n%s", rry.String(), y.String())
	}
}

func TestPorkbunUrlfwd_WithWildcard(t *testing.T) {
	y := &PORKBUNURLFWD{
		Hdr: dnsv2.Header{Name: "example.org.", Class: dnsv2.ClassINET},
		PORKBUNURLFWD: privatetypesrdata.PORKBUNURLFWD{
			Target:      "http://example.com",
			TypeName:    "urlfwd",
			IncludePath: "no",
			Wildcard:    "yes",
		},
	}
	rry, err := dnsv2.New(y.String())
	if err != nil {
		t.Fatal(err)
	}
	if rry.String() != y.String() {
		t.Fatalf("PORKBUN_URLFWD string presentations should be identical:\n%s\n%s", rry.String(), y.String())
	}
}
