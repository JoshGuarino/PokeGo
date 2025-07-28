# [Items Group](https://pokeapi.co/docs/v2#items-section)

The following are the resources available in the Items group of the [PokeAPI](https://pokeapi.co/).

## [Item](https://pokeapi.co/docs/v2#items)

##### Get single Item resource by Name or ID:

```go
// Main client example with ID
item, err := client.Items.GetItem("1")

// Individual resource group example with Name
item, err := itemsGroup.GetItem("master-ball")
```

##### Get list of Item resource:

```go
// Main client example
itemList, err := client.Items.GetItemList(limit, offset)

// Individual resource group example
itemList, err := itemsGroup.GetItemList(limit, offset)
```

##### Get Item resource URL:

```go
// Main client example
itemURL := client.Items.GetItemURL()

// Individual resource group example
itemURL := itemsGroup.GetItemURL()
```

## [Item Attribute](https://pokeapi.co/docs/v2#item-attributes)

##### Get single Item Attribute resource by Name or ID:

```go
// Main client example with ID
itemAttribute, err := client.Items.GetItemAttribute("1")

// Individual resource group example with Name
itemAttribute, err := itemsGroup.GetItemAttribute("countable")
```

##### Get list of Item Attribute resource:

```go
// Main client example 
itemAttributeList, err := client.Items.GetItemAttributeList(limit, offset)

// Individual resource group example 
itemAttributeList, err := itemsGroup.GetItemAttributeList(limit, offset)
```

##### Get Item Attribute resource URL:

```go
// Main client example
itemAttributeURL := client.Items.GetItemAttributeURL()

// Individual resource group example
itemAttributeURL := itemsGroup.GetItemAttributeURL()
```

## [Item Category](https://pokeapi.co/docs/v2#item-categories)

##### Get single Item Category resource by Name or ID:

```go
// Main client example with ID
itemCategory, err := client.Items.GetItemCategory("1")

// Individual resource group example with Name
itemCategory, err := itemsGroup.GetItemCategory("stat-boosts")
```

##### Get list of Item Category resource:

```go
// Main client example 
itemCategoryList, err := client.Items.GetItemCategoryList(limit, offset)

// Individual resource group example
itemCategoryList, err := itemsGroup.GetItemCategoryList(limit, offset)
```

##### Get Item Category resource URL:

```go
// Main client example
itemCategoryURL := client.Items.GetItemCategoryURL()

// Individual resource group example
itemCategoryURL := itemsGroup.GetItemCategoryURL()
```

## [Item Fling Effect](https://pokeapi.co/docs/v2#item-fling-effects)

##### Get single Item Fling Effect resource by Name or ID:

```go
// Main client example with ID
itemFlingEffect, err := client.Items.GetItemFlingEffect("1")

// Individual resource group example with Name
itemFlingEffect, err := itemsGroup.GetItemFlingEffect("badly-poison")
```

##### Get list of Item Fling Effect resource:

```go
// Main client example 
itemFlingEffectList, err := client.Items.GetItemFlingEffectList(limit, offset)

// Individual resource group example 
itemFlingEffectList, err := itemsGroup.GetItemFlingEffectList(limit, offset)
```

##### Get Item Fling Effect resource URL:

```go
// Main client example
itemFlingEffectURL := client.Items.GetItemFlingEffectURL()

// Individual resource group example
itemFlingEffectURL := itemsGroup.GetItemFlingEffectURL()
```

## [Item Pocket](https://pokeapi.co/docs/v2#item-pockets)

##### Get single Item Pocket resource by Name or ID:

```go
// Main client example with ID
itemPocket, err := client.Items.GetItemPocket("1")

// Individual resource group example with Name
itemPocket, err := itemsGroup.GetItemPocket("misc")
```

##### Get list of Item Pocket resource:

```go
// Main client example 
itemPocketList, err := client.Items.GetItemPocketList(limit, offset)

// Individual resource group 
itemPocketList, err := itemsGroup.GetItemPocketList(limit, offset)
```

##### Get Item Pocket resource URL:

```go
// Main client example
itemPocketURL := client.Items.GetItemPocketURL()

// Individual resource group example
itemPocketURL := itemsGroup.GetItemPocketURL()
```
