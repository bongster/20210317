package api

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/stretchr/testify/require"
)

func TestBankModuleNotEmptry(t *testing.T) {
	bankModule := bank.AppModuleBasic{}
	t.Logf("%v", bankModule)
	require.NotEmpty(t, bankModule)
}
