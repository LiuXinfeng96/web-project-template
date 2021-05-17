package services

import uuid "github.com/satori/go.uuid"

//Getuuid get uuid
func Getuuid() string {
	uuid := uuid.NewV4()
	return uuid.String()
}
