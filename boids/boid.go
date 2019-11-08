package main

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

func (b *Boid) calcAcceleration() Vector2D {
	upper, lower := b.position.AddV(viewRadius), b.position.AddV(-viewRadius)
	avgPosition, avgVelocity, separation := Vector2D{0, 0}, Vector2D{0, 0}, Vector2D{0, 0}
	count := 0.0
	//Use read lock to calc new velocities
	rWLock.RLock()
	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, screenWidth); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, screenHeight); j++ {
			if otherBoidId := boidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.id {
				dist := boids[otherBoidId].position.Distance(b.position)
				if dist < viewRadius {
					count++
					avgPosition = avgPosition.Add(boids[otherBoidId].position)
					avgVelocity = avgVelocity.Add(boids[otherBoidId].velocity)
					separation = separation.Add(b.position.Subtract(boids[otherBoidId].position).DivisionV(dist))
				}
			}
		}
	}
	//unlock
	rWLock.RUnlock()
	accel := Vector2D{b.borderBounce(b.position.x, screenWidth), b.borderBounce(b.position.y, screenHeight)}
	if count > 0 {
		avgPosition, avgVelocity = avgPosition.DivisionV(count), avgVelocity.DivisionV(count)
		accelVelocity := avgVelocity.Subtract(b.velocity).MultiplyV(adjRate)
		accelProximity := avgPosition.Subtract(b.position).MultiplyV(adjRate)
		accelSeparation := separation.MultiplyV(adjRate)
		accel = accel.Add(accelProximity).Add(accelVelocity).Add(accelSeparation)
	}
	return accel
}

func (b *Boid) moveOne(acceleration Vector2D) {
	rWLock.Lock()
	//Use write lock to update map
	b.velocity = b.velocity.Add(acceleration).limit(-1, 1)
	boidMap[int(b.position.x)][int(b.position.y)] = -1
	b.position = b.position.Add(b.velocity)
	boidMap[int(b.position.x)][int(b.position.y)] = b.id
	rWLock.Unlock()
}

func (b *Boid) borderBounce(pos, maxBorderPos float64) float64 {
	if pos < viewRadius {
		return 1 / pos
	} else if pos > maxBorderPos-viewRadius {
		return 1 / (pos - maxBorderPos)
	}
	return 0
}

func (b *Boid) start() {
	for {
		b.moveOne(b.calcAcceleration())
		time.Sleep(5 * time.Millisecond)
	}
}

func createBoid(bid int) {
	b := Boid{
		position: Vector2D{rand.Float64() * screenWidth, rand.Float64() * screenHeight},
		velocity: Vector2D{(rand.Float64() * 2) - 1.0, (rand.Float64() * 2) - 1.0},
		id:       bid,
	}
	boids[bid] = &b
	boidMap[int(b.position.x)][int(b.position.y)] = b.id
	go b.start()
}
