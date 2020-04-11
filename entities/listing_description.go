package entities

import "errors"

type ListingDescription struct {
	initialized bool
	description *string
}

func (wd ListingDescription) Initialized() bool {
	return wd.initialized
}

func (wd ListingDescription) String() string {
	return *wd.description
}

func NewListingDescription(description string) (ListingDescription, error) {

	if len(description) < 30 {
		return ListingDescription{}, errors.New("listing description is too short")
	}

	return ListingDescription{
		initialized: true,
		description: &description,
	}, nil

}
