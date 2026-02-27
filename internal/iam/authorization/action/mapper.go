package action

func toModel(action *ActionEntity) *ActionSchema {
	return &ActionSchema{
		ID:          action.ID,
		Resource:    action.Resource,
		Action:      action.Action,
		Label:       action.Label,
		Description: action.Description,
	}
}

func toDomain(model *ActionSchema) *ActionEntity {
	return &ActionEntity{
		ID:          model.ID,
		Resource:    model.Resource,
		Action:      model.Action,
		Label:       model.Label,
		Description: model.Description,
	}
}
