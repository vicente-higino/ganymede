package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetDefaultsEnablesNFOGeneration(t *testing.T) {
	t.Parallel()

	var cfg Config
	cfg.SetDefaults()

	require.True(t, cfg.Archive.GenerateNFOFiles)
}

func TestSetDefaultsUsesStreamCopyForVideoConversion(t *testing.T) {
	t.Parallel()

	var cfg Config
	cfg.SetDefaults()

	require.Equal(t, "-c:v copy -c:a copy", cfg.Parameters.VideoConvert)
}
