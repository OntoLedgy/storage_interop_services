package wrappers

type PathWrappers struct {
	Path       string //TODO replace with path library
	PathString string
	level      int
}

func (p *PathWrappers) initialise(pathString string) {
	p.Path = pathString
}

//@property
//def base_name(
//self) \
//-> str:
//return \
//str(
//self.__path.name)
//
//@property
//def level(
//self) \
//-> int:
//return \
//len(
//self.__path.parts)
//
//@property
//def path_string(
//self) \
//-> str:
//return \
//str(
//self.__path)
//
//@property
//def parent(self):
//return \
//self.__path.parent
//
//def extend_path(
//self,
//path_extension: str) \
//-> str:
//extended_path_string = \
//self.__path.joinpath(
//path_extension)
//
//return \
//extended_path_string
//
//def exists(
//self) \
//-> bool:
//exists = \
//self.__path.exists()
//
//return \
//exists
//
//def list_of_components(self):
//return \
//self.__path.parts
//
//def item_count(
//self) \
//-> int:
//item_count = \
//len(self.__path.parts)
//
//return \
//item_count
