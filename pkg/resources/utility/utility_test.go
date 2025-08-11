package utility

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/internal/env"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/stretchr/testify/assert"
)

var utility IUtility = NewUtilityGroup()
var url string = env.ENV.URL()

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
	assert.IsType(t, &models.Language{}, rById, "Expected Language instance to be returned")
}

func TestGetLanguageList(t *testing.T) {
	rList, _ := utility.GetLanguageList(20, 0)
	rPage, _ := utility.GetLanguageList(1, 1)
	assert.Equal(t, "ja-Hrkt", rList.Results[0].Name, "Unexpected Name for Language resource")
	assert.Equal(t, "roomaji", rPage.Results[0].Name, "Unexpected Name for Language resource")
	assert.Equal(t, 1, len(rPage.Results), "Unexpected number of results returned")
	assert.IsType(t, &models.NamedResourceList{}, rList, "Expected NamedResourceList instance to be returned")
}

func TestGetLanguageURL(t *testing.T) {
	languageURL := utility.GetLanguageURL()
	assert.Equal(t, url+LanguageEndpoint, languageURL, "Unexpected Language resource URL")
	assert.IsType(t, "", languageURL, "Expected Language resource URL to be a string")
}
