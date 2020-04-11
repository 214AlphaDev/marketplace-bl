package entities

import (
	"errors"
	"github.com/214alphadev/marketplace-bl/utils"
	. "github.com/214alphadev/marketplace-bl/value_objects"
)

type Listing struct {
	initialized bool
	id          *ListingID
	name        *ListingName
	description *ListingDescription
	creator     *MemberID
	createdAt   *ListingCreationDate
	price       *Price
	photo *[]byte
}

func (i Listing) Name() ListingName {
	return *i.name
}

func (i *Listing) ChangeName(name ListingName) error {

	err := utils.Initialized(name)
	if err != nil {
		return err
	}

	i.name = &name
	return nil

}

func (i Listing) Description() ListingDescription {
	return *i.description
}

func (i *Listing) ChangeDescription(description ListingDescription) error {

	err := utils.Initialized(description)
	if err != nil {
		return err
	}

	i.description = &description
	return nil

}

func (i Listing) Price() Price {
	return *i.price
}

func (i *Listing) ChangePrice(price Price) error {

	if err := utils.Initialized(price); err != nil {
		return err
	}

	i.price = &price
	return nil

}

func (i Listing) Photo() []byte {
	return *i.photo
}

func (i *Listing) ChangePhoto(photo []byte) error {

	if photo == nil || len(photo) == 0 {
		return errors.New("invalid photo")
	}

	i.photo = &photo
	return nil

}

func (i Listing) Seller() MemberID {
	return *i.creator
}

func (i Listing) CreatedAt() ListingCreationDate {
	return *i.createdAt
}

func (i Listing) Initialized() bool {
	return i.initialized
}

func (i Listing) ID() ListingID {
	return *i.id
}

func NewListing(creator MemberID, id ListingID, name ListingName, description ListingDescription, creationDate ListingCreationDate, price Price, photo []byte) (Listing, error) {

	if err := utils.Initialized(creator, id, name, description, creationDate, price); err != nil {
		return Listing{}, err
	}

	if photo == nil || len(photo) == 0 {
		return Listing{}, errors.New("received empty photo")
	}

	return Listing{
		id:          &id,
		name:        &name,
		description: &description,
		initialized: true,
		createdAt:   &creationDate,
		creator:     &creator,
		price:       &price,
		photo:       &photo,
	}, nil

}
