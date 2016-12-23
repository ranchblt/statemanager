/*
Package statemanager provides a way to manage states simply.

Create a new statemanager

    stateManager := New()

You will need to create your own states that have a Update, Draw, and ID function following the
State interface. The ID is used to keep track the states and allow you to switch to them. It must
be unique.

Add your states. Add will not allow you to add states with the same ID.

    err := stateManager.Add(menuState)
    err = stateManager.Add(gameState)

Make sure to set the first active state. An error will be returned if the id is not in stateManager

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
