package object_model

type Folders struct {
	*FileSystemObjects
	ParentFolder *Folders
	ChildFolders []*Folders
	ChildFiles   []*Files
}

//	def __init__(
//		self,
//		absolute_path_string: str,
//		parent_folder: 'Folders' = None):
func (folder *Folders) Initialise(absolute_path_string string, parent_folder *Folders) {

	folder.FileSystemObjects = &FileSystemObjects{}
	//super().__init__(
	folder.FileSystemObjects.Initialise(
		//absolute_path_string=absolute_path_string)
		absolute_path_string)

	//self.parent_folder = \
	folder.ParentFolder =
		//parent_folder
		parent_folder

	//	self.child_folders = \
	//[]
	folder.ChildFolders =
		[]*Folders{}

	//	self.child_files = \
	//	[]
	folder.ChildFiles =
		[]*Files{}

	//	self.__add_to_parent(
	//	parent_folder=parent_folder)
	folder.addToParent(parent_folder)
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

func (folder *Folders) addToChildFolders(
	child_file_system_object *Folders) {

	//self.child_folders.append(
	folder.ChildFolders = append(
		//child_file_system_object)
		folder.ChildFolders, child_file_system_object)

}

//def __add_to_parent(
//self,
func (folder *Folders) addToParent(
	//parent_folder: 'Folders'):
	parent_folder *Folders) {

	//if parent_folder is None:
	if parent_folder == nil {
		//return
		return
	}

	//parent_folder.add_to_child_folders(
	//self)
	parent_folder.addToChildFolders(folder)
}
