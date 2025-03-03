package machines

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/internal/constants"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/stretchr/testify/assert"
)

var machines IMachines = NewMachinesGroup()

func TestGetMachine(t *testing.T) {
	rById, _ := machines.GetMachine("1")
	_, err := machines.GetMachine("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Machine resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetMachineList(t *testing.T) {
	rList, _ := machines.GetMachineList(models.PaginationOptions{})
	rPage, _ := machines.GetMachineList(models.PaginationOptions{Limit: 1, Offest: 1})
	assert.Equal(t, constants.MachineEndpoint+"1/", rList.Results[0].URL, "Unexpected URL for Machine resource")
	assert.Equal(t, constants.MachineEndpoint+"2/", rPage.Results[0].URL, "Unexpected URL for Machine resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}
