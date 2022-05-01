package main

type Weapon struct {
	On    bool
	Ammo  int
	Power int
}

func (weapon *Weapon) Shoot() bool {
	if weapon.On == false {
		return false
	}
	if weapon.Ammo > 0 {
		weapon.Ammo--
		return true
	} else {
		return false
	}
}

func (weapon *Weapon) RideBike() bool {
	if weapon.On == false {
		return false
	}
	if weapon.Power > 0 {
		weapon.Power--
		return true
	} else {
		return false
	}
}
