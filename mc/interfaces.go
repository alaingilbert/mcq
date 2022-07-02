package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type IIdentifiable interface {
	ID() ID
}

type IItem interface {
	IIdentifiable
	Count() byte
	Slot() byte
	Display() string
	Tag() *nbt.TagNodeCompound
	Nbt() *nbt.TagNodeCompound
}

type IItems interface {
	Each(func(IItem))
}

type IEntities interface {
	Each(func(IEntity))
}

type IEntity interface {
	IIdentifiable
	CustomName() string
	CustomNameVisible() bool
	Glowing() bool
	Passengers() IEntities
	Pos() [3]float64
}

type IMob interface {
	IEntity
	HandItems() IItems
	ArmorItems() IItems
	Inventory() IItems
	CanPickUpLoot() bool
	LeftHanded() bool
	Health() float32
}

type IBlockEntity interface {
	IIdentifiable
	X() int
	Y() int
	Z() int
	CustomName() string
}

type IContainerEntity interface {
	IIdentifiable
	Items() IItems
}

type IItemFrame interface {
	IEntity
	Item() IItem
}

type INamed interface {
	CustomName() string
}
