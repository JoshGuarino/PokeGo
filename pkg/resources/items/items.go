package items

import (
	"github.com/JoshGuarino/PokeGo/internal/constants"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Items group interface
type IItems interface {
	GetItem(nameOrId string) (*models.Item, error)
	GetItemList(options models.PaginationOptions) (*models.ResourceList, error)
	GetItemAttribute(nameOrId string) (*models.ItemAttribute, error)
	GetItemAttributeList(options models.PaginationOptions) (*models.ResourceList, error)
	GetItemCategory(nameOrId string) (*models.ItemCategory, error)
	GetItemCategoryList(options models.PaginationOptions) (*models.ResourceList, error)
	GetItemFlingEffect(nameOrId string) (*models.ItemFlingEffect, error)
	GetItemFlingEffectList(options models.PaginationOptions) (*models.ResourceList, error)
	GetItemPocket(nameOrId string) (*models.ItemPocket, error)
	GetItemPocketList(options models.PaginationOptions) (*models.ResourceList, error)
}

// Items group struct
type Items struct{}

// Return an instance of Items resource group struct
func NewItemsGroup() Items {
	return Items{}
}

// Return a single Item resource by name or ID
func (i Items) GetItem(nameOrId string) (*models.Item, error) {
	item, err := request.GetSpecificResource[models.Item](constants.ItemEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// Return a list of Item resource
func (i Items) GetItemList(options models.PaginationOptions) (*models.ResourceList, error) {
	itemList, err := request.GetResourceList(constants.ItemEndpoint, options)
	if err != nil {
		return nil, err
	}
	return itemList, nil
}

// Return a single ItemAttribute resource by name or ID
func (i Items) GetItemAttribute(nameOrId string) (*models.ItemAttribute, error) {
	itemAttribute, err := request.GetSpecificResource[models.ItemAttribute](constants.ItemAttributeEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return itemAttribute, nil
}

// Return a list of ItemAttribute resource
func (i Items) GetItemAttributeList(options models.PaginationOptions) (*models.ResourceList, error) {
	itemAttributeList, err := request.GetResourceList(constants.ItemAttributeEndpoint, options)
	if err != nil {
		return nil, err
	}
	return itemAttributeList, nil
}

// Return a single ItemCategory resource by name or ID
func (i Items) GetItemCategory(nameOrId string) (*models.ItemCategory, error) {
	itemCategory, err := request.GetSpecificResource[models.ItemCategory](constants.ItemCategoryEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return itemCategory, nil
}

// Return a list of ItemCategory resource
func (i Items) GetItemCategoryList(options models.PaginationOptions) (*models.ResourceList, error) {
	itemCategoryList, err := request.GetResourceList(constants.ItemCategoryEndpoint, options)
	if err != nil {
		return nil, err
	}
	return itemCategoryList, nil
}

// Return a single ItemFlingEffect resource by name or ID
func (i Items) GetItemFlingEffect(nameOrId string) (*models.ItemFlingEffect, error) {
	itemFlingEffect, err := request.GetSpecificResource[models.ItemFlingEffect](constants.ItemFlingEffectEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return itemFlingEffect, nil
}

// Return a list of ItemFlingEffect resource
func (i Items) GetItemFlingEffectList(options models.PaginationOptions) (*models.ResourceList, error) {
	itemFlingEffectList, err := request.GetResourceList(constants.ItemFlingEffectEndpoint, options)
	if err != nil {
		return nil, err
	}
	return itemFlingEffectList, nil
}

// Return a single ItemPocket resource by name or ID
func (i Items) GetItemPocket(nameOrId string) (*models.ItemPocket, error) {
	itemPocket, err := request.GetSpecificResource[models.ItemPocket](constants.ItemPocketEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return itemPocket, nil
}

// Return a list of ItemPocket resource
func (i Items) GetItemPocketList(options models.PaginationOptions) (*models.ResourceList, error) {
	itemPocketList, err := request.GetResourceList(constants.ItemPocketEndpoint, options)
	if err != nil {
		return nil, err
	}
	return itemPocketList, nil
}
