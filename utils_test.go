package gossamer

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestUtilFunctions(t *testing.T) {
	assert.Equal(t, "Things(12345)", ResolveEntityLink("12345", ENTITY_THINGS))
	assert.Equal(t, "http://localhost:8000/v1.0/Things(12345)", ResolveSelfLinkUrl("12345", "Things"))
}
