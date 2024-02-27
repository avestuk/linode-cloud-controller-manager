package linodeid

import (
	"fmt"
	"strconv"
	"strings"
)

const ProviderIDPrefix = "linode://"

type InvalidProviderIDError struct {
	Value string
}

func (e InvalidProviderIDError) Error() string {
	return fmt.Sprintf("invalid provider ID %q", e.Value)
}

func IsLinodeProviderID(providerID string) bool {
	return strings.HasPrefix(providerID, ProviderIDPrefix)
}

func ParseProviderID(providerID string) (int, error) {
	if !IsLinodeProviderID(providerID) {
		return 0, InvalidProviderIDError{providerID}
	}
	id, err := strconv.Atoi(strings.TrimPrefix(providerID, ProviderIDPrefix))
	if err != nil {
		return 0, InvalidProviderIDError{providerID}
	}
	return id, nil
}
