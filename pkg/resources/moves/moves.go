package moves

import (
	"github.com/JoshGuarino/PokeGo/internal/constants"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Moves group interface
type IMoves interface {
	GetMove(nameOrId string) (*models.Move, error)
	GetMoveList(options models.PaginationOptions) (*models.ResourceList, error)
	GetMoveAilment(nameOrId string) (*models.MoveAilment, error)
	GetMoveAilmentList(options models.PaginationOptions) (*models.ResourceList, error)
	GetMoveBattleStyle(nameOrId string) (*models.MoveBattleStyle, error)
	GetMoveBattleStyleList(options models.PaginationOptions) (*models.ResourceList, error)
	GetMoveCategory(nameOrId string) (*models.MoveCategory, error)
	GetMoveCategoryList(options models.PaginationOptions) (*models.ResourceList, error)
	GetMoveDamageClass(nameOrId string) (*models.MoveDamageClass, error)
	GetMoveDamageClassList(options models.PaginationOptions) (*models.ResourceList, error)
	GetMoveLearnMethod(nameOrId string) (*models.MoveLearnMethod, error)
	GetMoveLearnMethodList(options models.PaginationOptions) (*models.ResourceList, error)
	GetMoveTarget(nameOrId string) (*models.MoveTarget, error)
	GetMoveTargetList(options models.PaginationOptions) (*models.ResourceList, error)
}

// Moves group struct
type Moves struct{}

// Return an instance of Moves resource group struct
func NewMovesGroup() Moves {
	return Moves{}
}

// Return a single Move resource by name or ID
func (m Moves) GetMove(nameOrId string) (*models.Move, error) {
	move, err := request.GetSpecificResource[models.Move](constants.MoveEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return move, nil
}

// Return a list of Move resource
func (m Moves) GetMoveList(options models.PaginationOptions) (*models.ResourceList, error) {
	moveList, err := request.GetResourceList(constants.MoveEndpoint, options)
	if err != nil {
		return nil, err
	}
	return moveList, nil
}

// Return a single MoveAilment resource by name or ID
func (m Moves) GetMoveAilment(nameOrId string) (*models.MoveAilment, error) {
	moveAilment, err := request.GetSpecificResource[models.MoveAilment](constants.MoveAilmentEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return moveAilment, nil
}

// Return a list of MoveAilment resource
func (m Moves) GetMoveAilmentList(options models.PaginationOptions) (*models.ResourceList, error) {
	moveAilmentList, err := request.GetResourceList(constants.MoveAilmentEndpoint, options)
	if err != nil {
		return nil, err
	}
	return moveAilmentList, nil
}

// Return a single MoveBattleStyle resource by name or ID
func (m Moves) GetMoveBattleStyle(nameOrId string) (*models.MoveBattleStyle, error) {
	moveBattleStyle, err := request.GetSpecificResource[models.MoveBattleStyle](constants.MoveBattleStyleEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return moveBattleStyle, nil
}

// Return a list of MoveBattleStyle resource
func (m Moves) GetMoveBattleStyleList(options models.PaginationOptions) (*models.ResourceList, error) {
	moveBattleStyleList, err := request.GetResourceList(constants.MoveBattleStyleEndpoint, options)
	if err != nil {
		return nil, err
	}
	return moveBattleStyleList, nil
}

// Return a single MoveCategory resource by name or ID
func (m Moves) GetMoveCategory(nameOrId string) (*models.MoveCategory, error) {
	moveCategory, err := request.GetSpecificResource[models.MoveCategory](constants.MoveCategoryEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return moveCategory, nil
}

// Return a list of MoveCategory resource
func (m Moves) GetMoveCategoryList(options models.PaginationOptions) (*models.ResourceList, error) {
	moveCategoryList, err := request.GetResourceList(constants.MoveCategoryEndpoint, options)
	if err != nil {
		return nil, err
	}
	return moveCategoryList, nil
}

// Return a single MoveDamageClass resource by name or ID
func (m Moves) GetMoveDamageClass(nameOrId string) (*models.MoveDamageClass, error) {
	moveDamageClass, err := request.GetSpecificResource[models.MoveDamageClass](constants.MoveDamageClassEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return moveDamageClass, nil
}

// Return a list of MoveDamageClass resource
func (m Moves) GetMoveDamageClassList(options models.PaginationOptions) (*models.ResourceList, error) {
	moveDamageClassList, err := request.GetResourceList(constants.MoveDamageClassEndpoint, options)
	if err != nil {
		return nil, err
	}
	return moveDamageClassList, nil
}

// Return a single MoveLearnMethod resource by name or ID
func (m Moves) GetMoveLearnMethod(nameOrId string) (*models.MoveLearnMethod, error) {
	moveLearnMethod, err := request.GetSpecificResource[models.MoveLearnMethod](constants.MoveLearnMethodEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return moveLearnMethod, nil
}

// Return a list of MoveLearnMethod resource
func (m Moves) GetMoveLearnMethodList(options models.PaginationOptions) (*models.ResourceList, error) {
	moveLearnMethodList, err := request.GetResourceList(constants.MoveLearnMethodEndpoint, options)
	if err != nil {
		return nil, err
	}
	return moveLearnMethodList, nil
}

// Return a single MoveTarget resource by name or ID
func (m Moves) GetMoveTarget(nameOrId string) (*models.MoveTarget, error) {
	moveTarget, err := request.GetSpecificResource[models.MoveTarget](constants.MoveTargetEndpoint + nameOrId)
	if err != nil {
		return nil, err
	}
	return moveTarget, nil
}

// Return a list of MoveTarget resource
func (m Moves) GetMoveTargetList(options models.PaginationOptions) (*models.ResourceList, error) {
	moveTargetList, err := request.GetResourceList(constants.MoveTargetEndpoint, options)
	if err != nil {
		return nil, err
	}
	return moveTargetList, nil
}
