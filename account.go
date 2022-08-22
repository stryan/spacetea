package main

import sim "git.saintnet.tech/stryan/spacetea/simulator"

//Account is a account
type Account struct {
	Player    *sim.Player
	Simulator *sim.Simulator
	Landmark  *sim.Landmark
}
