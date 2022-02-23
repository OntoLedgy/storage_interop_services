package object_model

type Folders struct {
	*FileSystemObjects
	ChildFolders []*Folders
	ChildFiles   []*Files
}

func (folder *Folders) Initialise() {

}

func (folder *Folders) addToChildFiles(file *Files) {

	//self,
	//	child_file_system_object: Files):
	//self.child_files.append(
	//	child_file_system_object)
}

//def add_to_child_folders(
//self,
//child_file_system_object: 'Folders'):
//self.child_folders.append(
//child_file_system_object)
