package helpers

import (
	repoModel "github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/repository/model"
)

func FilterParts(parts []repoModel.Part, filter repoModel.PartsFilter) []repoModel.Part {
	result := make([]repoModel.Part, 0, len(parts))
	for _, part := range parts {
		// Фильтрация по UUID
		if len(filter.Uuids) > 0 && !containsString(filter.Uuids, part.UUID) {
			continue
		}
		// Фильтрация по имени
		if len(filter.Names) > 0 && !containsString(filter.Names, part.Name) {
			continue
		}
		// Фильтрация по категории
		if len(filter.Categories) > 0 && !containsCategory(filter.Categories, part.Category) {
			continue
		}
		// Фильтрация по стране производителя
		if len(filter.ManufacturerCountries) > 0 {
			if !containsString(filter.ManufacturerCountries, part.Manufacturer.Country) {
				continue
			}
		}
		// Фильтрация по тегам (нужно, чтобы у детали был хотя бы один тег из фильтра)
		if len(filter.Tags) > 0 && !hasIntersection(part.Tags, filter.Tags) {
			continue
		}
		result = append(result, part)
	}
	return result
}

func containsString(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func containsCategory[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func hasIntersection(slice1, slice2 []string) bool {
	set := make(map[string]struct{}, len(slice1))
	for _, s := range slice1 {
		set[s] = struct{}{}
	}
	for _, s := range slice2 {
		if _, ok := set[s]; ok {
			return true
		}
	}
	return false
}
