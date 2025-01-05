package cmd

import (
	"github.com/nats-io/cliprompts/v2"
	"github.com/nats-io/nsc/v2/cmd/store"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
	"testing"
)

type tp struct {
	AccountUserContextParams
}

func (p *tp) SetDefaults(ctx ActionCtx) error {
	return p.AccountUserContextParams.SetDefaults(ctx)
}

func (p *tp) Validate(ctx ActionCtx) error {
	return p.AccountUserContextParams.Validate(ctx)
}

func (p *tp) Load(ctx ActionCtx) error {
	return nil
}

func (p *tp) Run(ctx ActionCtx) (store.Status, error) {
	return nil, nil
}

func (p *tp) PreInteractive(ctx ActionCtx) error {
	return p.AccountUserContextParams.Edit(ctx)
}

func (p *tp) PostInteractive(ctx ActionCtx) error {
	return nil
}

func Test_AccountUserContextParams(t *testing.T) {
	ts := NewTestStore(t, "O")
	defer ts.Done(t)

	ts.AddAccount(t, "A")
	ts.AddUser(t, "A", "a")
	ts.AddAccount(t, "B")
	ts.AddUser(t, "B", "b")

	var params tp
	cmd := &cobra.Command{
		Use:          "ucp",
		SilenceUsage: true,

		RunE: func(cmd *cobra.Command, args []string) error {
			return RunAction(cmd, args, &params)
		},
	}

	cliprompts.LogFn = t.Log
	// A, a
	inputs := []interface{}{0, 0}

	_, _, err := ExecuteInteractiveCmd(cmd, inputs)
	require.NoError(t, err)
	require.Equal(t, "A", params.AccountContextParams.Name)
	require.Equal(t, "a", params.UserContextParams.Name)
}
