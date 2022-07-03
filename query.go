package mcq

import (
	"fmt"
	"github.com/alaingilbert/mcq/mc"
	"github.com/alaingilbert/mcq/nbt"
	"math/bits"
)

type query struct {
	world       *World
	bboxes      []IBBox
	targets     map[mc.ID]struct{}
	entities    *EntitiesConf
	dim         mc.Dimension
	searchScope byte
}

func (q *query) hasTarget(targetID mc.ID) bool {
	if len(q.targets) == 0 {
		return true
	}
	_, found := q.targets[targetID]
	return found
}

type IBBox interface {
	Coord1() mc.ICoordinate
	Coord2() mc.ICoordinate
	Contains(coord mc.ICoordinate) bool
	Intersect(other IBBox) bool
}

type BBox struct {
	coord1, coord2 mc.ICoordinate
}

func New2DBBox(dim mc.Dimension, x1, z1, x2, z2 int) *BBox {
	return New3DBBox(dim, x1, -65536, z1, x2, 30000000, z2)
}

func New3DBBox(dim mc.Dimension, x1, y1, z1, x2, y2, z2 int) *BBox {
	if x2 < x1 {
		x1, x2 = x2, x1
	}
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if z2 < z1 {
		z1, z2 = z2, z1
	}
	return &BBox{
		coord1: mc.NewCoord(dim, x1, y1, z1),
		coord2: mc.NewCoord(dim, x2, y2, z2),
	}
}

func (b BBox) Coord1() mc.ICoordinate { return b.coord1 }
func (b BBox) Coord2() mc.ICoordinate { return b.coord2 }

func (b BBox) Contains(coord mc.ICoordinate) bool {
	return coord.X() >= b.coord1.X() && coord.X() <= b.coord2.X() &&
		coord.Y() >= b.coord1.Y() && coord.Y() <= b.coord2.Y() &&
		coord.Z() >= b.coord1.Z() && coord.Z() <= b.coord2.Z()
}

func (b BBox) Intersect(other IBBox) bool {
	return (b.Coord2().X() > other.Coord1().X()) &&
		(b.Coord1().X() < other.Coord2().X()) &&
		(b.Coord2().Y() > other.Coord1().Y()) &&
		(b.Coord1().Y() < other.Coord2().Y()) &&
		(b.Coord2().Z() > other.Coord1().Z()) &&
		(b.Coord1().Z() < other.Coord2().Z())
}

type regionQ struct {
	dim  mc.Dimension
	x, z int
}

type Result struct {
	coord       mc.ICoordinate
	Description string
	Item        mc.IIdentifiable
}

func NewResult(coord mc.ICoordinate, desc string, itemParsed mc.IIdentifiable) Result {
	return Result{
		coord:       coord,
		Description: desc,
		Item:        itemParsed,
	}
}

func (r Result) Coord() string {
	var shortDim string
	switch r.coord.Dim() {
	case mc.Overworld:
		shortDim = "O"
	case mc.Nether:
		shortDim = "N"
	case mc.TheEnd:
		shortDim = "E"
	}
	return fmt.Sprintf("[%s|%d %d %d]", shortDim, r.coord.X(), r.coord.Y(), r.coord.Z())
}

func Q(world *World) *query {
	q := new(query)
	q.world = world
	q.bboxes = make([]IBBox, 0)
	q.targets = make(map[mc.ID]struct{})
	return q
}

func (q *query) BBox(bbox IBBox) *query {
	q.bboxes = append(q.bboxes, bbox)
	return q
}

func (q *query) In(dim mc.Dimension) *query {
	q.dim = dim
	return q
}

func (q *query) Targets(targets ...string) *query {
	for _, target := range targets {
		q.targets[mc.ID(target)] = struct{}{}
	}
	return q
}

type EntitiesConf struct {
	CustomName *bool
	WithBlocks bool
	//WithEntities bool
	//WithItems    bool
}

type EntitiesOption func(conf *EntitiesConf)

var WithCustomName = func(customName bool) EntitiesOption {
	return func(conf *EntitiesConf) {
		conf.CustomName = &customName
	}
}

// WithBlocks will search all blocks in the chunks as well
var WithBlocks = func(conf *EntitiesConf) {
	conf.WithBlocks = true
}

//var WithItems = func(conf *EntitiesConf) {
//	conf.WithItems = true
//}
//
//var WithEntities = func(conf *EntitiesConf) {
//	conf.WithEntities = true
//}

const (
	EntitiesScope = 0b00000001
	ItemsScope    = 0b00000010
	BlocksScope   = 0b00000100
)

func setEntitiesScope(scope byte) byte {
	scope |= EntitiesScope
	return scope
}

func setItemsScope(scope byte) byte {
	scope |= ItemsScope
	return scope
}

func setBlockScope(scope byte) byte {
	scope |= BlocksScope
	return scope
}

func hasEntitiesScope(scope byte) bool {
	return scope&EntitiesScope == EntitiesScope
}

func hasItemsScope(scope byte) bool {
	return scope&ItemsScope == ItemsScope
}
func hasBlockScope(scope byte) bool {
	return scope&BlocksScope == BlocksScope
}

