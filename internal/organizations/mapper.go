package organizations

func toModel(organization *OrganizationEntity) *OrganizationSchema {
	return &OrganizationSchema{
		ID:        organization.ID,
		Name:      organization.Name,
		Cellphone: organization.Cellphone,
		Logo:      organization.Logo,
		CreatedAt: organization.CreatedAt,
		UpdatedAt: organization.UpdatedAt,
	}
}

func toDomain(model *OrganizationSchema) *OrganizationEntity {
	return &OrganizationEntity{
		ID:        model.ID,
		Name:      model.Name,
		Cellphone: model.Cellphone,
		Logo:      model.Logo,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
