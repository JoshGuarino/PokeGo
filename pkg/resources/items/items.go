package items

import (
	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/env"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/charmbracelet/log"
)

// Items group resource endpoints
const (
	ItemEndpoint            = "/item/"
	ItemAttributeEndpoint   = "/item-attribute/"
	ItemCategoryEndpoint    = "/item-category/"
	ItemFlingEffectEndpoint = "/item-fling-effect/"
	ItemPocketEndpoint      = "/item-pocket/"
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
	Cache *cache.Cache
	Env   *env.Env
}

// Initialize function
func init() {
	log.Info("Items resource group initialized")
}

// Return an instance of Items resource group struct
func NewItemsGroup() Items {
	return Items{
		Cache: cache.CACHE,
		Env:   env.ENV,
	}
}

// Return a single Item resource by name or ID
func (i Items) GetItem(nameOrId string) (*models.Item, error) {
	item, err := request.GetResource[models.Item](i.GetItemURL() + nameOrId)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// Return a list of Item resource
func (i Items) GetItemList(limit int, offset int) (*models.NamedResourceList, error) {
	itemList, err := request.GetResourceList[models.NamedResourceList](i.GetItemURL(), limit, offset)
	if err != nil {
		return nil, err
	}
	return itemList, nil
}

// Return the Item resource URL
func (i Items) GetItemURL() string {
	return i.Env.URL() + ItemEndpoint
}

// Return a single ItemAttribute resource by name or ID
func (i Items) GetItemAttribute(nameOrId string) (*models.ItemAttribute, error) {
	itemAttribute, err := request.GetResource[models.ItemAttribute](i.GetItemAttributeURL() + nameOrId)
	if err != nil {
		return nil, err
	}
	return itemAttribute, nil
}

// Return a list of ItemAttribute resource
func (i Items) GetItemAttributeList(limit int, offset int) (*models.NamedResourceList, error) {
	itemAttributeList, err := request.GetResourceList[models.NamedResourceList](i.GetItemAttributeURL(), limit, offset)
	if err != nil {
		return nil, err
	}
	return itemAttributeList, nil
}

// Return the ItemAttribute resource URL
func (i Items) GetItemAttributeURL() string {
	return i.Env.URL() + ItemAttributeEndpoint
}

// Return a single ItemCategory resource by name or ID
func (i Items) GetItemCategory(nameOrId string) (*models.ItemCategory, error) {
	itemCategory, err := request.GetResource[models.ItemCategory](i.GetItemCategoryURL() + nameOrId)
	if err != nil {
		return nil, err
	}
	return itemCategory, nil
}

// Return a list of ItemCategory resource
func (i Items) GetItemCategoryList(limit int, offset int) (*models.NamedResourceList, error) {
	itemCategoryList, err := request.GetResourceList[models.NamedResourceList](i.GetItemCategoryURL(), limit, offset)
	if err != nil {
		return nil, err
	}
	return itemCategoryList, nil
}

// Return the ItemCategory resource URL
func (i Items) GetItemCategoryURL() string {
	return i.Env.URL() + ItemCategoryEndpoint
}

// Return a single ItemFlingEffect resource by name or ID
func (i Items) GetItemFlingEffect(nameOrId string) (*models.ItemFlingEffect, error) {
	itemFlingEffect, err := request.GetResource[models.ItemFlingEffect](i.GetItemFlingEffectURL() + nameOrId)
	if err != nil {
		return nil, err
	}
	return itemFlingEffect, nil
}

// Return a list of ItemFlingEffect resource
func (i Items) GetItemFlingEffectList(limit int, offset int) (*models.NamedResourceList, error) {
	itemFlingEffectList, err := request.GetResourceList[models.NamedResourceList](i.GetItemFlingEffectURL(), limit, offset)
	if err != nil {
		return nil, err
	}
	return itemFlingEffectList, nil
}

// Return the ItemFlingEffect resource URL
func (i Items) GetItemFlingEffectURL() string {
	return i.Env.URL() + ItemFlingEffectEndpoint
}

// Return a single ItemPocket resource by name or ID
func (i Items) GetItemPocket(nameOrId string) (*models.ItemPocket, error) {
	itemPocket, err := request.GetResource[models.ItemPocket](i.GetItemPocketURL() + nameOrId)
	if err != nil {
		return nil, err
	}
	return itemPocket, nil
}

// Return a list of ItemPocket resource
func (i Items) GetItemPocketList(limit int, offset int) (*models.NamedResourceList, error) {
	itemPocketList, err := request.GetResourceList[models.NamedResourceList](i.GetItemPocketURL(), limit, offset)
	if err != nil {
		return nil, err
	}
	return itemPocketList, nil
}

// Return the ItemPocket resource URL
func (i Items) GetItemPocketURL() string {
	return i.Env.URL() + ItemPocketEndpoint
}
