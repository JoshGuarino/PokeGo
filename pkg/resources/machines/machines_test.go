package machines

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/internal/endpoints"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/stretchr/testify/assert"
)

var machines IMachines = NewMachinesGroup()
var url string = endpoints.BaseURL

func TestNewMachinesGroup(t *testing.T) {
	machines := NewMachinesGroup()
	assert.IsType(t, Machines{}, machines, "Expected Machines instance to be returned")
}

func TestGetMachine(t *testing.T) {
	rById, _ := machines.GetMachine("1")
	_, err := machines.GetMachine("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Machine resource")
	assert.Error(t, err, "Expected an error to be thrown.")
	assert.IsType(t, &models.Machine{}, rById, "Expected Machine instance to be returned")
}

func TestGetMachineList(t *testing.T) {
	rList, _ := machines.GetMachineList(20, 0)
	rPage, _ := machines.GetMachineList(1, 1)
	assert.Equal(t, machines.GetMachineURL()+"1/", rList.Results[0].URL, "Unexpected URL for Machine resource")
	assert.Equal(t, machines.GetMachineURL()+"2/", rPage.Results[0].URL, "Unexpected URL for Machine resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.ResourceList{}, rList, "Expected ResourceList instance to be returned")
}

func TestGetMachineURL(t *testing.T) {
	machineURL := machines.GetMachineURL()
	assert.Equal(t, url+endpoints.Machine, machineURL, "Unexpected Machine resource URL")
	assert.IsType(t, "", machineURL, "Expected Machine resource URL to be a string")
}
