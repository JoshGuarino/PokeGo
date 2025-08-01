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
	GetItemList(limit int, offset int) (*models.NamedResourceList, error)
	GetItemURL() string
	GetItemAttribute(nameOrId string) (*models.ItemAttribute, error)
	GetItemAttributeList(limit int, offset int) (*models.NamedResourceList, error)
	GetItemAttributeURL() string
	GetItemCategory(nameOrId string) (*models.ItemCategory, error)
	GetItemCategoryList(limit int, offset int) (*models.NamedResourceList, error)
	GetItemCategoryURL() string
	GetItemFlingEffect(nameOrId string) (*models.ItemFlingEffect, error)
	GetItemFlingEffectList(limit int, offset int) (*models.NamedResourceList, error)
	GetItemFlingEffectURL() string
	GetItemPocket(nameOrId string) (*models.ItemPocket, error)
	GetItemPocketList(limit int, offset int) (*models.NamedResourceList, error)
	GetItemPocketURL() string
}

// Items group struct
type Items struct {
	ItemURL            string
	ItemAttributeURL   string
	ItemCategoryURL    string
	ItemFlingEffectURL string
	ItemPocketURL      string
	Cache              *cache.Cache
}

// Initialize function
func init() {
	fmt.Println("Items resource group initialized")
}

// Return an instance of Items resource group struct
func NewItemsGroup() Items {
	url := endpoints.BaseURL
	return Items{
		ItemURL:            url + endpoints.Item,
		ItemAttributeURL:   url + endpoints.ItemAttribute,
		ItemCategoryURL:    url + endpoints.ItemCategory,
		ItemFlingEffectURL: url + endpoints.ItemFlingEffect,
		ItemPocketURL:      url + endpoints.ItemPocket,
		Cache:              cache.C,
	}
}

// Return a single Item resource by name or ID
func (i Items) GetItem(nameOrId string) (*models.Item, error) {
	item, err := request.GetResource[models.Item](i.ItemURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// Return a list of Item resource
func (i Items) GetItemList(limit int, offset int) (*models.NamedResourceList, error) {
	itemList, err := request.GetResourceList[models.NamedResourceList](i.ItemURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return itemList, nil
}

// Return the Item resource URL
func (i Items) GetItemURL() string {
	return i.ItemURL
}

// Return a single ItemAttribute resource by name or ID
func (i Items) GetItemAttribute(nameOrId string) (*models.ItemAttribute, error) {
	itemAttribute, err := request.GetResource[models.ItemAttribute](i.ItemAttributeURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return itemAttribute, nil
}

// Return a list of ItemAttribute resource
func (i Items) GetItemAttributeList(limit int, offset int) (*models.NamedResourceList, error) {
	itemAttributeList, err := request.GetResourceList[models.NamedResourceList](i.ItemAttributeURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return itemAttributeList, nil
}

// Return the ItemAttribute resource URL
func (i Items) GetItemAttributeURL() string {
	return i.ItemAttributeURL
}

// Return a single ItemCategory resource by name or ID
func (i Items) GetItemCategory(nameOrId string) (*models.ItemCategory, error) {
	itemCategory, err := request.GetResource[models.ItemCategory](i.ItemCategoryURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return itemCategory, nil
}

// Return a list of ItemCategory resource
func (i Items) GetItemCategoryList(limit int, offset int) (*models.NamedResourceList, error) {
	itemCategoryList, err := request.GetResourceList[models.NamedResourceList](i.ItemCategoryURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return itemCategoryList, nil
}

// Return the ItemCategory resource URL
func (i Items) GetItemCategoryURL() string {
	return i.ItemCategoryURL
}

// Return a single ItemFlingEffect resource by name or ID
func (i Items) GetItemFlingEffect(nameOrId string) (*models.ItemFlingEffect, error) {
	itemFlingEffect, err := request.GetResource[models.ItemFlingEffect](i.ItemFlingEffectURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return itemFlingEffect, nil
}

// Return a list of ItemFlingEffect resource
func (i Items) GetItemFlingEffectList(limit int, offset int) (*models.NamedResourceList, error) {
	itemFlingEffectList, err := request.GetResourceList[models.NamedResourceList](i.ItemFlingEffectURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return itemFlingEffectList, nil
}

// Return the ItemFlingEffect resource URL
func (i Items) GetItemFlingEffectURL() string {
	return i.ItemFlingEffectURL
}

// Return a single ItemPocket resource by name or ID
func (i Items) GetItemPocket(nameOrId string) (*models.ItemPocket, error) {
	itemPocket, err := request.GetResource[models.ItemPocket](i.ItemPocketURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return itemPocket, nil
}

// Return a list of ItemPocket resource
func (i Items) GetItemPocketList(limit int, offset int) (*models.NamedResourceList, error) {
	itemPocketList, err := request.GetResourceList[models.NamedResourceList](i.ItemPocketURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return itemPocketList, nil
}

// Return the ItemPocket resource URL
func (i Items) GetItemPocketURL() string {
	return i.ItemPocketURL
}
