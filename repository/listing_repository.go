package repository

import (
	. "github.com/214alphadev/marketplace-bl/entities"
)

type IListingRepository interface {
	Save(listing Listing) error
	Get(listingID ListingID) (*Listing, error)
	Query(from *ListingID, next uint32) ([]Listing, error)
	Delete(listingID ListingID) error
}
