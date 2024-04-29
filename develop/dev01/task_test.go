package currentTime

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCurrentTime(t *testing.T) {

	t.Run("exit code test", func(t *testing.T) {

		result := currentTime()
		require.Equal(t, 0, result)
	})
}
