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
	shouldBeEmptyField = map[string]bool{
		"Session":   true,
		"Region":    true,
		"CreatedAt": true,
		"UpdatedAt": true,
	}

	errorField = map[string]bool{
		"ErrorCode":    true,
		"ErrorName":    true,
		"ErrorMessage": true,
		"ErrorType":    true,
		"ErrorSource":  true,
	}

	validErrorSources = map[string]bool{
		string(ErrorSourceExternal): true,
		string(ErrorSourceInternal): true,
	}
)

type Relay struct {
	PoktChainID              string        `json:"poktChainID"`
	EndpointID               string        `json:"endpointID"`
	SessionKey               string        `json:"sessionKey"`
	RelaySourceURL           string        `json:"relaySourceUrl"`
	PoktNodeAddress          string        `json:"poktNodeAddress"`
	PoktNodeDomain           string        `json:"poktNodeDomain"`
	PoktNodePublicKey        string        `json:"poktNodePublicKey"`
	RelayStartDatetime       time.Time     `json:"relayStartDatetime"`
	RelayReturnDatetime      time.Time     `json:"relayReturnDatetime"`
	IsError                  bool          `json:"isError"` // this field must be before the other error fields for validation to work
	ErrorCode                int           `json:"errorCode,omitempty"`
	ErrorName                string        `json:"errorName,omitempty"`
	ErrorMessage             string        `json:"errorMessage,omitempty"`
	ErrorType                string        `json:"errorType,omitempty"`
	ErrorSource              ErrorSource   `json:"errorSource,omitempty"`
	RelayRoundtripTime       int           `json:"relayRoundtripTime"`
	RelayChainMethodIDs      []string      `json:"relayChainMethodID"`
	RelayDataSize            int           `json:"relayDataSize"`
	RelayPortalTripTime      int           `json:"relayPortalTripTime"`
	RelayNodeTripTime        int           `json:"relayNodeTripTime"`
	RelayURLIsPublicEndpoint bool          `json:"relayUrlIsPublicEndpoint"`
	PortalRegionName         string        `json:"portalRegionName"`
	IsAltruistRelay          bool          `json:"isAltruistRelay"`
	IsUserRelay              bool          `json:"isUserRelay"`
	RequestID                string        `json:"requestID"`
	PoktTxID                 string        `json:"poktTxID"`
	Session                  PocketSession `json:"session"`
	Region                   PortalRegion  `json:"region"`
	CreatedAt                time.Time     `json:"createdAt"`
	UpdatedAt                time.Time     `json:"updatedAt"`
}

func (r Relay) Validate() (err error) {
	structType := reflect.TypeOf(r)
	structVal := reflect.ValueOf(r)
	fieldNum := structVal.NumField()

	var isError bool

	// fields are in the order they are declared on the struct
	for i := 0; i < fieldNum; i++ {
		field := structVal.Field(i)
		fieldName := structType.Field(i).Name

		isSet := field.IsValid() && !field.IsZero()

		if isSet {
			// if isError is set it means it's true so it is an error relay
			if fieldName == "IsError" {
				isError = true
				continue
			}

			// shouldBeEmptyFields should never be set
			// error fields shoould just be set if is an error relay
			if shouldBeEmptyField[fieldName] || (!isError && errorField[fieldName]) {
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
			if shouldBeEmptyField[fieldName] || field.Kind() == reflect.Bool || (!isError && errorField[fieldName]) {
				continue
			}

			// if is not set and the field is none of the special cases it is an error
			return fmt.Errorf("%s is not set", fieldName)
		}
	}

	return nil
}
