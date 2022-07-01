package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type WanderingTrader struct{ Mob }

func WanderingTraderFromNbt(node *nbt.TagNodeCompound) *WanderingTrader {
	return &WanderingTrader{Mob: *MobFromNbt(node)}
}
