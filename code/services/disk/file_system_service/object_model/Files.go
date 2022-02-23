package object_model

type Files struct {
	FileSystemObjects
	parentFolder *Folders
}

func (file *Files) Initialise(
	absolutePathString string,
	additionalParentFolder ...*Folders) {

	file.FileSystemObjects.Initialise(absolutePathString)

	if len(additionalParentFolder) > 0 {
		file.addToParent(additionalParentFolder[0])
	}

}

func (file *Files) addToParent(parentFolder *Folders) {

	if parentFolder == nil {
		return
	}

	file.parentFolder = parentFolder

	parentFolder.addToChildFiles(file)

}

func (file *Files) ParentFolder() *Folders {
	return file.parentFolder
}
