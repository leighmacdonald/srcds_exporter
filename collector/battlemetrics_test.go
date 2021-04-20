package collector

import (
	"github.com/stretchr/testify/require"
	"os/exec"
	"testing"
)

func TestBMScrape(t *testing.T) {
	if _, err := exec.LookPath("fortune"); err != nil {
		t.Skipf("google-chrome not in path, skipping test")
		return
	}
	result, err := fetch()
	require.NoError(t, err)
	require.True(t, len(result.State.Servers.Servers) > 6)
}
