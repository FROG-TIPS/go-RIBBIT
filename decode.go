package ribbit

import (
	"encoding/asn1"
	"log"
)

type val interface{}

func DecodeTips(encodedCroak []byte) (tips []Tip, err error) {
	rest, err := asn1.Unmarshal(encodedCroak, &tips)
	if err != nil {
		log.Printf("Error! err: %v rest: %v", err, rest)
	}

	return tips, err
}

func DecodeTip(encodedCroak []byte) (tip Tip, err error) {
	rest, err := asn1.Unmarshal(encodedCroak, &tip)
	if err != nil {
		log.Printf("Error! err: %v rest: %v", err, rest)
	}
	return
}