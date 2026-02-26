package token

import "gorm.io/datatypes"

func toModel(token *TokenEntity) *TokenSchema {
	var metadata datatypes.JSONMap
	if token.Metadata != nil {
		metadata = datatypes.JSONMap(token.Metadata)
	}

	return &TokenSchema{
		ID:          token.ID,
		ReferenceID: token.ReferenceID,
		Token:       token.Token,
		Resource:    token.Resource,
		ExpiresAt:   token.ExpiresAt,
		Status:      token.Status,
		Metadata:    metadata,
		CreatedAt:   token.CreatedAt,
		UpdatedAt:   token.UpdatedAt,
	}
}

func toDomain(model *TokenSchema) *TokenEntity {
	var metadata map[string]any
	if model.Metadata != nil {
		metadata = map[string]any(model.Metadata)
	}

	return &TokenEntity{
		ID:          model.ID,
		ReferenceID: model.ReferenceID,
		Token:       model.Token,
		Resource:    model.Resource,
		ExpiresAt:   model.ExpiresAt,
		Status:      model.Status,
		Metadata:    metadata,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}
