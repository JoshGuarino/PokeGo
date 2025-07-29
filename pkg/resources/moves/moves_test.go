package moves

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/internal/endpoints"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/stretchr/testify/assert"
)

var moves IMoves = NewMovesGroup()
var url string = endpoints.BaseURL

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
	assert.IsType(t, &models.Move{}, rById, "Expected Move instance to be returned")
}

func TestGetMoveList(t *testing.T) {
	rList, _ := moves.GetMoveList(20, 0)
	rPage, _ := moves.GetMoveList(1, 1)
	assert.Equal(t, "pound", rList.Results[0].Name, "Unexpected Name for Move resource")
	assert.Equal(t, "karate-chop", rPage.Results[0].Name, "Unexpected Name for Move resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetMoveURL(t *testing.T) {
	moveURL := moves.GetMoveURL()
	assert.Equal(t, url+endpoints.Move, moveURL, "Unexpected Move resource URL")
	assert.IsType(t, "", moveURL, "Expected Move resource URL to be a string")
}

func TestGetMoveAilment(t *testing.T) {
	rById, _ := moves.GetMoveAilment("1")
	rByName, _ := moves.GetMoveAilment("unknown")
	_, err := moves.GetMoveAilment("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for MoveAilment resource")
	assert.Equal(t, "unknown", rByName.Name, "Unexpected Name for MoveAilment resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.MoveAilment{}, rById, "Expected MoveAilment instance to be returned")
}

func TestGetMoveAilmentList(t *testing.T) {
	rList, _ := moves.GetMoveAilmentList(20, 0)
	rPage, _ := moves.GetMoveAilmentList(1, 1)
	assert.Equal(t, "unknown", rList.Results[0].Name, "Unexpected Name for MoveAilment resource")
	assert.Equal(t, "none", rPage.Results[0].Name, "Unexpected Name for MoveAilment resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetMoveAilmentURL(t *testing.T) {
	moveAilmentURL := moves.GetMoveAilmentURL()
	assert.Equal(t, url+endpoints.MoveAilment, moveAilmentURL, "Unexpected MoveAilment resource URL")
	assert.IsType(t, "", moveAilmentURL, "Expected MoveAilment resource URL to be a string")
}

func TestGetMoveBattleStyle(t *testing.T) {
	rById, _ := moves.GetMoveBattleStyle("1")
	rByName, _ := moves.GetMoveBattleStyle("attack")
	_, err := moves.GetMoveBattleStyle("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for MoveBattleStyle resource")
	assert.Equal(t, "attack", rByName.Name, "Unexpected Name for MoveBattleStyle resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.MoveBattleStyle{}, rById, "Expected MoveBattleStyle instance to be returned")
}

func TestGetMoveBattleStyleList(t *testing.T) {
	rList, _ := moves.GetMoveBattleStyleList(20, 0)
	rPage, _ := moves.GetMoveBattleStyleList(1, 1)
	assert.Equal(t, "attack", rList.Results[0].Name, "Unexpected Name for MoveBattleStyle resource")
	assert.Equal(t, "defense", rPage.Results[0].Name, "Unexpected Name for MoveBattleStyle resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetMoveBattleStyleURL(t *testing.T) {
	moveBattleStyleURL := moves.GetMoveBattleStyleURL()
	assert.Equal(t, url+endpoints.MoveBattleStyle, moveBattleStyleURL, "Unexpected MoveBattleStyle resource URL")
	assert.IsType(t, "", moveBattleStyleURL, "Expected MoveBattleStyle resource URL to be a string")
}

func TestGetMoveCategory(t *testing.T) {
	rById, _ := moves.GetMoveCategory("1")
	rByName, _ := moves.GetMoveCategory("damage")
	_, err := moves.GetMoveCategory("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for MoveCategory resource")
	assert.Equal(t, "damage", rByName.Name, "Unexpected Name for MoveCategory resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.MoveCategory{}, rById, "Expected MoveCategory instance to be returned")
}

func TestGetMoveCategoryList(t *testing.T) {
	rList, _ := moves.GetMoveCategoryList(20, 0)
	rPage, _ := moves.GetMoveCategoryList(1, 1)
	assert.Equal(t, "damage", rList.Results[0].Name, "Unexpected Name for MoveCategory resource")
	assert.Equal(t, "ailment", rPage.Results[0].Name, "Unexpected Name for MoveCategory resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetMoveCategoryURL(t *testing.T) {
	moveCategoryURL := moves.GetMoveCategoryURL()
	assert.Equal(t, url+endpoints.MoveCategory, moveCategoryURL, "Unexpected MoveCategory resource URL")
	assert.IsType(t, "", moveCategoryURL, "Expected MoveCategory resource URL to be a string")
}

func TestGetMoveDamageClass(t *testing.T) {
	rById, _ := moves.GetMoveDamageClass("1")
	rByName, _ := moves.GetMoveDamageClass("status")
	_, err := moves.GetMoveDamageClass("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for MoveDamageClass resource")
	assert.Equal(t, "status", rByName.Name, "Unexpected Name for MoveDamageClass resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.MoveDamageClass{}, rById, "Expected MoveDamageClass instance to be returned")
}

func TestGetMoveDamageClassList(t *testing.T) {
	rList, _ := moves.GetMoveDamageClassList(20, 0)
	rPage, _ := moves.GetMoveDamageClassList(1, 1)
	assert.Equal(t, "status", rList.Results[0].Name, "Unexpected Name for MoveDamageClass resource")
	assert.Equal(t, "physical", rPage.Results[0].Name, "Unexpected Name for MoveDamageClass resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetMoveDamageClassURL(t *testing.T) {
	moveDamageClassURL := moves.GetMoveDamageClassURL()
	assert.Equal(t, url+endpoints.MoveDamageClass, moveDamageClassURL, "Unexpected MoveDamageClass resource URL")
	assert.IsType(t, "", moveDamageClassURL, "Expected MoveDamageClass resource URL to be a string")
}

func TestGetMoveLearnMethod(t *testing.T) {
	rById, _ := moves.GetMoveLearnMethod("1")
	rByName, _ := moves.GetMoveLearnMethod("level-up")
	_, err := moves.GetMoveLearnMethod("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for MoveLearnMethod resource")
	assert.Equal(t, "level-up", rByName.Name, "Unexpected Name for MoveLearnMethod resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.MoveLearnMethod{}, rById, "Expected MoveLearnMethod instance to be returned")
}

func TestGetMoveLearnMethodList(t *testing.T) {
	rList, _ := moves.GetMoveLearnMethodList(20, 0)
	rPage, _ := moves.GetMoveLearnMethodList(1, 1)
	assert.Equal(t, "level-up", rList.Results[0].Name, "Unexpected Name for MoveLearnMethod resource")
	assert.Equal(t, "egg", rPage.Results[0].Name, "Unexpected Name for MoveLearnMethod resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetMoveLearnMethodURL(t *testing.T) {
	moveLearnMethodURL := moves.GetMoveLearnMethodURL()
	assert.Equal(t, url+endpoints.MoveLearnMethod, moveLearnMethodURL, "Unexpected MoveLearnMethod resource URL")
	assert.IsType(t, "", moveLearnMethodURL, "Expected MoveLearnMethod resource URL to be a string")
}

func TestGetMoveTarget(t *testing.T) {
	rById, _ := moves.GetMoveTarget("1")
	rByName, _ := moves.GetMoveTarget("specific-move")
	_, err := moves.GetMoveTarget("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for MoveTarget resource")
	assert.Equal(t, "specific-move", rByName.Name, "Unexpected Name for MoveTarget resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.MoveTarget{}, rById, "Expected MoveTarget instance to be returned")
}

func TestGetMoveTargetList(t *testing.T) {
	rList, _ := moves.GetMoveTargetList(20, 0)
	rPage, _ := moves.GetMoveTargetList(1, 1)
	assert.Equal(t, "specific-move", rList.Results[0].Name, "Unexpected Name for MoveTarget resource")
	assert.Equal(t, "selected-pokemon-me-first", rPage.Results[0].Name, "Unexpected Name for MoveTarget resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetMoveTargetURL(t *testing.T) {
	moveTargetURL := moves.GetMoveTargetURL()
	assert.Equal(t, url+endpoints.MoveTarget, moveTargetURL, "Unexpected MoveTarget resource URL")
	assert.IsType(t, "", moveTargetURL, "Expected MoveTarget resource URL to be a string")
}
