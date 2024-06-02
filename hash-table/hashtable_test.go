package hashtable

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	ht := NewHashTable[string](5)
	fmt.Println(ht)
}
