package aghtest_test

import (
	"github.com/LensDNS/LensDNS/internal/aghtest"
	"github.com/LensDNS/LensDNS/internal/client"
)

// type check
//
// TODO(s.chzhen):  Resolve the import cycles and move it to aghtest.
var (
	_ client.AddressProcessor = (*aghtest.AddressProcessor)(nil)
	_ client.AddressUpdater   = (*aghtest.AddressUpdater)(nil)
)
