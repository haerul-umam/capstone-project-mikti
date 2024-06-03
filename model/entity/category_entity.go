package entity

import "github.com/haerul-umam/capstone-project-mikti/model/domain"

type CategoryEntity struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ToCategoryEntity(id int, name string) CategoryEntity {
	return CategoryEntity{id, name}
}

func ToCategoryEntities(categorys []domain.Category) []CategoryEntity {
	data := []CategoryEntity{}

	for _, category := range categorys {
		data = append(data, ToCategoryEntity(category.ID, category.Name))
	}

	return data
}
