package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type CaveSpider struct{ Mob }

func CaveSpiderFromNbt(node *nbt.TagNodeCompound) *CaveSpider {
	return &CaveSpider{Mob: *MobFromNbt(node)}
}
