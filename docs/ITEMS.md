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
// Main client example returning first page of 20 results
itemList, err := client.Items.GetItemList(20, 0)

// Individual resource group example returning second page of 20 results
itemList, err := itemsGroup.GetItemList(20, 20)
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
// Main client example returning first page of 20 results
itemAttributeList, err := client.Items.GetItemAttributeList(20, 0)

// Individual resource group example returning second page of 20 results
itemAttributeList, err := itemsGroup.GetItemAttributeList(20, 20)
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
// Main client example returning first page of 20 results
itemCategoryList, err := client.Items.GetItemCategoryList(20, 0)

// Individual resource group example returning second page of 20 results
itemCategoryList, err := itemsGroup.GetItemCategoryList(20, 20)
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
// Main client example returning first page of 20 results
itemFlingEffectList, err := client.Items.GetItemFlingEffectList(20, 0)

// Individual resource group example returning second page of 20 results
itemFlingEffectList, err := itemsGroup.GetItemFlingEffectList(20, 20)
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
// Main client example returning first page of 20 results
itemPocketList, err := client.Items.GetItemPocketList(20, 0)

// Individual resource group example returning second page of 20 results
itemPocketList, err := itemsGroup.GetItemPocketList(20, 20)
```
