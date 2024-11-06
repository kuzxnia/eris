package action

type ActionRegistry struct {
	actions map[string]Action
}

func ProvideActionRegistry(actions []Action) *ActionRegistry {
	actionsByName := map[string]Action{}
	for _, action := range actions {
		if _, isPresent := actionsByName[action.ActionType()]; isPresent {
			panic("action " + action.ActionType() + " already registered. Action types should be unique.")
		}
		actionsByName[action.ActionType()] = action
	}

	return &ActionRegistry{
		actions: actionsByName,
	}
}
