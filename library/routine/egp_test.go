package routine

import (
	"testing"
)

func TestErrGp(t *testing.T) {
	group := NewGroup(10)
	for i := 0; i < 10; i++ {
		num := i
		group.Go(func() error {
			t.Error(num)
			if num == 5{
				panic("出错")
			}
			return nil
		})
	}
	group.Wait()
}
