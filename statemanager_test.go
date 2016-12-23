package statemanager

import (
	"testing"

	"github.com/hajimehoshi/ebiten"
	"github.com/stretchr/testify/assert"
)

type testState struct {
	id      string
	counter int
}

func (s *testState) Draw(screen *ebiten.Image) error {
	return nil
}

func (s *testState) Update() error {
	s.counter++
	return nil
}

func (s *testState) ID() string {
	return s.id
}

func newTestState(id string) *testState {
	return &testState{
		id: id,
	}
}

func TestAddValid(t *testing.T) {
	stateManager := New()

	state := newTestState("test")

	err := stateManager.Add(state)

	assert.Nil(t, err)
}

func TestAddFail(t *testing.T) {
	stateManager := New()

	state := newTestState("test")

	err := stateManager.Add(state)
	assert.Nil(t, err)

	err = stateManager.Add(state)
	assert.NotNil(t, err)
}

func TestStates(t *testing.T) {
	stateManager := New()
	state := newTestState("test")
	stateManager.Add(state)

	assert.Len(t, stateManager.States(), 1)
}

func TestAddMultiple(t *testing.T) {
	stateManager := New()

	state := newTestState("test")

	err := stateManager.Add(state)
	assert.Nil(t, err)

	state2 := newTestState("test2")

	err = stateManager.Add(state2)
	assert.Nil(t, err)

	assert.Len(t, stateManager.States(), 2)
}

func TestUpdate(t *testing.T) {
	stateManager := New()
	state := newTestState("test")
	stateManager.Add(state)
	stateManager.SetActive(state.ID())

	err := stateManager.Update()
	assert.Nil(t, err)

	assert.Equal(t, 1, state.counter)
}

func TestUpdateNoActive(t *testing.T) {
	stateManager := New()
	state := newTestState("test")
	stateManager.Add(state)

	err := stateManager.Update()
	assert.NotNil(t, err)
}

func TestSetActive(t *testing.T) {
	stateManager := New()
	state := newTestState("test")
	stateManager.Add(state)
	state2 := newTestState("test2")
	stateManager.Add(state2)

	err := stateManager.SetActive("test")
	assert.Nil(t, err)
}

func TestSetActiveFail(t *testing.T) {
	stateManager := New()
	state := newTestState("test")
	stateManager.Add(state)

	err := stateManager.SetActive("not valid")
	assert.NotNil(t, err)
}
