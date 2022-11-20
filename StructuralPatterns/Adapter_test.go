package StructuralPatterns

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAdapter(t *testing.T) {
	hero := Hero{
		Name:   "neo",
		weapon: new(Knife),
	}

	require.Equal(t, hero.Skill(), knife)

	// 使用適配器 使出不相干的技能
	newHero := Hero{
		Name:   "Neal",
		weapon: NewAdapter(new(PowerOff)),
	}

	fmt.Println(newHero.Skill(), shutDown)

}
