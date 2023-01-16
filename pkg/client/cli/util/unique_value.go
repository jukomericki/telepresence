package util

import "github.com/denisbrodbeck/machineid"

const universalAppId = "dev-app"

func UniqueValue() (string, error) {
	id, err := machineid.ProtectedID(universalAppId)
	if err != nil {
		return "", err
	}

	return id, nil
}