func (q *query) Block(coord mc.ICoordinate, clb func(mc.ID)) {
	dim, x, y, z := coord.Unpack()
	yy := y + 64
	rx, rz := RegionCoordinatesFromWorldXZ(x, z)
	region := q.world.RegionManager().GetRegion(dim, rx, rz)
	chunk := region.GetChunkFromWorldXZ(x, z)
	sections, ok := chunk.GetData().Root().Entries["sections"].(*nbt.TagNodeList)
	if !ok {
		return
	}
	xRemaining := x & 0b1111
	zRemaining := z & 0b1111
	sectionY := yy / SectionHeight
	section := sections.Get(sectionY).(*nbt.TagNodeCompound)
	blockStates := section.Entries["block_states"].(*nbt.TagNodeCompound)
	palette := blockStates.Entries["palette"].(*nbt.TagNodeList)
	if palette.Length() == 1 {
		blockID := mc.ID(palette.Get(0).(*nbt.TagNodeCompound).Entries["Name"].(*nbt.TagNodeString).String())
		clb(blockID)
		return
	}
	data := blockStates.Entries["data"].(*nbt.TagNodeLongArray)
	yRemaining := yy % NbSection
	blockPos := yRemaining*ZDim*XDim + zRemaining*XDim + xRemaining
	mask := uint8(0b1111)
	if palette.Length() > 64 {
		mask = 0b111_1111
	} else if palette.Length() > 32 {
		mask = 0b11_1111
	} else if palette.Length() > 16 {
		mask = 0b1_1111
	}
	ones := bits.OnesCount8(mask)
	blockLngIdx := blockPos / (64 / ones)
	lng := data.Data()[blockLngIdx]
	indexRemaining := blockPos % (64 / ones)
	blockPaletteIndex := int(uint8(lng>>(indexRemaining*ones)) & mask)
	blockID := mc.ID(palette.Get(blockPaletteIndex).(*nbt.TagNodeCompound).Entries["Name"].(*nbt.TagNodeString).String())
	clb(blockID)
}

