package moves

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/stretchr/testify/assert"
)

var moves IMoves = NewMovesGroup()

func TestNewMovesGroup(t *testing.T) {
	moves := NewMovesGroup()
	assert.IsType(t, Moves{}, moves, "Expected Moves instance to be returned")
}

func TestGetMove(t *testing.T) {
	rById, _ := moves.GetMove("1")
	rByName, _ := moves.GetMove("pound")
	_, err := moves.GetMove("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Move resource")
	assert.Equal(t, "pound", rByName.Name, "Unexpected Name for Move resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetMoveList(t *testing.T) {
	rList, _ := moves.GetMoveList(models.PaginationOptions{})
	rPage, _ := moves.GetMoveList(models.PaginationOptions{Limit: 1, Offest: 1})
	assert.Equal(t, "pound", rList.Results[0].Name, "Unexpected Name for Move resource")
	assert.Equal(t, "karate-chop", rPage.Results[0].Name, "Unexpected Name for Move resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetMoveAilment(t *testing.T) {
	rById, _ := moves.GetMoveAilment("1")
	rByName, _ := moves.GetMoveAilment("unknown")
	_, err := moves.GetMoveAilment("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for MoveAilment resource")
	assert.Equal(t, "unknown", rByName.Name, "Unexpected Name for MoveAilment resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetMoveAilmentList(t *testing.T) {
	rList, _ := moves.GetMoveAilmentList(models.PaginationOptions{})
	rPage, _ := moves.GetMoveAilmentList(models.PaginationOptions{Limit: 1, Offest: 1})
	assert.Equal(t, "unknown", rList.Results[0].Name, "Unexpected Name for MoveAilment resource")
	assert.Equal(t, "none", rPage.Results[0].Name, "Unexpected Name for MoveAilment resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetMoveBattleStyle(t *testing.T) {
	rById, _ := moves.GetMoveBattleStyle("1")
	rByName, _ := moves.GetMoveBattleStyle("attack")
	_, err := moves.GetMoveBattleStyle("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for MoveBattleStyle resource")
	assert.Equal(t, "attack", rByName.Name, "Unexpected Name for MoveBattleStyle resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetMoveBattleStyleList(t *testing.T) {
	rList, _ := moves.GetMoveBattleStyleList(models.PaginationOptions{})
	rPage, _ := moves.GetMoveBattleStyleList(models.PaginationOptions{Limit: 1, Offest: 1})
	assert.Equal(t, "attack", rList.Results[0].Name, "Unexpected Name for MoveBattleStyle resource")
	assert.Equal(t, "defense", rPage.Results[0].Name, "Unexpected Name for MoveBattleStyle resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetMoveCategory(t *testing.T) {
	rById, _ := moves.GetMoveCategory("1")
	rByName, _ := moves.GetMoveCategory("damage")
	_, err := moves.GetMoveCategory("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for MoveCategory resource")
	assert.Equal(t, "damage", rByName.Name, "Unexpected Name for MoveCategory resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetMoveCategoryList(t *testing.T) {
	rList, _ := moves.GetMoveCategoryList(models.PaginationOptions{})
	rPage, _ := moves.GetMoveCategoryList(models.PaginationOptions{Limit: 1, Offest: 1})
	assert.Equal(t, "damage", rList.Results[0].Name, "Unexpected Name for MoveCategory resource")
	assert.Equal(t, "ailment", rPage.Results[0].Name, "Unexpected Name for MoveCategory resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetMoveDamageClass(t *testing.T) {
	rById, _ := moves.GetMoveDamageClass("1")
	rByName, _ := moves.GetMoveDamageClass("status")
	_, err := moves.GetMoveDamageClass("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for MoveDamageClass resource")
	assert.Equal(t, "status", rByName.Name, "Unexpected Name for MoveDamageClass resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetMoveDamageClassList(t *testing.T) {
	rList, _ := moves.GetMoveDamageClassList(models.PaginationOptions{})
	rPage, _ := moves.GetMoveDamageClassList(models.PaginationOptions{Limit: 1, Offest: 1})
	assert.Equal(t, "status", rList.Results[0].Name, "Unexpected Name for MoveDamageClass resource")
	assert.Equal(t, "physical", rPage.Results[0].Name, "Unexpected Name for MoveDamageClass resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetMoveLearnMethod(t *testing.T) {
	rById, _ := moves.GetMoveLearnMethod("1")
	rByName, _ := moves.GetMoveLearnMethod("level-up")
	_, err := moves.GetMoveLearnMethod("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for MoveLearnMethod resource")
	assert.Equal(t, "level-up", rByName.Name, "Unexpected Name for MoveLearnMethod resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetMoveLearnMethodList(t *testing.T) {
	rList, _ := moves.GetMoveLearnMethodList(models.PaginationOptions{})
	rPage, _ := moves.GetMoveLearnMethodList(models.PaginationOptions{Limit: 1, Offest: 1})
	assert.Equal(t, "level-up", rList.Results[0].Name, "Unexpected Name for MoveLearnMethod resource")
	assert.Equal(t, "egg", rPage.Results[0].Name, "Unexpected Name for MoveLearnMethod resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}

func TestGetMoveTarget(t *testing.T) {
	rById, _ := moves.GetMoveTarget("1")
	rByName, _ := moves.GetMoveTarget("specific-move")
	_, err := moves.GetMoveTarget("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for MoveTarget resource")
	assert.Equal(t, "specific-move", rByName.Name, "Unexpected Name for MoveTarget resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetMoveTargetList(t *testing.T) {
	rList, _ := moves.GetMoveTargetList(models.PaginationOptions{})
	rPage, _ := moves.GetMoveTargetList(models.PaginationOptions{Limit: 1, Offest: 1})
	assert.Equal(t, "specific-move", rList.Results[0].Name, "Unexpected Name for MoveTarget resource")
	assert.Equal(t, "selected-pokemon-me-first", rPage.Results[0].Name, "Unexpected Name for MoveTarget resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}
