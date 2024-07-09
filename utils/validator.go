package utils

import (
	"github.com/eymen-iron/map-api-task/models"
	validation "github.com/go-ozzo/ozzo-validation"
)

func Validate(l models.Location) error {
	var rules []*validation.FieldRules

	if l.Latitude != 0 {
		rules = append(rules, validation.Field(&l.Latitude, validation.Min(-90.0), validation.Max(90.0), validation.Required))
	}
	if int(l.ID) > 0 {
		rules = append(rules, validation.Field(&l.ID, validation.Min(1), validation.Max(1000), validation.Required))
	}
	if l.Longitude != 0 {
		rules = append(rules, validation.Field(&l.Longitude, validation.Min(-180.0), validation.Max(180.0), validation.Required))
	}
	if l.Name != "" {
		rules = append(rules, validation.Field(&l.Name, validation.Length(1, 255), validation.Required))
	}
	if l.Marker != "" {
		rules = append(rules, validation.Field(&l.Marker, validation.Length(3, 18), validation.Required))
	}

	return validation.ValidateStruct(&l, rules...)
}
