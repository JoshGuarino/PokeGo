package items

import (
	"fmt"

	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/endpoints"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Items group interface
type IItems interface {
	GetItem(nameOrId string) (*models.Item, error)
	GetItemList(limit int, offset int) (*models.ResourceList, error)
	GetItemAttribute(nameOrId string) (*models.ItemAttribute, error)
	GetItemAttributeList(limit int, offset int) (*models.ResourceList, error)
	GetItemCategory(nameOrId string) (*models.ItemCategory, error)
	GetItemCategoryList(limit int, offset int) (*models.ResourceList, error)
	GetItemFlingEffect(nameOrId string) (*models.ItemFlingEffect, error)
	GetItemFlingEffectList(limit int, offset int) (*models.ResourceList, error)
	GetItemPocket(nameOrId string) (*models.ItemPocket, error)
	GetItemPocketList(limit int, offset int) (*models.ResourceList, error)
}

// Items group struct
type Items struct {
	Cache *cache.Cache
}

// Initialize function
func init() {
	fmt.Println("Items resource group initialized")
}

// Return an instance of Items resource group struct
func NewItemsGroup() Items {
	return Items{
		Cache: cache.C,
	}
}

// Return a single Item resource by name or ID
func (i Items) GetItem(nameOrId string) (*models.Item, error) {
	item, err := request.GetSpecificResource[models.Item](endpoints.Item + nameOrId)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// Return a list of Item resource
func (i Items) GetItemList(limit int, offset int) (*models.ResourceList, error) {
	itemList, err := request.GetResourceList(endpoints.Item, limit, offset)
	if err != nil {
		return nil, err
	}
	return itemList, nil
}

// Return a single ItemAttribute resource by name or ID
func (i Items) GetItemAttribute(nameOrId string) (*models.ItemAttribute, error) {
	itemAttribute, err := request.GetSpecificResource[models.ItemAttribute](endpoints.ItemAttribute + nameOrId)
	if err != nil {
		return nil, err
	}
	return itemAttribute, nil
}

// Return a list of ItemAttribute resource
func (i Items) GetItemAttributeList(limit int, offset int) (*models.ResourceList, error) {
	itemAttributeList, err := request.GetResourceList(endpoints.ItemAttribute, limit, offset)
	if err != nil {
		return nil, err
	}
	return itemAttributeList, nil
}

// Return a single ItemCategory resource by name or ID
func (i Items) GetItemCategory(nameOrId string) (*models.ItemCategory, error) {
	itemCategory, err := request.GetSpecificResource[models.ItemCategory](endpoints.ItemCategory + nameOrId)
	if err != nil {
		return nil, err
	}
	return itemCategory, nil
}

// Return a list of ItemCategory resource
func (i Items) GetItemCategoryList(limit int, offset int) (*models.ResourceList, error) {
	itemCategoryList, err := request.GetResourceList(endpoints.ItemCategory, limit, offset)
	if err != nil {
		return nil, err
	}
	return itemCategoryList, nil
}

// Return a single ItemFlingEffect resource by name or ID
func (i Items) GetItemFlingEffect(nameOrId string) (*models.ItemFlingEffect, error) {
	itemFlingEffect, err := request.GetSpecificResource[models.ItemFlingEffect](endpoints.ItemFlingEffect + nameOrId)
	if err != nil {
		return nil, err
	}
	return itemFlingEffect, nil
}

// Return a list of ItemFlingEffect resource
func (i Items) GetItemFlingEffectList(limit int, offset int) (*models.ResourceList, error) {
	itemFlingEffectList, err := request.GetResourceList(endpoints.ItemFlingEffect, limit, offset)
	if err != nil {
		return nil, err
	}
	return itemFlingEffectList, nil
}

// Return a single ItemPocket resource by name or ID
func (i Items) GetItemPocket(nameOrId string) (*models.ItemPocket, error) {
	itemPocket, err := request.GetSpecificResource[models.ItemPocket](endpoints.ItemPocket + nameOrId)
	if err != nil {
		return nil, err
	}
	return itemPocket, nil
}

// Return a list of ItemPocket resource
func (i Items) GetItemPocketList(limit int, offset int) (*models.ResourceList, error) {
	itemPocketList, err := request.GetResourceList(endpoints.ItemPocket, limit, offset)
	if err != nil {
		return nil, err
	}
	return itemPocketList, nil
}
