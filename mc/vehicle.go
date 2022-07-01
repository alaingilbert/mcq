package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

// https://minecraft.fandom.com/wiki/Entity_format#Vehicles
//
// boat, chest_boat, minecart, chest_minecart, command_block_minecart,
// furnace_minecart, hopper_minecart, spawner_minecart, tnt_minecart

type vehicle struct {
	entity
}

func VehicleFromNbt(node *nbt.TagNodeCompound) *vehicle {
	t := new(vehicle)
	t.entity = *EntityFromNbt(node)
	if passengers, ok := node.Entries["Passengers"].(*nbt.TagNodeList); ok {
		passengers.Each(func(passengerRaw nbt.ITagNode) {
			passenger := passengerRaw.(*nbt.TagNodeCompound)
			parsed := ParseEntity(passenger)
			t.passengers = append(t.passengers, parsed)
		})
	}
	return t
}
