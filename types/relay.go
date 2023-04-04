package types

import (
	"fmt"
	"reflect"
	"time"
)

type ErrorType string

const (
	ErrorTypeSyncCheck  ErrorType = "sync_check"
	ErrorTypeChainCheck ErrorType = "chain_check"
	ErrorTypeRelay      ErrorType = "relay"
)

var (
	// this fields shpould be empty because they are set after db record is created
	shouldBeEmptyField = map[string]bool{
		"RelayID":   true,
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
	}

	validErrorTypes = map[string]bool{
		string(ErrorTypeSyncCheck):  true,
		string(ErrorTypeChainCheck): true,
		string(ErrorTypeRelay):      true,
	}
)

type Relay struct {
	RelayID                  int64         `json:"relayID"`
	ChainID                  int32         `json:"chainID"`
	EndpointID               int32         `json:"endpointID"`
	SessionKey               string        `json:"sessionKey"`
	PoktNodeAddress          string        `json:"poktNodeAddress"`
	RelayStartDatetime       time.Time     `json:"relayStartDatetime"`
	RelayReturnDatetime      time.Time     `json:"relayReturnDatetime"`
	IsError                  bool          `json:"isError"` // this field must be before the other error fields for validation to work
	ErrorCode                int32         `json:"errorCode,omitempty"`
	ErrorName                string        `json:"errorName,omitempty"`
	ErrorMessage             string        `json:"errorMessage,omitempty"`
	ErrorType                ErrorType     `json:"errorType,omitempty"`
	RelayRoundtripTime       int32         `json:"relayRoundtripTime"`
	RelayChainMethodID       int32         `json:"relayChainMethodID"`
	RelayDataSize            int32         `json:"relayDataSize"`
	RelayPortalTripTime      int32         `json:"relayPortalTripTime"`
	RelayNodeTripTime        int32         `json:"relayNodeTripTime"`
	RelayURLIsPublicEndpoint bool          `json:"relayUrlIsPublicEndpoint"`
	PortalOriginRegionID     int32         `json:"portalOriginRegionID"`
	IsAltruistRelay          bool          `json:"isAltruistRelay"`
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

			// errorType field just has some valid error types
			if fieldName == "ErrorType" && !validErrorTypes[field.String()] {
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