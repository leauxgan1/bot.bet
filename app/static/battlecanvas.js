
function randomRoll(minRoll, maxRoll) {
  return Math.floor(minRoll + Math.random() * maxRoll);
}

function passFail(successChance) {
  return Math.random() >= successChance
}

/**
  * Framework fighter, used to battle other frameworks
  * Name : official name of framework
  * Strength : Damage dealt on attack
  * Agility : Ability to dodge
  * MaxHealth : Maximum health value, default health
*/
class Framework {
  constructor(_name,_strength,_agility,_maxHealth) {
    this.name = _name;
    this.strength = _strength;
    this.agility = _agility;
    this.health = _maxHealth;
  }

  attack(other) {
    /** This framework attempts to attack another framework
     * returns the amount of health remaining for the defender
     */
    console.log(this.name + " is attacking " + other.name + "...");
    if (passFail(other.agility / 10)) {
      let damage = this.strength;
      other.health -= damage;
      console.log("Hit for " + damage + " damage!");
    } else {
      console.log("Missed :-(");
    }
    return other.health;
  }
}

let framework1 = new Framework("React",  4, 6,25);
let framework2 = new Framework("Svelte", 7, 4,18);

function gameloop() {
  /**
     * Turn order 
     * Randomly choose a starter (coinflip) 
     * Starter makes attack against other
     * Hits or misses
     * Swap, defender attacking starter
     * If anyone's health is < 1, end game and declare winner
    */

  let coinflip = passFail(0.5);
  let starter;
  let defender;
  let winner;
  if (coinflip) {
    starter = framework1;
    defender = framework2;
  } else {
    starter = framework2;
    defender = framework1;
  }
  while (starter.health > 0 && defender.health > 0) {
    let remainingHP = starter.attack(defender);
    if (remainingHP < 0) {
      winner = defender; 
    } else {
      let temp = defender;
      defender = starter;
      starter = temp;
    }
  }
  console.log("Winner is: " + winner.name);

}


gameloop()
