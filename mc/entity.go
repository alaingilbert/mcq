package mc

import (
	"encoding/json"
	"github.com/alaingilbert/mcq/nbt"
)

type entity struct {
	id                ID
	customName        string
	customNameVisible bool
	glowing           bool
	passengers        Entities
	pos               [3]float64
}

func EntityFromNbt(node *nbt.TagNodeCompound) *entity {
	type txt struct {
		Text string `json:"text"`
	}
	e := new(entity)
	e.id = ID(node.Entries["id"].(*nbt.TagNodeString).String())

	if customName, ok := node.Entries["CustomName"].(*nbt.TagNodeString); ok {
		var t txt
		_ = json.Unmarshal([]byte(customName.String()), &t)
		e.customName = t.Text
	}
	if customNameVisible, ok := node.Entries["CustomNameVisible"].(*nbt.TagNodeByte); ok {
		e.customNameVisible = customNameVisible.Byte() == 1
	}
	if customNameVisible, ok := node.Entries["Glowing"].(*nbt.TagNodeByte); ok {
		e.customNameVisible = customNameVisible.Byte() == 1
	}
	pos := node.Entries["Pos"].(*nbt.TagNodeList)
	x := pos.Get(0).(*nbt.TagNodeDouble).Data
	y := pos.Get(1).(*nbt.TagNodeDouble).Data
	z := pos.Get(2).(*nbt.TagNodeDouble).Data
	e.pos = [3]float64{x, y, z}
	return e
}

func (e entity) CustomName() string      { return e.customName }
func (e entity) CustomNameVisible() bool { return e.customNameVisible }
func (e entity) Glowing() bool           { return e.glowing }
func (e entity) Pos() [3]float64         { return e.pos }
func (e entity) ID() ID                  { return e.id }
func (e entity) Passengers() IEntities   { return e.passengers }

type Entities []IEntity

func (e Entities) Each(clb func(IEntity)) {
	for _, el := range e {
		clb(el)
	}
}

