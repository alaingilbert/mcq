# MCQ (Minecraft Query)

Library to query anything in a minecraft world.

## Examples

Find all written/writable books in the world
```go
world := core.NewWorld("/path/to/minecraft/saves/world")
mcq.Q(world).Targets(mc.WrittenBookID, mc.WritableBookID).Find(func(result mcq.Result) {
    if bookItem, ok := result.Item.(*mc.BookItem); ok {
        switch book := bookItem.Book.(type) {
        case *mc.WrittenBook:
            fmt.Printf("WRITTEN BOOK: %s %v %v\n", result.Coord(), result.Description, book)
        case *mc.WritableBook:
            fmt.Printf("WRITABLE BOOK: %s %v %v\n", result.Coord(), result.Description, book)
        }
    }
})
```

Find all ShulkerBox in a region
```go
mcq.Q(world).BBox(mcq.NewBBox(mcq.Overworld, 170, 160, 266, 226)).Targets(mc.ShulkerBoxID).Find(func(result mcq.Result) {
    if container, ok := result.Item.(mc.IContainerEntity); ok {
        fmt.Printf("SHULKER: %s %s %v\n", result.Coord(), result.Description, container)
    }
})
```

Find all ShulkerBox in a region where the first slot is an Apple
```go
mcq.Q(world).BBox(mcq.NewBBox(mcq.Overworld, 170, 160, 266, 226)).Targets(mc.ShulkerBoxID).Find(func(result mcq.Result) {
    if shulkerBox, ok := result.Item.(mc.IContainerEntity); ok {
        shulkerBox.Items().Each(func(item mc.IItem) {
            if item.Slot() == 0 && item.ID() == mc.AppleID {
                fmt.Printf("Found ShulkerBox with apple in first slot at %s\n", result.Coord())
            }
        })
    }
})
```

Find all named items/entities
```go
mcq.Q(world).Find(func(result mcq.Result) {
    if i, ok := result.Item.(mc.INamed); ok {
        fmt.Printf("%s %v %s\n", result.Coord(), result, i.CustomName())
    }
}, mcq.WithCustomName(true))
```

Find all Sign with text
```go
mcq.Q(world).Targets(mc.SignID).Find(func(result mcq.Result) {
    if sign, ok := result.Item.(*mc.Sign); ok {
        if sign.HasText() {
            fmt.Printf("%s %v\n", result.Coord(), sign.InlineText())
        }
    }
})
```