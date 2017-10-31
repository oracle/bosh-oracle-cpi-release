package test

import (
	"github.com/oracle/bosh-oracle-cpi/oci/stemcell"
	"net/url"
	"path"
	"testing"
)

func Test_StemcellOpsFind(t *testing.T) {

	state := NewConnectionFixture()
	state.Setup(t)
	defer state.TearDown(t)

	t.Run("FindStemcell with valid image id", func(t *testing.T) {

		finder := stemcell.NewFinder(state.Connector(), state.Logger())
		stemCellId, err := finder.FindStemcell(state.StemcellImageID())
		assertIsNil(t, err, " Unexpected error from FindStemcell")
		assertEqual(t, state.StemcellImageID(), stemCellId, "Unexpected imageID as stemcell ID")

	})

	t.Run("FindStemcell with invalid image id", func(t *testing.T) {

		finder := stemcell.NewFinder(state.Connector(), state.Logger())
		_, err := finder.FindStemcell("non-existing-id")
		assertNotNil(t, err, "FindStemcell didn't fail as expected")

	})

}

func Test_StemcellOpsDelete(t *testing.T) {

	state := NewConnectionFixture()
	state.Setup(t)
	defer state.TearDown(t)

	t.Run("DeleteStemcell with invalid image id", func(t *testing.T) {

		destroyer := stemcell.NewDestroyer(state.Connector(), state.Logger())
		err := destroyer.DeleteStemcell("non-existing-id")
		assertNotNil(t, err, "DeleteStemcell didn't fail as expected")

	})

}

func Test_StemcellOpsCreateDelete(t *testing.T) {

	state := NewConnectionFixture()
	state.Setup(t)
	defer state.TearDown(t)

	var createdStemcellID string

	t.Run("Delete created stemcell", func(t *testing.T) {
		t.Run("Create stemcell from valid image source url", func(t *testing.T) {

			creator := stemcell.NewCreator(state.Connector(), state.Logger())
			url, _ := url.Parse(state.StemcellImageSourceURI())
			imageName := "test-image-" + path.Base(url.Path)

			stemcellID, err := creator.CreateStemcell(state.StemcellImageSourceURI(), imageName)
			assertIsNil(t, err, "Unexpected failure in CreateStemcell")
			assertNotNil(t, stemcellID, "Unexpected nil stemcellID")

			createdStemcellID = stemcellID

		})

		if createdStemcellID != "" {
			destroyer := stemcell.NewDestroyer(state.Connector(), state.Logger())
			err := destroyer.DeleteStemcell(createdStemcellID)
			assertIsNil(t, err, "Unexpected failure in DeleteStemcell")

		}
	})

}
