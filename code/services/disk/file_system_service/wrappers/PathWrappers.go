package wrappers

import (
	"github.com/chigopher/pathlib"
	"github.com/spf13/afero"
)

type PathWrappers struct {
	*pathlib.Path
}

func (p *PathWrappers) Initialise(
	pathString string) {
	p.Path =
		pathlib.NewPathAfero(
			pathString,
			afero.NewOsFs())

}

//@property
//def base_name(
//self) \
func (p *PathWrappers) BaseName() string {

	//return \
	//str(
	//self.__path.name)
	return p.Path.Name()
}

//@property
//def level(
//self) \
//-> int:
func (p *PathWrappers) Level() int {
	//return \
	//len(
	//self.__path.parts)
	return p.Level()
}

//@property
//def path_string(
//self) \
//-> str:
func (p *PathWrappers) PathString() string {
	//return \
	//str(
	//self.__path)
	return p.String()
}

//@property
//def parent(self):
func (p *PathWrappers) Parent() *PathWrappers {

	//return \
	//self.__path.parent
	return &PathWrappers{p.Path.Parent()}
}

//def extend_path(
//self,
//path_extension: str) \
func (p *PathWrappers) ExtendPath(
	pathExtension string) string {

	//extended_path_string = \
	extendedPathString :=
		//self.__path.joinpath(
		p.Join(
			//path_extension)
			pathExtension)

	//return \
	//extended_path_string
	return extendedPathString.String()
}

//def exists(
//self) \
//-> bool:
func (p *PathWrappers) Exists() bool {

	//exists = \
	exists, error :=
		//self.__path.exists()
		p.Path.Exists()

	if error != nil {
		panic(error)
	}
	//return \
	//exists
	return exists
}

//def list_of_components(self):
func (p *PathWrappers) ListOfComponents() []string {
	//return \
	//self.__path.parts
	return p.Parts()
}

//def item_count(
//self) \
//-> int:
func (p *PathWrappers) ItemCount() int {

	//item_count = \
	//len(self.__path.parts)
	itemCount := len(p.Parts())

	//return \
	//item_count
	return itemCount
}
