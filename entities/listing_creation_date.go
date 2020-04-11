package entities

import (
	"errors"
	"time"
)

type ListingCreationDate struct {
	initialized  bool
	creationDate *int64
}

func (cd ListingCreationDate) Initialized() bool {
	return cd.initialized
}

func (cd ListingCreationDate) Time() time.Time {
	return time.Unix(*cd.creationDate, 0)
}

func NewListingCreationDate(creationDate int64) (ListingCreationDate, error) {

	if creationDate <= 0 {
		return ListingCreationDate{}, errors.New("invalid creation date")
	}

	return ListingCreationDate{
		initialized:  true,
		creationDate: &creationDate,
	}, nil

}
