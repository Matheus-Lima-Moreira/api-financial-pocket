package group_permission

func toModel(groupPermission *GroupPermissionEntity) *GroupPermissionSchema {
	return &GroupPermissionSchema{
		ID:        groupPermission.ID,
		Name:      groupPermission.Name,
		Type:      groupPermission.Type,
		CreatedAt: groupPermission.CreatedAt,
		UpdatedAt: groupPermission.UpdatedAt,
	}
}

func toDomain(model *GroupPermissionSchema) *GroupPermissionEntity {
	return &GroupPermissionEntity{
		ID:        model.ID,
		Name:      model.Name,
		Type:      model.Type,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
