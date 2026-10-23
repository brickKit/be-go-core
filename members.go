package main

import (
	"github.com/brickKit/be-acceptance/probes/go/v3/probe"
	"github.com/brickKit/be-sdk-go/v3/component"
	"github.com/brickKit/erp-billing/v3/billing"
	"github.com/brickKit/erp-credit/v3/credit"
	"github.com/brickKit/erp-inventory/v3/inventory"
	"github.com/brickKit/erp-sales/v3/sales"
)

// members are the definitions compiled into this shell, at the versions go.mod pins (by tag, never replace); they
// must equal shell.members in component.yaml (TestMembersMatchTheManifest).
func members() []component.Definition {
	return append(probe.Definitions(), // the experiment's probes
		inventory.Definition(), credit.Definition(), sales.Definition(), billing.Definition())
}
