package boids

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	position Vector2D
	velocity Vector2D
	id       int
}

func (b *Boid) moveOne() {
	acceleration := b.calcAcceleration()

	rWLock.Lock()
	b.velocity = b.velocity.Add(acceleration).Limit(-1, 1)

	// before moving reset the boid's position to -1
	boidMap[int(b.position.x)][int(b.position.y)] = -1

	b.position = b.position.Add(b.velocity)

	// update the position of boid
	boidMap[int(b.position.x)][int(b.position.y)] = b.id

	rWLock.Unlock()
}

func (b *Boid) calcAcceleration() Vector2D {
	upper, lower := b.position.AddValue(viewRadius), b.position.AddValue(-viewRadius)
	avgPosition, avgVelocity, seperation := Vector2D{0, 0}, Vector2D{0, 0}, Vector2D{0, 0}
	var count = 0.0

	rWLock.RLock()
	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, screenWidth); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, screenHeight); j++ {
			if otherBId := boidMap[int(i)][int(j)]; otherBId != -1 && otherBId != b.id {
				if dist := boids[otherBId].position.Distance(b.position); dist < viewRadius {
					count++
					avgVelocity = avgVelocity.Add(boids[otherBId].velocity)
					avgPosition = avgPosition.Add(boids[otherBId].position)
					seperation = seperation.Add(b.position.Subtract(boids[otherBId].position).DivisionValue(dist))
				}
			}
		}
	}
	rWLock.RUnlock()

	accel := Vector2D{b.borderBounce(b.position.x, screenWidth),
		b.borderBounce(b.position.y, screenHeight)}

	if count > 0 {
		avgVelocity = avgVelocity.DivisionValue(count)
		avgPosition = avgPosition.DivisionValue(count)
		accelAlignment := avgVelocity.Subtract(b.velocity).MultiplyValue(adjRate)
		accelCohesion := avgPosition.Subtract(b.position).MultiplyValue(adjRate)
		accelSeparation := seperation.MultiplyValue(adjRate)
		accel = accel.Add(accelAlignment).Add(accelCohesion).Add(accelSeparation)
	}

	return accel
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func createBoid(id int) {
	b := Boid{
		position: Vector2D{rand.Float64() * screenWidth, rand.Float64() * screenHeight},
		velocity: Vector2D{(rand.Float64() * 2) - 1, (rand.Float64() * 2) - 1},
		id:       id,
	}

	boids[id] = &b
	// add boid's initial position to boidMap
	boidMap[int(b.position.x)][int(b.position.y)] = b.id

	// start each boid in separate thead(Green-Thread/User-Level Threads/Go-Routines)
	go b.start()
}

func (b *Boid) borderBounce(pos, maxBorderPos float64) float64 {
	if pos < viewRadius {
		return 1 / pos
	} else if pos > maxBorderPos-viewRadius {
		return 1 / (pos - maxBorderPos)
	}
	return 0
}
