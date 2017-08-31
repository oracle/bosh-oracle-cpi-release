package test

import (
	"github.com/oracle/bosh-oracle-cpi/bmc/stemcell"
	"testing"
)

func Test_StemcellOperations(t *testing.T) {

	state := NewConnectionFixture()
	state.Setup(t)
	defer state.TearDown(t)

	t.Run("CreateStemcell valid image id", func(t *testing.T) {

		creator := stemcell.NewCreator(state.Connector(), state.Logger())
		stemCellId, err := creator.CreateStemcell(state.StemcellImageID())
		assertIsNil(t, err, " Unexpected error from CreateStemcell")
		assertEqual(t, state.StemcellImageID(), stemCellId, "Unexpected imageID as stemcell ID")

	})

	t.Run("CreateStemcell invalid image id", func(t *testing.T) {

		creator := stemcell.NewCreator(state.Connector(), state.Logger())
		_, err := creator.CreateStemcell("non-existing-id")
		assertNotNil(t, err, "Create Stemcell didn't fail as expected")

	})

	t.Run("DeleteStemcell valid image id", func(t *testing.T) {

		destroyer := stemcell.NewDestroyer(state.Connector(), state.Logger())
		err := destroyer.DeleteStemcell(state.StemcellImageID())
		assertIsNil(t, err, "Unexpected failure in DeleteStemcell")

	})
	t.Run("DeleteStemcell invalid image id", func(t *testing.T) {

		destroyer := stemcell.NewDestroyer(state.Connector(), state.Logger())
		err := destroyer.DeleteStemcell(state.StemcellImageID())
		assertIsNil(t, err, "Unexpected failure in DeleteStemcell")

	})

}
