package mc

import (
	"fmt"
	"strings"
)

type ID string

func (i ID) String() string {
	trimmed := strings.TrimPrefix(string(i), "minecraft:")
	return fmt.Sprintf("[ID:%s]", trimmed)
}

const (
	ItemID = "minecraft:item"

	AirID   = "minecraft:air" // Blocks IDs
	GrassID = "minecraft:grass"
	PoppyID = "minecraft:poppy"
	StoneID = "minecraft:stone"

	AppleID             = "minecraft:apple"
	DiamondChestplateID = "minecraft:diamond_chestplate"

	SignID         = "minecraft:sign"
	WritableBookID = "minecraft:writable_book"
	WrittenBookID  = "minecraft:written_book"

	BoatID     = "minecraft:boat" // Transport IDs
	MinecartID = "minecraft:minecart"

	BarrelID       = "minecraft:barrel" // Block entity IDs
	ChestID        = "minecraft:chest"
	DropperID      = "minecraft:dropper"
	DispenserID    = "minecraft:dispenser"
	FurnaceID      = "minecraft:furnace"
	LecternID      = "minecraft:lectern"
	ShulkerBoxID   = "minecraft:shulker_box"
	TrappedChestID = "minecraft:trapped_chest"

	GlowItemFrameID = "minecraft:glow_item_frame"
	ItemFrameID     = "minecraft:item_frame"

	AllayID           = "minecraft:allay" // Entities IDs
	AxolotlID         = "minecraft:axolotl"
	BatID             = "minecraft:bat"
	BeeID             = "minecraft:bee"
	BlazeID           = "minecraft:blaze"
	CatID             = "minecraft:cat"
	CaveSpiderID      = "minecraft:cave_spider"
	ChickenID         = "minecraft:chicken"
	CodID             = "minecraft:cod"
	CowID             = "minecraft:cow"
	CreeperID         = "minecraft:creeper"
	DolphinID         = "minecraft:dolphin"
	DonkeyID          = "minecraft:donkey"
	DrownedID         = "minecraft:drowned"
	ElderGuardianID   = "minecraft:elder_guardian"
	EnderDragonID     = "minecraft:ender_dragon"
	EndermanID        = "minecraft:enderman"
	EndermiteID       = "minecraft:endermite"
	EvokerID          = "minecraft:evoker"
	FoxID             = "minecraft:fox"
	FrogID            = "minecraft:frog"
	GhastID           = "minecraft:ghast"
	GiantID           = "minecraft:giant"
	GlowSquidID       = "minecraft:glow_squid"
	GoatID            = "minecraft:goat"
	GuardianID        = "minecraft:guardian"
	HoglinID          = "minecraft:hoglin"
	HorseID           = "minecraft:horse"
	HuskID            = "minecraft:husk"
	IllusionerID      = "minecraft:illusioner"
	IronGolemID       = "minecraft:iron_golem"
	LlamaID           = "minecraft:llama"
	MagmaCubeID       = "minecraft:magma_cube"
	MooshroomID       = "minecraft:mooshroom"
	MuleID            = "minecraft:mule"
	OcelotID          = "minecraft:ocelot"
	PandaID           = "minecraft:panda"
	ParrotID          = "minecraft:parrot"
	PhantomID         = "minecraft:phantom"
	PigID             = "minecraft:pig"
	PiglinID          = "minecraft:piglin"
	PiglinBruteID     = "minecraft:piglin_brute"
	PillagerID        = "minecraft:pillager"
	PolarBearID       = "minecraft:polar_bear"
	PufferfishID      = "minecraft:pufferfish"
	RabbitID          = "minecraft:rabbit"
	RavagerID         = "minecraft:ravager"
	SalmonID          = "minecraft:salmon"
	SheepID           = "minecraft:sheep"
	ShulkerID         = "minecraft:shulker"
	SilverfishID      = "minecraft:silverfish"
	SkeletonID        = "minecraft:skeleton"
	SkeletonHorseID   = "minecraft:skeleton_horse"
	SlimeID           = "minecraft:slime"
	SnowGolemID       = "minecraft:snow_golem"
	SpiderID          = "minecraft:spider"
	StriderID         = "minecraft:strider"
	SquidID           = "minecraft:squid"
	StrayID           = "minecraft:stray"
	TadpoleID         = "minecraft:tadpole"
	TraderLlamaID     = "minecraft:trader_llama"
	TropicalFishID    = "minecraft:tropical_fish"
	TurtleID          = "minecraft:turtle"
	VexID             = "minecraft:vex"
	VillagerID        = "minecraft:villager"
	VindicatorID      = "minecraft:vindicator"
	WanderingTraderID = "minecraft:wandering_trader"
	WardenID          = "minecraft:warden"
	WitchID           = "minecraft:witch"
	WitherID          = "minecraft:wither"
	WitherSkeletonID  = "minecraft:wither_skeleton"
	WolfID            = "minecraft:wolf"
	ZoglinID          = "minecraft:zoglin"
	ZombieID          = "minecraft:zombie"
	ZombieHorseID     = "minecraft:zombie_horse"
	ZombieVillagerID  = "minecraft:zombie_villager"
	ZombifiedPiglinID = "minecraft:zombified_piglin"
)

func (i ID) IsTransport() bool {
	return i == BoatID ||
		i == MinecartID
}

func (i ID) IsEntity() bool {
	return i == GlowItemFrameID ||
		i == ItemFrameID ||
		i == ZombieID
}
