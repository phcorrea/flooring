package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestRepositoryCreate(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)

	repo := NewRepository()
	repo.db.AutoMigrate(&modelTest{})

	testModel := &modelTest{StringAttribute: "not_empty"}
	err := repo.Create(testModel)

	assert.NotEqual(t, uint(0), testModel.ID, "Model not persisted")
	assert.Equal(t, "not_empty", testModel.StringAttribute, "Model's stringAttribute not persisted")
	assert.Equal(t, nil, err, "Error persisting the model")
}

func TestRepositoryFindBy(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)

	repo := NewRepository()
	repo.db.AutoMigrate(&modelTest{})

	createdModel := &modelTest{StringAttribute: "not_empty"}
	err := repo.Create(createdModel)
	assert.Equal(t, nil, err, "Error persisting the model")

	readModel := &modelTest{}
	err = repo.FindBy(readModel, "id = ?", createdModel.ID)

	assert.Equal(t, createdModel.ID, readModel.ID, "Read Model with different ID")
	assert.Equal(t, createdModel.StringAttribute, readModel.StringAttribute, "Model's StringAttribute not read sucessfully")
	assert.Equal(t, nil, err, "Error persisting the model")
}

func TestRepositorySearch(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)

	repo := NewRepository()
	createdModel := &modelTest{StringAttribute: "randomstring"}
	err := repo.Create(createdModel)
	assert.Equal(t, nil, err, "Error persisting the model")

	readModels := []modelTest{}
	err = repo.FindBy(&readModels, "string_attribute = ?", createdModel.StringAttribute)

	assert.Equal(t, createdModel.ID, readModels[0].ID, "Read Model with different ID")
	assert.Equal(t, createdModel.StringAttribute, readModels[0].StringAttribute, "Model's StringAttribute not read sucessfully")
	assert.Equal(t, nil, err, "Error persisting the model")
}

type modelTest struct {
	gorm.Model
	StringAttribute string
}

func setupTest(t *testing.T) func(t *testing.T) {
	repo := NewRepository()
	repo.db.AutoMigrate(&modelTest{})

	return func(t *testing.T) {
		repo := NewRepository()
		repo.db.Migrator().DropTable(&modelTest{})
	}
}
