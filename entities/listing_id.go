package entities

import (
	"crypto/rand"
	"errors"
	"regexp"
)

type ListingID struct {
	initialized bool
	id          *string
}

func (wi ListingID) Initialized() bool {
	return wi.initialized
}

func (wi ListingID) String() string {
	return *wi.id
}

func NewListingID(listingID string) (ListingID, error) {

	if len(listingID) != 7 {
		return ListingID{}, errors.New("listing id must exactly be exactly 7 characters long")
	}

	valid, err := regexp.Match("^[A-Z]*$", []byte(listingID))
	if err != nil {
		return ListingID{}, err
	}
	if !valid {
		return ListingID{}, errors.New("listing id is invalid")
	}

	return ListingID{
		id:          &listingID,
		initialized: true,
	}, nil

}

func ListingIDFactory() (ListingID, error) {

	idBytes := make([]byte, 7)

	allowedCharacters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	_, err := rand.Read(idBytes)
	if err != nil {
		return ListingID{}, err
	}

	id := ""
	for _, idByte := range idBytes {
		id += allowedCharacters[int(idByte)%int(len(allowedCharacters))]
	}

	return NewListingID(id)

}
