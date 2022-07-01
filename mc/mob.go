package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

// https://minecraft.fandom.com/wiki/Entity_format#Mobs
//
// allay, axolotl, bat, bee, blaze, cat, cave_spider, chicken, cod, cow, creeper, dolphin, donkey, drowned,
// elder_guardian, ender_dragon, enderman, endermite, evoker, fox, frog, ghast, giant, glow_squid, goat, guardian,
// hoglin, horse, husk, illusioner, iron_golem, llama, magma_cube, mooshroom, mule, ocelot, panda, parrot, phantom,
// pig, piglin, piglin_brute, pillager, polar_bear, pufferfish, rabbit, ravager, salmon, sheep, shulker, silverfish,
// skeleton, skeleton_horse, slime, snow_golem, spider, strider, squid, stray, tadpole, trader_llama, tropical_fish,
// turtle, vex, villager, vindicator, wandering_trader, warden, witch, wither, wither_skeleton, wolf, zoglin, zombie,
// zombie_horse, zombie_villager, zombified_piglin

type Mob struct {
	entity
	health        float32
	leftHanded    bool
	canPickUpLoot bool
	handItems     BaseItems
	armorItems    BaseItems
	inventory     BaseItems
}

func MobFromNbt(node *nbt.TagNodeCompound) *Mob {
	e := new(Mob)
	e.entity = *EntityFromNbt(node)
	if handItems, ok := node.Entries["HandItems"].(*nbt.TagNodeList); ok {
		handItems.Each(func(handItemRaw nbt.ITagNode) {
			handItem := handItemRaw.(*nbt.TagNodeCompound)
			if len(handItem.Entries) > 0 {
				parsed := parseItemFromNbt(handItem)
				e.handItems = append(e.handItems, parsed)
			}
		})
	}
	if armorItems, ok := node.Entries["ArmorItems"].(*nbt.TagNodeList); ok {
		armorItems.Each(func(armorItemRaw nbt.ITagNode) {
			armorItem := armorItemRaw.(*nbt.TagNodeCompound)
			if len(armorItem.Entries) > 0 {
				parsed := parseItemFromNbt(armorItem)
				e.armorItems = append(e.armorItems, parsed)
			}
		})
	}
	if health, ok := node.Entries["Health"].(*nbt.TagNodeFloat); ok {
		e.health = health.Float32()
	}
	if canPickUpLoot, ok := node.Entries["CanPickUpLoot"].(*nbt.TagNodeByte); ok {
		e.canPickUpLoot = canPickUpLoot.Byte() == 1
	}
	return e
}

func (e Mob) CanPickUpLoot() bool { return e.canPickUpLoot }
func (e Mob) Health() float32     { return e.health }
func (e Mob) LeftHanded() bool    { return e.leftHanded }
func (e Mob) HandItems() IItems   { return e.handItems }
func (e Mob) ArmorItems() IItems  { return e.armorItems }
func (e Mob) Inventory() IItems   { return e.inventory }
