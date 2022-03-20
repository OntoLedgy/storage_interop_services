package disk

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/disk/file_system_service"
	"testing"
)

func TestHashing(t *testing.T) {

	sourceFolder := file_system_service.SelectFolder("select a folder")

	outputReportFile := file_system_service.CreateFile()

	file_system_service.Get_file_hashes_for_folder(sourceFolder.Path.PathString(),
		outputReportFile.AbsolutePathString(),
		"sha256",
		0,
		100)

}
