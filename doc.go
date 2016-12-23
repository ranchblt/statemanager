/*
Package statemanager provides a way to manage states simply.

Create a new statemanager

    stateManager := New()

You will need to create your own states that have a Update, Draw, ID, OnExit, and OnEnter function following the
State interface. The ID is used to keep track the states and allow you to switch to them. It must
be unique.

    type menuState struct {}

    // Draw takes a screen and hanles any drawing
    func (s *menuState) Draw(screen *ebiten.Image) error {
        return nil
    }

    // Update Handle any updating this state needs to do.
    func (s *menuState) Update() error {
        return nil
    }

    // ID is the unique ID of this state
    func (s *menuState) ID() string {
        return "Main_Menu"
    }

    // OnEnter does any setup for the state. It could be used to reset a state.
    func (s *menuState) OnEnter() error {
        return nil
    }

    // OnExit handles teardown of this state if required.
    func (s *menuState) OnExit() error {
        return nil
    }

Add your states. Add will not allow you to add states with the same ID.

    err := stateManager.Add(menuState)
    err = stateManager.Add(gameState)

Make sure to set the first active state. An error will be returned if the id is not in stateManager. When
setting the active state the State.OnExit and State.OnEnter methods are run to handle any setup
or teardown.

    err := stateManager.SetActive(menuState.ID())

The states manage themselves so you just need to run the Update and Draw functions correctly in your main
game loop.

    func GameLoop(screen *ebiten.Image) error {
       if err := stateManager.Update(); err != nil {
		    return err
	    }

        if ebiten.IsRunningSlowly() {
            return nil
        }

        if err := stateManager.Draw(screen); err != nil {
		    return err
	    }
    }

*/
package statemanager
