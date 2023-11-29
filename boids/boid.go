package boids

import "math/rand"

type Boid struct {
	position Vector2D
	velocity Vector2D
	id       int
}

func createBoid(id int) {
	b := Boid{
		position: Vector2D{rand.Float64() * screenWidth, rand.Float64() * screenHeight},
		velocity: Vector2D{(rand.Float64() * 2) - 1, (rand.Float64() * 2) - 1},
		id:       id,
	}

	boids[id] = &b
	go b.start()
}
