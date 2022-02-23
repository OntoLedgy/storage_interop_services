package testing

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/disk/file_system_service"
	"testing"
)

func TestHashing(t *testing.T) {

	sourceFolder := file_system_service.SelectFolder()

	outputReportFile := file_system_service.CreateFile()

	file_system_service.Get_file_hashes_for_folder(sourceFolder,
		outputReportFile.AbsolutePathString(),
		"sha256",
		0,
		100)

}
