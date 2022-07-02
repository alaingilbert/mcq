package mcq

import (
	"fmt"
	"github.com/alaingilbert/mcq/mc"
	"github.com/alaingilbert/mcq/nbt"
)

type query struct {
	world       *World
	bboxes      []BBox
	targets     map[mc.ID]struct{}
	entities    *EntitiesConf
	dim         Dimension
	searchScope byte
}

func (q *query) hasTarget(targetID mc.ID) bool {
	if len(q.targets) == 0 {
		return true
	}
	_, found := q.targets[targetID]
	return found
}

type BBox struct {
	dim            Dimension
	x1, z1, x2, z2 int
}

func NewBBox(dim Dimension, x1, z1, x2, z2 int) BBox {
	return BBox{dim: dim, x1: x1, z1: z1, x2: x2, z2: z2}
}

func (b *BBox) Contains(x, z int) bool {
	return x >= b.x1 && x <= b.x2 && z >= b.z1 && z <= b.z2
}

type regionQ struct {
	dim  Dimension
	x, z int
}

type Result struct {
	Dim         Dimension
	X, Y, Z     int
	Description string
	NbtItem     *nbt.TagNodeCompound // TODO: Remove?
	Item        any
}

func NewResult(dim Dimension, x, y, z int, desc string, itemParsed any) Result {
	return Result{
		Dim:         dim,
		X:           x,
		Y:           y,
		Z:           z,
		Description: desc,
		Item:        itemParsed,
	}
}

func (r Result) Coord() string {
	var shortDim string
	switch r.Dim {
	case Overworld:
		shortDim = "O"
	case Nether:
		shortDim = "N"
	case TheEnd:
		shortDim = "E"
	}
	return fmt.Sprintf("[%s|%d %d %d]", shortDim, r.X, r.Y, r.Z)
}

func Q(world *World) *query {
	q := new(query)
	q.world = world
	q.bboxes = make([]BBox, 0)
	q.targets = make(map[mc.ID]struct{})
	return q
}

func (q *query) BBox(bbox BBox) *query {
	q.bboxes = append(q.bboxes, bbox)
	return q
}

func (q *query) In(dim Dimension) *query {
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
	//WithEntities bool
	//WithItems    bool
}

type EntitiesOption func(conf *EntitiesConf)

