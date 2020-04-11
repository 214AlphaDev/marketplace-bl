package entities

import "errors"

type ListingName struct {
	initialized bool
	name        *string
}

func (wn ListingName) Initialized() bool {
	return wn.initialized
}

func (wn ListingName) String() string {
	return *wn.name
}

func NewListingName(name string) (ListingName, error) {

	if len(name) < 5 {
		return ListingName{}, errors.New("listing name is too short")
	}

	return ListingName{
		initialized: true,
		name:        &name,
	}, nil
}
