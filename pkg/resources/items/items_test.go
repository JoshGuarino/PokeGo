package items

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/stretchr/testify/assert"
)

var items IItems = NewItemsGroup()

func TestGetItem(t *testing.T) {
	rById, _ := items.GetItem("1")
	rByName, _ := items.GetItem("master-ball")
	_, err := items.GetItem("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Item resource")
	assert.Equal(t, "master-ball", rByName.Name, "Unexpected Name for Item resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetItemList(t *testing.T) {
	rList, _ := items.GetItemList(models.PaginationOptions{})
	rPage, _ := items.GetItemList(models.PaginationOptions{Limit: 1, Offest: 1})
	assert.Equal(t, "master-ball", rList.Results[0].Name, "Unexpected Name for Item resource")
	assert.Equal(t, "ultra-ball", rPage.Results[0].Name, "Unexpected Name for Item resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetItemAttribute(t *testing.T) {
	rById, _ := items.GetItemAttribute("1")
	rByName, _ := items.GetItemAttribute("countable")
	_, err := items.GetItemAttribute("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for ItemAttribute resource")
	assert.Equal(t, "countable", rByName.Name, "Unexpected Name for ItemAttribute resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetItemAttributeList(t *testing.T) {
	rList, _ := items.GetItemAttributeList(models.PaginationOptions{})
	rPage, _ := items.GetItemAttributeList(models.PaginationOptions{Limit: 1, Offest: 1})
	assert.Equal(t, "countable", rList.Results[0].Name, "Unexpected Name for ItemAttribute resource")
	assert.Equal(t, "consumable", rPage.Results[0].Name, "Unexpected Name for ItemAttribute resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetItemCategory(t *testing.T) {
	rById, _ := items.GetItemCategory("1")
	rByName, _ := items.GetItemCategory("stat-boosts")
	_, err := items.GetItemCategory("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for ItemCategory resource")
	assert.Equal(t, "stat-boosts", rByName.Name, "Unexpected Name for ItemCategory resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetItemCategoryList(t *testing.T) {
	rList, _ := items.GetItemCategoryList(models.PaginationOptions{})
	rPage, _ := items.GetItemCategoryList(models.PaginationOptions{Limit: 1, Offest: 1})
	assert.Equal(t, "stat-boosts", rList.Results[0].Name, "Unexpected Name for ItemCategory' resource")
	assert.Equal(t, "effort-drop", rPage.Results[0].Name, "Unexpected Name for ItemCategory resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetItemFlingEffect(t *testing.T) {
	rById, _ := items.GetItemFlingEffect("1")
	rByName, _ := items.GetItemFlingEffect("badly-poison")
	_, err := items.GetItemFlingEffect("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for ItemFlingEffect resource")
	assert.Equal(t, "badly-poison", rByName.Name, "Unexpected Name for ItemFlingEffect resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetItemFlingEffectList(t *testing.T) {
	rList, _ := items.GetItemFlingEffectList(models.PaginationOptions{})
	rPage, _ := items.GetItemFlingEffectList(models.PaginationOptions{Limit: 1, Offest: 1})
	assert.Equal(t, "badly-poison", rList.Results[0].Name, "Unexpected Name for ItemFlingEffect resource")
	assert.Equal(t, "burn", rPage.Results[0].Name, "Unexpected Name for ItemFlingEffect resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetItemPocket(t *testing.T) {
	rById, _ := items.GetItemPocket("1")
	rByName, _ := items.GetItemPocket("misc")
	_, err := items.GetItemPocket("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for ItemPocket resource")
	assert.Equal(t, "misc", rByName.Name, "Unexpected Name for ItemPocket resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetItemPocketList(t *testing.T) {
	rList, _ := items.GetItemPocketList(models.PaginationOptions{})
	rPage, _ := items.GetItemPocketList(models.PaginationOptions{Limit: 1, Offest: 1})
	assert.Equal(t, "misc", rList.Results[0].Name, "Unexpected Name for ItemPocket resource")
	assert.Equal(t, "medicine", rPage.Results[0].Name, "Unexpected Name for ItemPocket resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}