var WithCustomName = func(customName bool) EntitiesOption {
	return func(conf *EntitiesConf) {
		conf.CustomName = &customName
	}
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

func hasEntitiesScope(scope byte) bool {
	return scope&EntitiesScope == EntitiesScope
}

func hasItemsScope(scope byte) bool {
	return scope&ItemsScope == ItemsScope
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

	type regionNbbox struct {
		region *Region
		bbox   *BBox
	}

	regionsNbbox := make([]regionNbbox, 0)

	if len(q.bboxes) == 0 {
		if q.dim == 0 {
			q.world.RegionManager().Each(Overworld, func(region *Region) {
				regionsNbbox = append(regionsNbbox, regionNbbox{region, nil})
			})
			q.world.RegionManager().Each(Nether, func(region *Region) {
				regionsNbbox = append(regionsNbbox, regionNbbox{region, nil})
			})
			q.world.RegionManager().Each(TheEnd, func(region *Region) {
				regionsNbbox = append(regionsNbbox, regionNbbox{region, nil})
			})
		} else {
			q.world.RegionManager().Each(q.dim, func(region *Region) {
				regionsNbbox = append(regionsNbbox, regionNbbox{region, nil})
			})
		}
	} else {
		for _, bb := range q.bboxes {
			if bb.x2 < bb.x1 {
				bb.x1, bb.x2 = bb.x2, bb.x1
			}
			if bb.z2 < bb.z1 {
				bb.z1, bb.z2 = bb.z2, bb.z1
			}
			startX, startZ := RegionCoordinatesFromWorldXZ(bb.x1, bb.z1)
			startX *= 16 * 32
			startZ *= 16 * 32
			for tmpx := startX; tmpx <= bb.x2; tmpx += 32 * 16 {
				for tmpz := startZ; tmpz <= bb.z2; tmpz += 32 * 16 {
					rx, rz := RegionCoordinatesFromWorldXZ(tmpx, tmpz)
					region := q.world.RegionManager().GetRegion(bb.dim, rx, rz)
					regionsNbbox = append(regionsNbbox, regionNbbox{region, &bb})
				}
			}
		}
	}

	processResult := func(dim Dimension, x, y, z int, item mc.IIdentifiable, desc string) {
		if q.hasTarget(item.ID()) {
			if q.entities.CustomName != nil {
				if *q.entities.CustomName {
					if i, ok := item.(mc.IItem); ok {
						if i.Display() == "" {
							return
						}
					} else if e, ok := item.(mc.IEntity); ok {
						if e.CustomName() == "" {
							return
						}
					} else if e, ok := item.(mc.IBlockEntity); ok {
						if e.CustomName() == "" {
							return
						}
					} else {
						return
					}
				} else {
					if i, ok := item.(mc.IItem); ok {
						if i.Display() != "" {
							return
						}
					} else if e, ok := item.(mc.IEntity); ok {
						if e.CustomName() != "" {
							return
						}
					} else if e, ok := item.(mc.IBlockEntity); ok {
						if e.CustomName() != "" {
							return
						}
					} else {
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
			clb(NewResult(dim, x, y, z, description, item))
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
						if t.bbox != nil && !t.bbox.Contains(x, z) {
							return
						}

						processResult1 := func(item mc.IIdentifiable, desc string) {
							processResult(t.region.GetDimension(), x, y, z, item, desc)
						}

						// Process the block entity itself
						processResult1(blockEntity, "")

						// Lectern special case
						if lectern, ok := blockEntity.(*mc.Lectern); ok {
							if lectern.HasBook() {
								processResult1(lectern.BookItem(), " on lectern")
							}
						}

						// Process containers such as chest/barrel/shulker
						if itemsHolder, ok := blockEntity.(mc.IContainerEntity); ok {
							itemsHolder.Items().Each(func(item mc.IItem) {
								processResult1(item, " in "+blockEntity.ID().String())
								if shulkerBoxItem, ok := item.(*mc.ShulkerBoxItem); ok {
									shulkerBoxItem.ShulkerBox.Items().Each(func(item mc.IItem) {
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
						if t.bbox != nil && !t.bbox.Contains(x, z) {
							return
						}

						processResult1 := func(item mc.IIdentifiable, desc string) {
							processResult(t.region.GetDimension(), x, y, z, item, desc)
						}

						// Process entities themselves
						processResult1(entity, "")

						// Handle vehicles (boat, minecart...)
						entity.Passengers().Each(func(passenger mc.IEntity) {
							processResult1(passenger, " in "+entity.ID().String())
							// Mobs that hold something in their hand
							if mob, ok := passenger.(mc.IMob); ok {
								mob.HandItems().Each(func(handItem mc.IItem) {
									processResult1(handItem, " in "+passenger.ID().String()+" hand in "+entity.ID().String())
									if shulkerBoxItem, ok := handItem.(*mc.ShulkerBoxItem); ok {
										shulkerBoxItem.ShulkerBox.Items().Each(func(item mc.IItem) {
											processResult1(item, " in "+shulkerBoxItem.ID().String()+" in "+passenger.ID().String()+" hand in "+entity.ID().String())
										})
									}
								})
								mob.ArmorItems().Each(func(armorItem mc.IItem) {
									processResult1(armorItem, " in "+passenger.ID().String()+" armor in "+entity.ID().String())
								})
							}
						})

						// Mobs that hold something in their hand
						if mob, ok := entity.(mc.IMob); ok {
							mob.HandItems().Each(func(handItem mc.IItem) {
								processResult1(handItem, " in "+entity.ID().String()+" hand")
								if shulkerBoxItem, ok := handItem.(*mc.ShulkerBoxItem); ok {
									shulkerBoxItem.ShulkerBox.Items().Each(func(item mc.IItem) {
										processResult1(item, " in "+shulkerBoxItem.ID().String()+" in "+entity.ID().String()+" hand")
									})
								}
							})
							mob.ArmorItems().Each(func(armorItem mc.IItem) {
								processResult1(armorItem, " in "+mob.ID().String()+" armor")
							})
						}

						// Special case for item frame
						if itemFrame, ok := entity.(mc.IItemFrame); ok {
							if itemFrame.Item() != nil {
								processResult1(itemFrame.Item(), " in "+itemFrame.ID().String())
							}
							if shulkerBoxItem, ok := itemFrame.Item().(*mc.ShulkerBoxItem); ok {
								shulkerBoxItem.ShulkerBox.Items().Each(func(item mc.IItem) {
									processResult1(item, " in "+shulkerBoxItem.ID().String()+" in "+itemFrame.ID().String())
								})
							}
						}

					})
				}
			})
		}
	}
}
