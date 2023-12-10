package repositories

import (
	"github.com/bshome19/programming-languages-with-db/configs"
	"github.com/bshome19/programming-languages-with-db/models"
)


type Repos struct{}

var singleton *Repos

func ReposInstance() *Repos {
	if singleton == nil {
		singleton = &Repos{}
	}
	return singleton
}

func (_ *Repos) GetAllLanguages() ([]models.ProgrammingLanguage, error) {
	var languages []models.ProgrammingLanguage
	result := configs.DB.Find(&languages)
	if result.Error != nil {
		return []models.ProgrammingLanguage{}, result.Error
	}
	return languages, nil
}

func (_ *Repos) GetLanguagesById(id string) (models.ProgrammingLanguage, error) {
	var language models.ProgrammingLanguage
	result := configs.DB.Find(&language, "id = ?", id)
	if result.Error != nil {
		return models.ProgrammingLanguage{}, result.Error
	}
	return language, nil
}

func (_ *Repos) CreateNewLanguage(language models.ProgrammingLanguage) (models.ProgrammingLanguage, error) {
	
	result := configs.DB.Create(&language)
	if result.Error != nil {
		return models.ProgrammingLanguage{}, result.Error
	}
	return language, nil
}

func (_ *Repos) DeleteLanguageById(id string) (models.ProgrammingLanguage, error) {
	var language models.ProgrammingLanguage

	result := configs.DB.Delete(&language, "id = ?", id)
	if result.Error != nil {
		return models.ProgrammingLanguage{}, result.Error
	}
	return language, nil

}

func (_ *Repos) UpdateLanguageDataById(updatedlanguage models.ProgrammingLanguage) (models.ProgrammingLanguage, error) {

	id := updatedlanguage.Id
	result := configs.DB.Where("id = ?", id).Updates(&updatedlanguage)
	if result.Error != nil {
		return models.ProgrammingLanguage{}, result.Error
	}
	return updatedlanguage, nil
}