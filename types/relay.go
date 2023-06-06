package types

import (
	"fmt"
	"reflect"
	"time"
)

type ErrorSource string

const (
	ErrorSourceInternal ErrorSource = "internal"
	ErrorSourceExternal ErrorSource = "external"
)

var (
	// this fields shpould be empty because they are set after db record is created
	shouldBeEmptyRelayField = map[string]bool{
		"RelayID":   true,
		"Session":   true,
		"Region":    true,
		"CreatedAt": true,
		"UpdatedAt": true,
	}

	// TODO: add fields here after discussion
	shouldBeEmptyAltruist = map[string]bool{
		// "PoktNodePublicKey": true,
		// "PoktTxID":          true,
	}

	relayErrorField = map[string]bool{
		"ErrorCode":    true,
		"ErrorName":    true,
		"ErrorMessage": true,
		"ErrorType":    true,
		"ErrorSource":  true,
	}

	relayOptionalFields = map[string]bool{
		"RelaySourceURL":      true,
		"PoktNodeAddress":     true,
		"PoktNodeDomain":      true,
		"GigastakeAppID":      true,
		"RelayChainMethodIDs": true,
		"RelayDataSize":       true,
		"PoktNodePublicKey":   true,
		"PoktTxID":            true,
	}

	validErrorSources = map[string]bool{
		string(ErrorSourceExternal): true,
		string(ErrorSourceInternal): true,
	}
)

type Relay struct {
	RelayID                  int           `json:"relayID"`
	PoktChainID              string        `json:"poktChainID"`
	EndpointID               string        `json:"endpointID"`
	SessionKey               string        `json:"sessionKey"`
	ProtocolAppPublicKey     string        `json:"protocolAppPublicKey"`
	RelaySourceURL           string        `json:"relaySourceUrl"`
	PoktNodeAddress          string        `json:"poktNodeAddress"`
	PoktNodeDomain           string        `json:"poktNodeDomain"`
	PoktNodePublicKey        string        `json:"poktNodePublicKey"`
	RelayStartDatetime       time.Time     `json:"relayStartDatetime"`
	RelayReturnDatetime      time.Time     `json:"relayReturnDatetime"`
	IsError                  bool          `json:"isError"`
	ErrorCode                int           `json:"errorCode,omitempty"`
	ErrorName                string        `json:"errorName,omitempty"`
	ErrorMessage             string        `json:"errorMessage,omitempty"`
	ErrorType                string        `json:"errorType,omitempty"`
	ErrorSource              ErrorSource   `json:"errorSource,omitempty"`
	RelayRoundtripTime       float64       `json:"relayRoundtripTime"`
	RelayChainMethodIDs      []string      `json:"relayChainMethodID"`
	RelayDataSize            int           `json:"relayDataSize"`
	RelayPortalTripTime      float64       `json:"relayPortalTripTime"`
	RelayNodeTripTime        float64       `json:"relayNodeTripTime"`
	RelayURLIsPublicEndpoint bool          `json:"relayUrlIsPublicEndpoint"`
	PortalRegionName         string        `json:"portalRegionName"`
	IsAltruistRelay          bool          `json:"isAltruistRelay"`
	IsUserRelay              bool          `json:"isUserRelay"`
	RequestID                string        `json:"requestID"`
	PoktTxID                 string        `json:"poktTxID"`
	GigastakeAppID           string        `json:"gigastakeAppID"`
	Session                  PocketSession `json:"session"`
	Region                   PortalRegion  `json:"region"`
	CreatedAt                time.Time     `json:"createdAt"`
	UpdatedAt                time.Time     `json:"updatedAt"`
}

func (r *Relay) Validate() (err error) {
	structType := reflect.TypeOf(*r)
	structVal := reflect.ValueOf(*r)
	fieldNum := structVal.NumField()

	// fields are in the order they are declared on the struct
	for i := 0; i < fieldNum; i++ {
		field := structVal.Field(i)
		fieldName := structType.Field(i).Name

		isSet := field.IsValid() && !field.IsZero()

		if isSet {
			// shouldBeEmptyFields should never be set
			// error fields shoould just be set if is an error relay
			if shouldBeEmptyRelayField[fieldName] ||
				(!r.IsError && relayErrorField[fieldName]) ||
				(r.IsAltruistRelay && shouldBeEmptyAltruist[fieldName]) {
				return fmt.Errorf("%s should not be set", fieldName)
			}

			// errorSource field just has some valid error sources
			if fieldName == "ErrorSource" && !validErrorSources[field.String()] {
				return fmt.Errorf("%s is not valid", fieldName)
			}
		}

		if !isSet {
			// shouldBeEmptyField can be empty
			// bools zero value is false which is a valid value
			// error fields can be empty if it is an error relay
			// optional fields can or cannot be set
			// if the relay is an error we should just validate the error fields
			if shouldBeEmptyRelayField[fieldName] ||
				field.Kind() == reflect.Bool ||
				(!r.IsError && relayErrorField[fieldName]) ||
				(r.IsAltruistRelay && shouldBeEmptyAltruist[fieldName]) ||
				relayOptionalFields[fieldName] ||
				(r.IsError && !relayErrorField[fieldName]) {
				continue
			}

			// if is not set and the field is none of the special cases it is an error
			return fmt.Errorf("%s is not set", fieldName)
		}
	}

	return nil
}
