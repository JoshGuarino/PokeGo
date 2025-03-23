package utility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var utility IUtility = NewUtilityGroup()

func TestNewUtilityGroup(t *testing.T) {
	utility := NewUtilityGroup()
	assert.IsType(t, Utility{}, utility, "Expected Utility instance to be returned")
}

func TestGetLanguage(t *testing.T) {
	rById, _ := utility.GetLanguage("1")
	rByname, _ := utility.GetLanguage("en")
	_, err := utility.GetLanguage("test")
	assert.Equal(t, 1, rById.ID, "Unexpected ID for Language resource")
	assert.Equal(t, "en", rByname.Name, "Unexpected ID for Language resource")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetLanguageList(t *testing.T) {
	rList, _ := utility.GetLanguageList(20, 0)
	rPage, _ := utility.GetLanguageList(1, 1)
	assert.Equal(t, "ja-Hrkt", rList.Results[0].Name, "Unexpected Name for Language resource")
	assert.Equal(t, "roomaji", rPage.Results[0].Name, "Unexpected Name for Language resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
}
