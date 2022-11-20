package StructuralPatterns

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFacade(t *testing.T) {
	livingRoom := new(LivingRoom)
	livingRoom.DoMovie()

	require.False(t, livingRoom.light.status)
	require.True(t, livingRoom.tv.status)
	require.True(t, livingRoom.audio.status)

	livingRoom.DoDinner()

	require.True(t, livingRoom.light.status)
	require.False(t, livingRoom.tv.status)
	require.False(t, livingRoom.audio.status)
}
