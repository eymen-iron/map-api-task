package utils

import (
	"github.com/eymen-iron/map-api-task/models"
	validation "github.com/go-ozzo/ozzo-validation"
)

func Validate(l models.Location) error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Latitude, validation.Required, validation.Min(-90.0), validation.Max(90.0)),
		validation.Field(&l.Longitude, validation.Required, validation.Min(-180.0), validation.Max(180.0)),
		validation.Field(&l.Name, validation.Required, validation.Length(1, 255)),
		validation.Field(&l.Marker, validation.Required, validation.Length(3, 18)),
	)
}