func (q *query) Find(clb func(Result), opts ...EntitiesOption) {
	var searchScope byte
	if q.searchScope == 0 {
		for targetID := range q.targets {
			if targetID.IsEntity() {
				searchScope = setEntitiesScope(searchScope)
			} else {
				searchScope = setItemsScope(searchScope)
			}
		}
		if len(q.targets) == 0 {
			searchScope = setEntitiesScope(searchScope)
			searchScope = setItemsScope(searchScope)
		}
	} else {
		searchScope = q.searchScope
	}

	conf := new(EntitiesConf)
	for _, opt := range opts {
		opt(conf)
	}
	q.entities = conf

	if q.entities.CustomName != nil {
		searchScope = setEntitiesScope(searchScope)
	}
	if q.entities.WithBlocks {
		searchScope = setBlockScope(searchScope)
	}

	type regionNbbox struct {
		region *Region
		bbox   IBBox
	}

	regionsNbbox := make([]regionNbbox, 0)

	if len(q.bboxes) == 0 {
		if q.dim == 0 {
			q.world.RegionManager().Each(mc.Overworld, func(region *Region) {
				regionsNbbox = append(regionsNbbox, regionNbbox{region, nil})
			})
			q.world.RegionManager().Each(mc.Nether, func(region *Region) {
				regionsNbbox = append(regionsNbbox, regionNbbox{region, nil})
			})
			q.world.RegionManager().Each(mc.TheEnd, func(region *Region) {
				regionsNbbox = append(regionsNbbox, regionNbbox{region, nil})
			})
		} else {
			q.world.RegionManager().Each(q.dim, func(region *Region) {
				regionsNbbox = append(regionsNbbox, regionNbbox{region, nil})
			})
		}
	} else {
		for _, bb := range q.bboxes {
			startX, startZ := RegionCoordinatesFromWorldXZ(bb.Coord1().X(), bb.Coord1().Z())
			startX *= 16 * 32
			startZ *= 16 * 32
			for tmpx := startX; tmpx <= bb.Coord2().X(); tmpx += 32 * 16 {
				for tmpz := startZ; tmpz <= bb.Coord2().Z(); tmpz += 32 * 16 {
					rx, rz := RegionCoordinatesFromWorldXZ(tmpx, tmpz)
					region := q.world.RegionManager().GetRegion(bb.Coord1().Dim(), rx, rz)
					regionsNbbox = append(regionsNbbox, regionNbbox{region, bb})
				}
			}
		}
	}

	processResult := func(dim mc.Dimension, x, y, z int, item mc.IIdentifiable, desc string) {
		if q.hasTarget(item.ID()) {
			if q.entities.CustomName != nil {
				if *q.entities.CustomName {
					// Keep only named things
					if i, ok := item.(mc.INamed); ok {
						if i.CustomName() == "" {
							return
						}
					} else {
						return
					}
				} else {
					// Skip all named things
					if i, ok := item.(mc.INamed); ok && i.CustomName() != "" {
						return
					}
				}
			}

			description := ""
			if e, ok := item.(mc.IEntity); ok {
				description += "entity " + e.ID().String()
			} else {
				description += "found " + item.ID().String()
			}
			description += desc
			coord := mc.NewCoord(dim, x, y, z)
			clb(NewResult(coord, description, item))
		}
	}

	if hasBlockScope(searchScope) {
		for _, t := range regionsNbbox {
			t.region.Each(func(chunk *Chunk) {

				chunkBBox := New2DBBox(t.region.dim,
					chunk.GetWorldX(),
					chunk.GetWorldZ(),
					chunk.GetWorldX()+16,
					chunk.GetWorldZ()+16)

				if t.bbox.Intersect(chunkBBox) {
					chunk.Each(func(block mc.Block) {
						if t.bbox != nil && !t.bbox.Contains(block) {
							return
						}
						processResult(t.region.dim, block.X(), block.Y(), block.Z(), block, "")
					})
				}
			})
		}
	}

	for _, t := range regionsNbbox {
		if hasItemsScope(searchScope) {
			t.region.Each(func(chunk *Chunk) {
				if blockEntities, ok := chunk.GetData().Root().Entries["block_entities"].(*nbt.TagNodeList); ok {
					blockEntities.Each(func(node nbt.ITagNode) {
						blockEntityNbt := node.(*nbt.TagNodeCompound)
						blockEntity := mc.ParseBlockEntity(blockEntityNbt)
						x, y, z := blockEntity.X(), blockEntity.Y(), blockEntity.Z()
						blockCoord := mc.NewCoord(t.region.GetDimension(), x, y, z)
						if t.bbox != nil && !t.bbox.Contains(blockCoord) {
							return
						}

						processResult1 := func(item mc.IIdentifiable, desc string) {
							processResult(t.region.GetDimension(), x, y, z, item, desc)
						}

						// Process the block entity itself
						processResult1(blockEntity, "")

						// Process containers such as chest/barrel/shulker
						if itemsHolder, ok := blockEntity.(mc.IContainerEntity); ok {
							itemsHolder.Items().Each(func(item mc.IItem) {
								processResult1(item, " in "+blockEntity.ID().String())
								if shulkerBoxItem, ok := item.(*mc.ShulkerBoxItem); ok {
									shulkerBoxItem.Items().Each(func(item mc.IItem) {
										processResult1(item, " in "+shulkerBoxItem.ID().String()+" in "+blockEntity.ID().String())
									})
								}
							})
						}
					})
				}
			})
		}

		if hasItemsScope(searchScope) || hasEntitiesScope(searchScope) {
			t.region.EachEntities(func(chunk *Chunk) {
				if entities, ok := chunk.GetData().Root().Entries["Entities"].(*nbt.TagNodeList); ok {
					entities.Each(func(entityRaw nbt.ITagNode) {
						entityNbt := entityRaw.(*nbt.TagNodeCompound)
						entity := mc.ParseEntity(entityNbt)
						x, y, z := int(entity.Pos()[0]), int(entity.Pos()[1]), int(entity.Pos()[2])
						entityCoord := mc.NewCoord(t.region.GetDimension(), x, y, z)
						if t.bbox != nil && !t.bbox.Contains(entityCoord) {
							return
						}

						processResult1 := func(item mc.IIdentifiable, desc string) {
							processResult(t.region.GetDimension(), x, y, z, item, desc)
						}

						processMob := func(mob mc.IMob, desc string) {
							mob.HandItems().Each(func(handItem mc.IItem) {
								processResult1(handItem, " in "+mob.ID().String()+" hand"+desc)
								if shulkerBoxItem, ok := handItem.(*mc.ShulkerBoxItem); ok {
									shulkerBoxItem.Items().Each(func(item mc.IItem) {
										processResult1(item, " in "+shulkerBoxItem.ID().String()+" in "+mob.ID().String()+" hand"+desc)
									})
								}
							})
							mob.ArmorItems().Each(func(armorItem mc.IItem) {
								processResult1(armorItem, " in "+mob.ID().String()+" armor"+desc)
							})
						}

						processEntity := func(entity mc.IEntity, desc string) {
							// Process entities themselves
							processResult1(entity, desc)
							// Mobs that hold something in their hand/armor
							if mob, ok := entity.(mc.IMob); ok {
								processMob(mob, desc)
							}
						}

						// Process entities themselves
						processEntity(entity, "")

						// Handle vehicles (boat, minecart...)
						entity.Passengers().Each(func(passenger mc.IEntity) {
							processEntity(passenger, " in "+entity.ID().String())
						})

						// Container entities such as minecart_hopper
						if container, ok := entity.(mc.IContainerEntity); ok {
							container.Items().Each(func(item mc.IItem) {
								processResult1(item, " in "+container.ID().String())
								if shulkerBoxItem, ok := item.(*mc.ShulkerBoxItem); ok {
									shulkerBoxItem.Items().Each(func(item mc.IItem) {
										processResult1(item, " in "+shulkerBoxItem.ID().String()+" in "+container.ID().String())
									})
								}
							})
						}

					})
				}
			})
		}
	}
}
