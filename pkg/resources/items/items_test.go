package items

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/internal/env"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/stretchr/testify/assert"
)

var items IItems = NewItemsGroup()
var url string = env.ENV.URL()

func TestNewItemsGroup(t *testing.T) {
	items := NewItemsGroup()
	assert.IsType(t, Items{}, items, "Expected Items instance to be returned")
}

func TestGetItem(t *testing.T) {
	rById, _ := items.GetItem("1")
	rByName, _ := items.GetItem("master-ball")
	_, err := items.GetItem("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Item resource")
	assert.Equal(t, "master-ball", rByName.Name, "Unexpected Name for Item resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.Item{}, rById, "Expected Item instance to be returned")
}

func TestGetItemList(t *testing.T) {
	rList, _ := items.GetItemList(20, 0)
	rPage, _ := items.GetItemList(1, 1)
	assert.Equal(t, "master-ball", rList.Results[0].Name, "Unexpected Name for Item resource")
	assert.Equal(t, "ultra-ball", rPage.Results[0].Name, "Unexpected Name for Item resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetItemURL(t *testing.T) {
	itemURL := items.GetItemURL()
	assert.Equal(t, url+ItemEndpoint, itemURL, "Unexpected Item resource URL")
	assert.IsType(t, "", itemURL, "Expected Item resource URL to be a string")
}

func TestGetItemAttribute(t *testing.T) {
	rById, _ := items.GetItemAttribute("1")
	rByName, _ := items.GetItemAttribute("countable")
	_, err := items.GetItemAttribute("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for ItemAttribute resource")
	assert.Equal(t, "countable", rByName.Name, "Unexpected Name for ItemAttribute resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.ItemAttribute{}, rById, "Expected ItemAttribute instance to be returned")
}

func TestGetItemAttributeList(t *testing.T) {
	rList, _ := items.GetItemAttributeList(20, 0)
	rPage, _ := items.GetItemAttributeList(1, 1)
	assert.Equal(t, "countable", rList.Results[0].Name, "Unexpected Name for ItemAttribute resource")
	assert.Equal(t, "consumable", rPage.Results[0].Name, "Unexpected Name for ItemAttribute resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetItemAttributeURL(t *testing.T) {
	itemAttributeURL := items.GetItemAttributeURL()
	assert.Equal(t, url+ItemAttributeEndpoint, itemAttributeURL, "Unexpected ItemAttribute resource URL")
	assert.IsType(t, "", itemAttributeURL, "Expected ItemAttribute resource URL to be a string")
}

func TestGetItemCategory(t *testing.T) {
	rById, _ := items.GetItemCategory("1")
	rByName, _ := items.GetItemCategory("stat-boosts")
	_, err := items.GetItemCategory("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for ItemCategory resource")
	assert.Equal(t, "stat-boosts", rByName.Name, "Unexpected Name for ItemCategory resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.ItemCategory{}, rById, "Expected ItemCategory instance to be returned")
}

func TestGetItemCategoryList(t *testing.T) {
	rList, _ := items.GetItemCategoryList(20, 0)
	rPage, _ := items.GetItemCategoryList(1, 1)
	assert.Equal(t, "stat-boosts", rList.Results[0].Name, "Unexpected Name for ItemCategory' resource")
	assert.Equal(t, "effort-drop", rPage.Results[0].Name, "Unexpected Name for ItemCategory resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetItemCategoryURL(t *testing.T) {
	itemCategoryURL := items.GetItemCategoryURL()
	assert.Equal(t, url+ItemCategoryEndpoint, itemCategoryURL, "Unexpected ItemCategory resource URL")
	assert.IsType(t, "", itemCategoryURL, "Expected ItemCategory resource URL to be a string")
}

func TestGetItemFlingEffect(t *testing.T) {
	rById, _ := items.GetItemFlingEffect("1")
	rByName, _ := items.GetItemFlingEffect("badly-poison")
	_, err := items.GetItemFlingEffect("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for ItemFlingEffect resource")
	assert.Equal(t, "badly-poison", rByName.Name, "Unexpected Name for ItemFlingEffect resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.ItemFlingEffect{}, rById, "Expected ItemFlingEffect instance to be returned")
}

func TestGetItemFlingEffectList(t *testing.T) {
	rList, _ := items.GetItemFlingEffectList(20, 0)
	rPage, _ := items.GetItemFlingEffectList(1, 1)
	assert.Equal(t, "badly-poison", rList.Results[0].Name, "Unexpected Name for ItemFlingEffect resource")
	assert.Equal(t, "burn", rPage.Results[0].Name, "Unexpected Name for ItemFlingEffect resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetItemFlingEffectURL(t *testing.T) {
	itemFlingEffectURL := items.GetItemFlingEffectURL()
	assert.Equal(t, url+ItemFlingEffectEndpoint, itemFlingEffectURL, "Unexpected ItemFlingEffect resource URL")
	assert.IsType(t, "", itemFlingEffectURL, "Expected ItemFlingEffect resource URL to be a string")
}

func TestGetItemPocket(t *testing.T) {
	rById, _ := items.GetItemPocket("1")
	rByName, _ := items.GetItemPocket("misc")
	_, err := items.GetItemPocket("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for ItemPocket resource")
	assert.Equal(t, "misc", rByName.Name, "Unexpected Name for ItemPocket resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.ItemPocket{}, rById, "Expected ItemPocket instance to be returned")
}

func TestGetItemPocketList(t *testing.T) {
	rList, _ := items.GetItemPocketList(20, 0)
	rPage, _ := items.GetItemPocketList(1, 1)
	assert.Equal(t, "misc", rList.Results[0].Name, "Unexpected Name for ItemPocket resource")
	assert.Equal(t, "medicine", rPage.Results[0].Name, "Unexpected Name for ItemPocket resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetItemPocketURL(t *testing.T) {
	itemPocketURL := items.GetItemPocketURL()
	assert.Equal(t, url+ItemPocketEndpoint, itemPocketURL, "Unexpected ItemPocket resource URL")
	assert.IsType(t, "", itemPocketURL, "Expected ItemPocket resource URL to be a string")
}
