package main

import (
	"fmt"
	"testing"
)

//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

func newPlayer() Player {
	return Player{
		name:      "test",
		health:    100,
		maxHealth: 100,
		energy:    100,
		maxEnergy: 500,
	}
}

func TestHealth(test *testing.T) {
	player := newPlayer()
	player.addHealth(999)
	if player.health > player.maxHealth {
		test.Fatalf("Health went over limit: %v, expect to: %v", player.health, player.maxHealth)
	}

	player.applyDamage(player.maxHealth + 1)

	if player.health < 0 {
		test.Fatalf("Health: %v. Minimum: 0", player.health)
	}

	if player.health > player.maxHealth {
		test.Fatalf("Health: %v. Maximum: %v", player.health, player.maxHealth)
	}

	fmt.Println("player: ", player)
}

func TestEnergy(test *testing.T) {
	player := newPlayer()
	player.addEnergy(999)
	if player.energy > player.maxEnergy {
		test.Fatalf("Energy went over limit: %v, expect to: %v", player.energy, player.maxEnergy)
	}

	player.consumeEnergy(player.maxEnergy + 1)

	if player.energy < 0 {
		test.Fatalf("Energy: %v. Minimum: 0", player.energy)
	}

	if player.energy > player.maxEnergy {
		test.Fatalf("Energy: %v. Maximum: %v", player.energy, player.maxEnergy)
	}

	fmt.Println("player: ", player)
}
