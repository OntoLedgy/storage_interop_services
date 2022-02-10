package object_model

type Files struct {
	*FileSystemObjects
	parentFolder *Folders
}

func (file *Files) initialise(absolutePathString string, parentFolder *Folders) {
	file.FileSystemObjects.initialise(absolutePathString)
	file.addToParent(parentFolder)
}

func (file *Files) addToParent(parentFolder *Folders) {
	if parentFolder == nil {
		return
	}

	file.parentFolder = parentFolder

	parentFolder.addToChildFiles(file)

}
