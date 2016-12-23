# StateManager [![GoDoc](https://godoc.org/github.com/ranchblt/statemanager?status.svg)](https://godoc.org/github.com/ranchblt/statemanager)
StateManager is simple state management for ebiten games

## Documentation

Read [GoDoc](https://godoc.org/github.com/ranchblt/statemanager)

## Usage
```Go
    stateManager := New()
    stateManager.Add(menuState)
    stateManager.SetActive(menuState.ID())

    // Main game loop function
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
```

## License
[MIT License](LICENSE)