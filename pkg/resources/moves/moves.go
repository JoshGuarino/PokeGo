package moves

import (
	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/endpoints"
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/charmbracelet/log"
)

// Moves group interface
type IMoves interface {
	GetMove(nameOrId string) (*models.Move, error)
	GetMoveList(limit int, offset int) (*models.NamedResourceList, error)
	GetMoveURL() string
	GetMoveAilment(nameOrId string) (*models.MoveAilment, error)
	GetMoveAilmentList(limit int, offset int) (*models.NamedResourceList, error)
	GetMoveAilmentURL() string
	GetMoveBattleStyle(nameOrId string) (*models.MoveBattleStyle, error)
	GetMoveBattleStyleList(limit int, offset int) (*models.NamedResourceList, error)
	GetMoveBattleStyleURL() string
	GetMoveCategory(nameOrId string) (*models.MoveCategory, error)
	GetMoveCategoryList(limit int, offset int) (*models.NamedResourceList, error)
	GetMoveCategoryURL() string
	GetMoveDamageClass(nameOrId string) (*models.MoveDamageClass, error)
	GetMoveDamageClassList(limit int, offset int) (*models.NamedResourceList, error)
	GetMoveDamageClassURL() string
	GetMoveLearnMethod(nameOrId string) (*models.MoveLearnMethod, error)
	GetMoveLearnMethodList(limit int, offset int) (*models.NamedResourceList, error)
	GetMoveLearnMethodURL() string
	GetMoveTarget(nameOrId string) (*models.MoveTarget, error)
	GetMoveTargetList(limit int, offset int) (*models.NamedResourceList, error)
	GetMoveTargetURL() string
}

// Moves group struct
type Moves struct {
	MoveURL            string
	MoveAilmentURL     string
	MoveBattleStyleURL string
	MoveCategoryURL    string
	MoveDamageClassURL string
	MoveLearnMethodURL string
	MoveTargetURL      string
	Cache              *cache.Cache
}

// Initialize function
func init() {
	log.Info("Moves resource group initialized")
}

// Return an instance of Moves resource group struct
func NewMovesGroup() Moves {
	url := endpoints.BaseURL
	return Moves{
		MoveURL:            url + endpoints.Move,
		MoveAilmentURL:     url + endpoints.MoveAilment,
		MoveBattleStyleURL: url + endpoints.MoveBattleStyle,
		MoveCategoryURL:    url + endpoints.MoveCategory,
		MoveDamageClassURL: url + endpoints.MoveDamageClass,
		MoveLearnMethodURL: url + endpoints.MoveLearnMethod,
		MoveTargetURL:      url + endpoints.MoveTarget,
		Cache:              cache.C,
	}
}

// Return a single Move resource by name or ID
func (m Moves) GetMove(nameOrId string) (*models.Move, error) {
	move, err := request.GetResource[models.Move](m.MoveURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return move, nil
}

// Return a list of Move resource
func (m Moves) GetMoveList(limit int, offset int) (*models.NamedResourceList, error) {
	moveList, err := request.GetResourceList[models.NamedResourceList](m.MoveURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return moveList, nil
}

// Return the Move resource url
func (m Moves) GetMoveURL() string {
	return m.MoveURL
}

// Return a single MoveAilment resource by name or ID
func (m Moves) GetMoveAilment(nameOrId string) (*models.MoveAilment, error) {
	moveAilment, err := request.GetResource[models.MoveAilment](m.MoveAilmentURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return moveAilment, nil
}

// Return a list of MoveAilment resource
func (m Moves) GetMoveAilmentList(limit int, offset int) (*models.NamedResourceList, error) {
	moveAilmentList, err := request.GetResourceList[models.NamedResourceList](m.MoveAilmentURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return moveAilmentList, nil
}

// Retunr the MoveAilment resource url
func (m Moves) GetMoveAilmentURL() string {
	return m.MoveAilmentURL
}

// Return a single MoveBattleStyle resource by name or ID
func (m Moves) GetMoveBattleStyle(nameOrId string) (*models.MoveBattleStyle, error) {
	moveBattleStyle, err := request.GetResource[models.MoveBattleStyle](m.MoveBattleStyleURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return moveBattleStyle, nil
}

// Return a list of MoveBattleStyle resource
func (m Moves) GetMoveBattleStyleList(limit int, offset int) (*models.NamedResourceList, error) {
	moveBattleStyleList, err := request.GetResourceList[models.NamedResourceList](m.MoveBattleStyleURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return moveBattleStyleList, nil
}

// Return the MoveBattleStyle resource url
func (m Moves) GetMoveBattleStyleURL() string {
	return m.MoveBattleStyleURL
}

// Return a single MoveCategory resource by name or ID
func (m Moves) GetMoveCategory(nameOrId string) (*models.MoveCategory, error) {
	moveCategory, err := request.GetResource[models.MoveCategory](m.MoveCategoryURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return moveCategory, nil
}

// Return a list of MoveCategory resource
func (m Moves) GetMoveCategoryList(limit int, offset int) (*models.NamedResourceList, error) {
	moveCategoryList, err := request.GetResourceList[models.NamedResourceList](m.MoveCategoryURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return moveCategoryList, nil
}

// Return the MoveCategory resource url
func (m Moves) GetMoveCategoryURL() string {
	return m.MoveCategoryURL
}

// Return a single MoveDamageClass resource by name or ID
func (m Moves) GetMoveDamageClass(nameOrId string) (*models.MoveDamageClass, error) {
	moveDamageClass, err := request.GetResource[models.MoveDamageClass](m.MoveDamageClassURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return moveDamageClass, nil
}

// Return a list of MoveDamageClass resource
func (m Moves) GetMoveDamageClassList(limit int, offset int) (*models.NamedResourceList, error) {
	moveDamageClassList, err := request.GetResourceList[models.NamedResourceList](m.MoveDamageClassURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return moveDamageClassList, nil
}

// Return the MoveDamageClass resource url
func (m Moves) GetMoveDamageClassURL() string {
	return m.MoveDamageClassURL
}

// Return a single MoveLearnMethod resource by name or ID
func (m Moves) GetMoveLearnMethod(nameOrId string) (*models.MoveLearnMethod, error) {
	moveLearnMethod, err := request.GetResource[models.MoveLearnMethod](m.MoveLearnMethodURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return moveLearnMethod, nil
}

// Return a list of MoveLearnMethod resource
func (m Moves) GetMoveLearnMethodList(limit int, offset int) (*models.NamedResourceList, error) {
	moveLearnMethodList, err := request.GetResourceList[models.NamedResourceList](m.MoveLearnMethodURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return moveLearnMethodList, nil
}

// Return the MoveLearnMethod resource url
func (m Moves) GetMoveLearnMethodURL() string {
	return m.MoveLearnMethodURL
}

// Return a single MoveTarget resource by name or ID
func (m Moves) GetMoveTarget(nameOrId string) (*models.MoveTarget, error) {
	moveTarget, err := request.GetResource[models.MoveTarget](m.MoveTargetURL + nameOrId)
	if err != nil {
		return nil, err
	}
	return moveTarget, nil
}

// Return a list of MoveTarget resource
func (m Moves) GetMoveTargetList(limit int, offset int) (*models.NamedResourceList, error) {
	moveTargetList, err := request.GetResourceList[models.NamedResourceList](m.MoveTargetURL, limit, offset)
	if err != nil {
		return nil, err
	}
	return moveTargetList, nil
}

// Return the MoveTarget resource url
func (m Moves) GetMoveTargetURL() string {
	return m.MoveTargetURL
}
