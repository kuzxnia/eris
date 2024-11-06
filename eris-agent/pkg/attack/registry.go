package attack

type AttackRegistry struct {
	attacks map[string]Attack
}

func ProvideAttackRegistry(attacks []Attack) *AttackRegistry {
	attacksMap := map[string]Attack{}
	for _, attack := range attacks {
		if _, preset := attacksMap[attack.AttackType()]; preset {
			panic("attack " + attack.AttackType() + " already registered. Attack types should be unique.")
		}
		attacksMap[attack.AttackType()] = attack
	}

	return &AttackRegistry{
		attacks: attacksMap,
	}
}
