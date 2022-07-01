package mcq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegionCoordinatesFromChunkXZ(t *testing.T) {
	x, z := regionCoordinatesFromChunkXZ(30, -3)
	assert.Equal(t, 0, x)
	assert.Equal(t, -1, z)

	x, z = regionCoordinatesFromChunkXZ(1500, -600)
	assert.Equal(t, 46, x)
	assert.Equal(t, -19, z)
}
