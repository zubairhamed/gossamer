package gossamer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResolveEntityLinkFunctions(t *testing.T) {
	assert.Equal(t, "Things(12345)", ResolveEntityLink("12345", ENTITY_THINGS))
}

func TestResolveSelfLinkUrlFunctions(t *testing.T) {
	assert.Equal(t, "http://localhost:8000/v1.0/Things(12345)", ResolveSelfLinkUrl("12345", "Things"))
}
