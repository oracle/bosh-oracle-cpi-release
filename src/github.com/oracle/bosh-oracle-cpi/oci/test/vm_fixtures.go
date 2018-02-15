package test

import (
	"testing"
)

type VMFixtures struct {
	fixtures []*VMFixture
}

func NewVMFixtures(count int) *VMFixtures {
	arr := make([]*VMFixture, count)
	for i, _ := range arr {
		arr[i] = NewVMFixture()
	}
	return &VMFixtures{fixtures: arr}
}

func (vf *VMFixtures) Setup(t *testing.T) error {
	for _, f := range vf.fixtures {
		if err := f.Setup(t); err != nil {
			return err
		}
	}
	return nil
}

func (vf *VMFixtures) TearDown(t *testing.T) error {
	for _, f := range vf.fixtures {
		if err := f.TearDown(t); err != nil {
			return err
		}
	}
	return nil
}

func (vf *VMFixtures) Fixtures() []*VMFixture {
	return vf.fixtures
}
