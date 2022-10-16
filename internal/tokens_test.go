package internal

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/service/tokens"
	"github.com/databricks/databricks-sdk-go/workspaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccTokens(t *testing.T) {
	env := GetEnvOrSkipTest(t, "CLOUD_ENV")
	t.Log(env)
	ctx := context.Background()
	wsc := workspaces.New()

	token, err := wsc.Tokens.Create(ctx, tokens.CreateTokenRequest{
		Comment:         "xyz",
		LifetimeSeconds: 300,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err = wsc.Tokens.DeleteByTokenId(ctx, token.TokenInfo.TokenId)
		require.NoError(t, err)
	})

	wscInner := workspaces.New(&databricks.Config{
		Host:     wsc.Config.Host,
		Token:    token.TokenValue,
		AuthType: "pat",
	})

	me, err := wsc.CurrentUser.Me(ctx)
	require.NoError(t, err)
	me2, err := wscInner.CurrentUser.Me(ctx)
	require.NoError(t, err)
	assert.Equal(t, me2.UserName, me.UserName)
}