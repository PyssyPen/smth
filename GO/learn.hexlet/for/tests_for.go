package main

/*


SOLUTION

А ВОТ КАК С ЭТОЙ ХУЙНЕЙ РАБОТАТЬ ХУЙ КТО ОБЪЯСНИТ БЛЯТЬ

Map iterates through the strs slice and modifies each element via mapFunc.
The func is safe and strs won't be modified.

func Map(strs []string, mapFunc func(s string) string) []string {
	mapped := make([]string, len(strs))
	for i, s := range strs {
		mapped[i] = mapFunc(s)
	}
	return mapped
}













import (
	"github.com/stretchr/testify/assert"
	"hexlet/exercise/pkg"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	a := assert.New(t)
	testMap(a, []string{"John", "Peter", "Fedor"}, []string{"john", "peter", "fedor"}, func(s string) string {
		return strings.Title(s) //nolint
	})
	testMap(a, []string{"hello", "world"}, []string{"HELLO", "WORLD"}, func(s string) string {
		return strings.ToLower(s)
	})
}

func testMap(a *assert.Assertions, expected, input []string, mapFunc func(s string) string) {
	inputCopy := make([]string, len(input))
	copy(inputCopy, input)

	a.Equal(expected, pkg.Map(input, mapFunc))
	a.Equal(inputCopy, input) // check that the initial slice hasn't been modified.
}
*/
