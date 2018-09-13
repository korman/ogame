package ogame

// BaseShip base struct for ships
type BaseShip struct {
	BaseDefender
	CargoCapacity    int
	BaseSpeed        int
	FuelConsumption  int
	RapidfireAgainst map[ID]int
	Price            Resources
}

// GetCargoCapacity returns ship cargo capacity
func (b BaseShip) GetCargoCapacity() int {
	return b.CargoCapacity
}

// GetFuelConsumption returns ship fuel consumption
func (b BaseShip) GetFuelConsumption() int {
	return b.FuelConsumption
}

// GetRapidfireAgainst returns which ships/defenses we have rapid fire against
func (b BaseShip) GetRapidfireAgainst() map[ID]int {
	return b.RapidfireAgainst
}

// GetSpeed returns speed of the ship
func (b BaseShip) GetSpeed(techs Researches) int {
	techDriveLvl := 0
	if b.ID == SmallCargoID && techs.ImpulseDrive >= 5 {
		return int(float64(b.BaseSpeed) + (float64(b.BaseSpeed)*0.2)*float64(techs.ImpulseDrive))
	}
	if minLvl, ok := b.Requirements[CombustionDrive.ID]; ok {
		techDriveLvl = techs.CombustionDrive
		if techDriveLvl < minLvl {
			techDriveLvl = minLvl
		}
		return int(float64(b.BaseSpeed) + (float64(b.BaseSpeed)*0.1)*float64(techDriveLvl))
	} else if minLvl, ok := b.Requirements[ImpulseDrive.ID]; ok {
		techDriveLvl = techs.ImpulseDrive
		if techDriveLvl < minLvl {
			techDriveLvl = minLvl
		}
		return int(float64(b.BaseSpeed) + (float64(b.BaseSpeed)*0.2)*float64(techDriveLvl))
	} else if minLvl, ok := b.Requirements[HyperspaceDrive.ID]; ok {
		techDriveLvl = techs.HyperspaceDrive
		if techDriveLvl < minLvl {
			techDriveLvl = minLvl
		}
		return int(float64(b.BaseSpeed) + (float64(b.BaseSpeed)*0.3)*float64(techDriveLvl))
	}
	return int(float64(b.BaseSpeed) + (float64(b.BaseSpeed)*0.2)*float64(techDriveLvl))
}