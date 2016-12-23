package statemanager

import (
	"errors"

	"github.com/hajimehoshi/ebiten"
)

// StateManager manages States by IDs and knows which State is active. It
// runs update/draw for that State.
type StateManager interface {
	// Add a new state to the manager
	Add(state State) error

	// SetActive sets the state to update/draw
	SetActive(id string) error

	// Update and Draw need to be run from the game
	Update() error
	Draw(screen *ebiten.Image) error

	// States is a list of ids that are currently in statemanager
	States() []string
}

// State is a game state
type State interface {
	OnEnter() error
	OnExit() error
	Draw(screen *ebiten.Image) error
	Update() error
	ID() string
}

type stateManager struct {
	states       map[string]State
	keys         []string
	currentState State
}

// New returns a new StateManager
func New() StateManager {
	states := make(map[string]State)

	return &stateManager{
		states: states,
	}
}

func (s *stateManager) Add(state State) error {
	if _, ok := s.states[state.ID()]; ok {
		return errors.New("State with that ID already exists")
	}

	s.states[state.ID()] = state
	s.keys = append(s.keys, state.ID())
	return nil
}

func (s *stateManager) SetActive(id string) error {
	if _, ok := s.states[id]; ok {

		if s.currentState != nil {
			s.currentState.OnExit()
		}

		s.currentState = s.states[id]
		s.currentState.OnEnter()

		return nil
	}

	return errors.New("State does not exist in manager")
}

func (s *stateManager) Update() error {
	if s.currentState == nil {
		return errors.New("No active state set")
	}

	return s.currentState.Update()
}

func (s *stateManager) Draw(screen *ebiten.Image) error {
	return s.currentState.Draw(screen)
}

func (s *stateManager) States() []string {
	return s.keys
}
