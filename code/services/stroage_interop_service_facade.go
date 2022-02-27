package services

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model/configurations"
)

type StorageServices struct {
	StorageServiceSettings string
	StorageTechnologyType  string //database, in-memory, disk
	DatabaseSettings       configurations.Settings
	errorDetails           string
}

func (storageServices *StorageServices) StoreData(dataToStore interface{}) error {

	switch storageServices.StorageTechnologyType {
	case "database":
		StoreDataInDatabase(dataToStore)
		return nil
	case "in-memory":
		StoreDataInMemory(dataToStore)
	case "disk":
		StoreDataInDisk(dataToStore)

	default:
		storageServices.errorDetails = "select storage technology"

	}

	return storageServices
}

func StoreDataInDisk(store interface{}) {

}

func StoreDataInMemory(store interface{}) {

}

func StoreDataInDatabase(store interface{}) {

}

func (storageServices *StorageServices) Error() string {

	return storageServices.errorDetails
}
