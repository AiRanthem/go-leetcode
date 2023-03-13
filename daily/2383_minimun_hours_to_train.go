package daily

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) (result int) {
	var energySum, expSum int
	var maxEnergy, maxExp int
	for i := 0; i < len(energy); i++ {
		energySum += energy[i]
		energy[i] = energySum
		expSum += experience[i]
		experience[i] = 2*experience[i] - expSum
		if energy[i] > maxEnergy {
			maxEnergy = energy[i]
		}
		if experience[i] > maxExp {
			maxExp = experience[i]
		}
	}
	maxExp++
	maxEnergy++ // 严格大于
	if maxExp > initialExperience {
		result += maxExp - initialExperience
	}
	if maxEnergy > initialEnergy {
		result += maxEnergy - initialEnergy
	}
	return
}
