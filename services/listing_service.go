package services

import (
	"errors"
	. "github.com/214alphadev/marketplace-bl/entities"
	. "github.com/214alphadev/marketplace-bl/repository"
	"github.com/214alphadev/marketplace-bl/utils"
	. "github.com/214alphadev/marketplace-bl/value_objects"
	"time"
)

type IListingService interface {
	List(memberID MemberID, name ListingName, description ListingDescription, price Price, photo []byte) (ListingID, error)
	Listings(start *ListingID, next uint32) ([]Listing, error)
	GetByID(listingID ListingID) (*Listing, error)
	Update(listing Listing) error
	Delete(listing ListingID) error
}

type listingService struct {
	listingRepository IListingRepository
}

func (ws *listingService) List(memberID MemberID, name ListingName, description ListingDescription, price Price, photo []byte) (ListingID, error) {

	if err := utils.Initialized(memberID, name, description); err != nil {
		return ListingID{}, err
	}

	id, err := ListingIDFactory()
	if err != nil {
		return ListingID{}, err
	}

	creationDate, err := NewListingCreationDate(time.Now().Unix())
	if err != nil {
		return ListingID{}, err
	}

	listing, err := NewListing(memberID, id, name, description, creationDate, price, photo)
	if err != nil {
		return ListingID{}, err
	}

	if err := ws.listingRepository.Save(listing); err != nil {
		return ListingID{}, err
	}

	return id, nil

}

func (ws *listingService) Listings(start *ListingID, next uint32) ([]Listing, error) {

	if err := utils.Initialized(start); err != nil {
		return nil, err
	}

	return ws.listingRepository.Query(start, next)

}

func (ws *listingService) GetByID(listingID ListingID) (*Listing, error) {

	w, err := ws.listingRepository.Get(listingID)

	if err != nil {
		return nil, err
	}

	return w, nil

}

func (ws *listingService) Delete(listing ListingID) error {

	if err := utils.Initialized(listing); err != nil {
		return err
	}

	fetchedListing, err := ws.listingRepository.Get(listing)
	if err != nil {
		return err
	}
	if fetchedListing == nil {
		return errors.New("ListingDoesNotExist")
	}

	return ws.listingRepository.Delete(listing)

}

func (ws *listingService) Update(listing Listing) error {

	if err := utils.Initialized(listing); err != nil {
		return err
	}

	fetchedListing, err := ws.listingRepository.Get(listing.ID())
	if err != nil {
		return err
	}
	if fetchedListing == nil {
		return errors.New("ListingDoesNotExist")
	}

	return ws.listingRepository.Save(listing)

}

func NewListingService(listingRepository IListingRepository) IListingService {
	return &listingService{listingRepository: listingRepository}
}
