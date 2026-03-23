package utils

import (
	"fmt"
	"math/rand"
	"time"
)

var adjectives = []string{
	"Brave", "Silent", "Curious", "Lazy", "Crazy",
	"Happy", "Sad", "Angry", "Chill", "Wild",
	"Fierce", "Gentle", "Swift", "Rapid", "Sneaky",
	"Lucky", "Unlucky", "Epic", "Legendary", "Mythic",
	"Dark", "Bright", "Shiny", "Dull", "Mysterious",
	"Ancient", "Modern", "Future", "Retro", "Cosmic",
	"Electric", "Frozen", "Burning", "Stormy", "Windy",
	"Thunder", "Shadow", "Golden", "Silver", "Crimson",
	"Azure", "Neon", "Quantum", "Turbo", "Hyper",
	"Mini", "Mega", "Ultra", "Prime", "Alpha",
}

var nouns = []string{
	"Tiger", "Wolf", "Coder", "Ninja", "Eagle",
	"Dragon", "Panda", "Shark", "Falcon", "Lion",
	"Panther", "Otter", "Fox", "Bear", "Hawk",
	"Rider", "Hunter", "Warrior", "Knight", "Samurai",
	"Wizard", "Ghost", "Phantom", "Shadow", "Spirit",
	"Storm", "Blaze", "Flame", "Frost", "Thunder",
	"Spark", "Pixel", "Byte", "Server", "Hacker",
	"Bot", "AI", "Engine", "Rocket", "Comet",
	"Galaxy", "Planet", "Orbit", "Nova", "Asteroid",
	"Wave", "Echo", "Signal", "Core", "Matrix",
}

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateUsername() string {
	adj := adjectives[seededRand.Intn(len(adjectives))]
	noun := nouns[seededRand.Intn(len(nouns))]
	number := seededRand.Intn(100000) // 0–99,999

	return fmt.Sprintf("%s%s%d", adj, noun, number)
}