func ParseEntity(node *nbt.TagNodeCompound) (parsed IEntity) {
	id := ID(node.Entries["id"].(*nbt.TagNodeString).String())
	switch id {
	case BoatID:
		parsed = BoatFromNbt(node)
	case MinecartID:
		parsed = MinecartFromNbt(node)

	case ItemFrameID:
		parsed = ItemFrameFromNbt(node)
	case GlowItemFrameID:
		parsed = GlowItemFrameFromNbt(node)

	case AllayID:
		parsed = AllayFromNbt(node)
	case AxolotlID:
		parsed = AxolotlFromNbt(node)
	case BatID:
		parsed = BatFromNbt(node)
	case BeeID:
		parsed = BeeFromNbt(node)
	case BlazeID:
		parsed = BlazeFromNbt(node)
	case CatID:
		parsed = CatFromNbt(node)
	case CaveSpiderID:
		parsed = CaveSpiderFromNbt(node)
	case ChickenID:
		parsed = ChickenFromNbt(node)
	case CodID:
		parsed = CodFromNbt(node)
	case CowID:
		parsed = CowFromNbt(node)
	case CreeperID:
		parsed = CreeperFromNbt(node)
	case DolphinID:
		parsed = DolphinFromNbt(node)
	case DonkeyID:
		parsed = DonkeyFromNbt(node)
	case DrownedID:
		parsed = DrownedFromNbt(node)
	case ElderGuardianID:
		parsed = ElderGuardianFromNbt(node)
	case EnderDragonID:
		parsed = EnderDragonFromNbt(node)
	case EndermanID:
		parsed = EndermanFromNbt(node)
	case EndermiteID:
		parsed = EndermiteFromNbt(node)
	case EvokerID:
		parsed = EvokerFromNbt(node)
	case FoxID:
		parsed = FoxFromNbt(node)
	case FrogID:
		parsed = FrogFromNbt(node)
	case GhastID:
		parsed = GhastFromNbt(node)
	case GiantID:
		parsed = GiantFromNbt(node)
	case GlowSquidID:
		parsed = GlowSquidFromNbt(node)
	case GoatID:
		parsed = GoatFromNbt(node)
	case GuardianID:
		parsed = GuardianFromNbt(node)
	case HoglinID:
		parsed = HoglinFromNbt(node)
	case HorseID:
		parsed = HorseFromNbt(node)
	case HuskID:
		parsed = HuskFromNbt(node)
	case IllusionerID:
		parsed = IllusionerFromNbt(node)
	case IronGolemID:
		parsed = IronGolemFromNbt(node)
	case LlamaID:
		parsed = LlamaFromNbt(node)
	case MagmaCubeID:
		parsed = MagmaCubeFromNbt(node)
	case MooshroomID:
		parsed = MooshroomFromNbt(node)
	case MuleID:
		parsed = MuleFromNbt(node)
	case OcelotID:
		parsed = OcelotFromNbt(node)
	case PandaID:
		parsed = PandaFromNbt(node)
	case ParrotID:
		parsed = ParrotFromNbt(node)
	case PhantomID:
		parsed = PhantomFromNbt(node)
	case PigID:
		parsed = PigFromNbt(node)
	case PiglinID:
		parsed = PiglinFromNbt(node)
	case PiglinBruteID:
		parsed = PiglinBruteFromNbt(node)
	case PillagerID:
		parsed = PillagerFromNbt(node)
	case PolarBearID:
		parsed = PolarBearFromNbt(node)
	case PufferfishID:
		parsed = PufferfishFromNbt(node)
	case RabbitID:
		parsed = RabbitFromNbt(node)
	case RavagerID:
		parsed = RavagerFromNbt(node)
	case SalmonID:
		parsed = SalmonFromNbt(node)
	case SheepID:
		parsed = SheepFromNbt(node)
	case ShulkerID:
		parsed = ShulkerFromNbt(node)
	case SilverfishID:
		parsed = SilverfishFromNbt(node)
	case SkeletonID:
		parsed = SkeletonFromNbt(node)
	case SkeletonHorseID:
		parsed = SkeletonHorseFromNbt(node)
	case SlimeID:
		parsed = SlimeFromNbt(node)
	case SnowGolemID:
		parsed = SnowGolemFromNbt(node)
	case SpiderID:
		parsed = SpiderFromNbt(node)
	case StriderID:
		parsed = StriderFromNbt(node)
	case SquidID:
		parsed = SquidFromNbt(node)
	case StrayID:
		parsed = StrayFromNbt(node)
	case TadpoleID:
		parsed = TadpoleFromNbt(node)
	case TraderLlamaID:
		parsed = TraderLlamaFromNbt(node)
	case TropicalFishID:
		parsed = TropicalFishFromNbt(node)
	case TurtleID:
		parsed = TurtleFromNbt(node)
	case VexID:
		parsed = VexFromNbt(node)
	case VillagerID:
		parsed = VillagerFromNbt(node)
	case VindicatorID:
		parsed = VindicatorFromNbt(node)
	case WanderingTraderID:
		parsed = WanderingTraderFromNbt(node)
	case WardenID:
		parsed = WardenFromNbt(node)
	case WitchID:
		parsed = WitchFromNbt(node)
	case WitherID:
		parsed = WitherFromNbt(node)
	case WitherSkeletonID:
		parsed = WitherSkeletonFromNbt(node)
	case WolfID:
		parsed = WolfFromNbt(node)
	case ZoglinID:
		parsed = ZoglinFromNbt(node)
	case ZombieID:
		parsed = ZombieFromNbt(node)
	case ZombieHorseID:
		parsed = ZombieHorseFromNbt(node)
	case ZombieVillagerID:
		parsed = ZombieVillagerFromNbt(node)
	case ZombifiedPiglinID:
		parsed = ZombifiedPiglinFromNbt(node)

	default:
		parsed = EntityFromNbt(node)
	}
	return
}
