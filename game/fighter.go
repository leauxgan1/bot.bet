package game

import (
	"math/rand"
)

type Fighter struct {
	Health int
	Damage int
	DefenseFlat int
	
	Accuracy float32
	StunChance float32
	CritChance float32
	CritMult float32
	DefenseRatio float32
	DodgeChance float32

}



func (f *Fighter) New(health int, damage int) *Fighter {
	return &Fighter{
		health,
		damage,
		0,
		0.0,
		0.0,
		1.0,
		0.0,
		0.0,
		0.0,
	}
}


func (f *Fighter) Attack(other *Fighter)  {
	// Include stun check on other

	attackRoll := rand.Float32()
	dodgeRoll := rand.Float32()
	if dodgeRoll < other.DodgeChance {
		// Enemy dodged, return with no effect
		// Send message that attack failed
		return
	}

	if attackRoll <= f.Accuracy {
		// Enemy was hit, deal Damage
		// Send message that attack was successful as well as the amt of damage done 
		critRoll := rand.Float32()
		damageDealt := f.Damage
		if critRoll < f.CritChance {
			damageDealt = int(float32(damageDealt) * f.CritMult)
		}
		other.Health -= damageDealt 
	}
}
