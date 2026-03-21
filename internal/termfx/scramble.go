package termfx

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

var charset = []rune(
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" +
		"#$%&*+-/<=>?@[]{}" +
		"█▓▒░",
)

func mutateChars(s []rune, changes int) {
	for i := 0; i < changes; i++ {
		idx := rand.Intn(len(s))
		s[idx] = charset[rand.Intn(len(charset))]
	}
}

func Scrambler(ctx context.Context) {
	s := make([]rune, 30)
	for i := range s {
		s[i] = charset[rand.Intn(len(charset))]
	}
	for {
		select {
		case <-ctx.Done():
			return
		default:
			mutateChars(s, 5)
			fmt.Printf("\r%s", string(s))
			time.Sleep(50 * time.Millisecond)
		}
	}
}
