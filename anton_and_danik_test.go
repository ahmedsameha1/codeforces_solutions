package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAntonAndDanik(t *testing.T) {
	assert.Equal(t, "Anton", AntonAndDanik("ADAAAA"))
	assert.Equal(t, "Danik", AntonAndDanik("DDDAADA"))
	assert.Equal(t, "Friendship", AntonAndDanik("DADADA"))
	assert.Equal(t, "Danik", AntonAndDanik(strings.Repeat("D", 100000)))
	assert.Equal(t, "Anton", AntonAndDanik(strings.Repeat("A", 100000)))
	assert.Equal(t, "Friendship", AntonAndDanik(strings.Repeat("ADAD", 25000)))
	assert.Equal(t, "Friendship", AntonAndDanik(strings.Repeat("ADADDDDAAA", 10000)))
}
