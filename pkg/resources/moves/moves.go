package moves

import (
	"fmt"

	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/endpoints"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Moves group interface
type IMoves interface {
	GetMove(nameOrId string) (*models.Move, error)
	GetMoveList(limit int, offset int) (*models.NamedResourceList, error)
	GetMoveAilment(nameOrId string) (*models.MoveAilment, error)
	GetMoveAilmentList(limit int, offset int) (*models.NamedResourceList, error)
	GetMoveBattleStyle(nameOrId string) (*models.MoveBattleStyle, error)
	GetMoveBattleStyleList(limit int, offset int) (*models.NamedResourceList, error)
	GetMoveCategory(nameOrId string) (*models.MoveCategory, error)
	GetMoveCategoryList(limit int, offset int) (*models.NamedResourceList, error)
	GetMoveDamageClass(nameOrId string) (*models.MoveDamageClass, error)
	GetMoveDamageClassList(limit int, offset int) (*models.NamedResourceList, error)
	GetMoveLearnMethod(nameOrId string) (*models.MoveLearnMethod, error)
	GetMoveLearnMethodList(limit int, offset int) (*models.NamedResourceList, error)
	GetMoveTarget(nameOrId string) (*models.MoveTarget, error)
	GetMoveTargetList(limit int, offset int) (*models.NamedResourceList, error)
}

// Moves group struct
type Moves struct {
	Cache *cache.Cache
}

// Initialize function
func init() {
	fmt.Println("Moves resource group initialized")
}

// Return an instance of Moves resource group struct
func NewMovesGroup() Moves {
	return Moves{
		Cache: cache.C,
	}
}

// Return a single Move resource by name or ID
func (m Moves) GetMove(nameOrId string) (*models.Move, error) {
	move, err := request.GetResource[models.Move](endpoints.Move + nameOrId)
	if err != nil {
		return nil, err
	}
	return move, nil
}

// Return a list of Move resource
func (m Moves) GetMoveList(limit int, offset int) (*models.NamedResourceList, error) {
	moveList, err := request.GetResourceList[models.NamedResourceList](endpoints.Move, limit, offset)
	if err != nil {
		return nil, err
	}
	return moveList, nil
}

// Return a single MoveAilment resource by name or ID
func (m Moves) GetMoveAilment(nameOrId string) (*models.MoveAilment, error) {
	moveAilment, err := request.GetResource[models.MoveAilment](endpoints.MoveAilment + nameOrId)
	if err != nil {
		return nil, err
	}
	return moveAilment, nil
}

// Return a list of MoveAilment resource
func (m Moves) GetMoveAilmentList(limit int, offset int) (*models.NamedResourceList, error) {
	moveAilmentList, err := request.GetResourceList[models.NamedResourceList](endpoints.MoveAilment, limit, offset)
	if err != nil {
		return nil, err
	}
	return moveAilmentList, nil
}

// Return a single MoveBattleStyle resource by name or ID
func (m Moves) GetMoveBattleStyle(nameOrId string) (*models.MoveBattleStyle, error) {
	moveBattleStyle, err := request.GetResource[models.MoveBattleStyle](endpoints.MoveBattleStyle + nameOrId)
	if err != nil {
		return nil, err
	}
	return moveBattleStyle, nil
}

// Return a list of MoveBattleStyle resource
func (m Moves) GetMoveBattleStyleList(limit int, offset int) (*models.NamedResourceList, error) {
	moveBattleStyleList, err := request.GetResourceList[models.NamedResourceList](endpoints.MoveBattleStyle, limit, offset)
	if err != nil {
		return nil, err
	}
	return moveBattleStyleList, nil
}

// Return a single MoveCategory resource by name or ID
func (m Moves) GetMoveCategory(nameOrId string) (*models.MoveCategory, error) {
	moveCategory, err := request.GetResource[models.MoveCategory](endpoints.MoveCategory + nameOrId)
	if err != nil {
		return nil, err
	}
	return moveCategory, nil
}

// Return a list of MoveCategory resource
func (m Moves) GetMoveCategoryList(limit int, offset int) (*models.NamedResourceList, error) {
	moveCategoryList, err := request.GetResourceList[models.NamedResourceList](endpoints.MoveCategory, limit, offset)
	if err != nil {
		return nil, err
	}
	return moveCategoryList, nil
}

// Return a single MoveDamageClass resource by name or ID
func (m Moves) GetMoveDamageClass(nameOrId string) (*models.MoveDamageClass, error) {
	moveDamageClass, err := request.GetResource[models.MoveDamageClass](endpoints.MoveDamageClass + nameOrId)
	if err != nil {
		return nil, err
	}
	return moveDamageClass, nil
}

// Return a list of MoveDamageClass resource
func (m Moves) GetMoveDamageClassList(limit int, offset int) (*models.NamedResourceList, error) {
	moveDamageClassList, err := request.GetResourceList[models.NamedResourceList](endpoints.MoveDamageClass, limit, offset)
	if err != nil {
		return nil, err
	}
	return moveDamageClassList, nil
}

// Return a single MoveLearnMethod resource by name or ID
func (m Moves) GetMoveLearnMethod(nameOrId string) (*models.MoveLearnMethod, error) {
	moveLearnMethod, err := request.GetResource[models.MoveLearnMethod](endpoints.MoveLearnMethod + nameOrId)
	if err != nil {
		return nil, err
	}
	return moveLearnMethod, nil
}

// Return a list of MoveLearnMethod resource
func (m Moves) GetMoveLearnMethodList(limit int, offset int) (*models.NamedResourceList, error) {
	moveLearnMethodList, err := request.GetResourceList[models.NamedResourceList](endpoints.MoveLearnMethod, limit, offset)
	if err != nil {
		return nil, err
	}
	return moveLearnMethodList, nil
}

// Return a single MoveTarget resource by name or ID
func (m Moves) GetMoveTarget(nameOrId string) (*models.MoveTarget, error) {
	moveTarget, err := request.GetResource[models.MoveTarget](endpoints.MoveTarget + nameOrId)
	if err != nil {
		return nil, err
	}
	return moveTarget, nil
}

// Return a list of MoveTarget resource
func (m Moves) GetMoveTargetList(limit int, offset int) (*models.NamedResourceList, error) {
	moveTargetList, err := request.GetResourceList[models.NamedResourceList](endpoints.MoveTarget, limit, offset)
	if err != nil {
		return nil, err
	}
	return moveTargetList, nil
}
