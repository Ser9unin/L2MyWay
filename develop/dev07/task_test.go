package orchannel

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestOr(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	t.Run("Много каналов с разной длительностью", func(t *testing.T) {
		start := time.Now()
		<-or(
			sig(2*time.Hour),
			sig(5*time.Minute),
			sig(2*time.Second),
			sig(1*time.Second),
			sig(1*time.Hour),
			sig(1*time.Minute),
		)

		require.LessOrEqual(t, int(time.Since(start).Seconds()), 3)
	})

	t.Run("Один канал 200 миллисекунд", func(t *testing.T) {
		start := time.Now()

		<-or(sig(200 * time.Millisecond))

		require.LessOrEqual(t, int(time.Since(start).Milliseconds()), 250)
	})
}
