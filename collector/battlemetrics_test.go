package collector

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBMScrape(t *testing.T) {
	result, err := fetch(gez)
	require.NoError(t, err)
	require.True(t, len(result.State.Servers.Servers) > 6)
}
