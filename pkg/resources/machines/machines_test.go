package machines

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/internal/endpoints"
	"github.com/stretchr/testify/assert"
)

var machines IMachines = NewMachinesGroup()

func TestNewMachinesGroup(t *testing.T) {
	machines := NewMachinesGroup()
	assert.IsType(t, Machines{}, machines, "Expected Machines instance to be returned")
}

func TestGetMachine(t *testing.T) {
	rById, _ := machines.GetMachine("1")
	_, err := machines.GetMachine("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Machine resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetMachineList(t *testing.T) {
	rList, _ := machines.GetMachineList(20, 0)
	rPage, _ := machines.GetMachineList(1, 1)
	assert.Equal(t, endpoints.Machine+"1/", rList.Results[0].URL, "Unexpected URL for Machine resource")
	assert.Equal(t, endpoints.Machine+"2/", rPage.Results[0].URL, "Unexpected URL for Machine resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}
