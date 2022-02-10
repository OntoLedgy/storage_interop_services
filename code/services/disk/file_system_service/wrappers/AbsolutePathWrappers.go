package wrappers

type AbsolutePathWrappers struct {
	*PathWrappers
}

func (a *AbsolutePathWrappers) initialise(absolutePathString string) {

	a.PathWrappers.initialise(absolutePathString)

}

func (a *AbsolutePathWrappers) AbsolutePathString() string {

	absolutePathString := a.PathWrappers.PathString

	return absolutePathString
}

func (a *AbsolutePathWrappers) AbsoluteLevel() int {

	absoluteLevel := a.PathWrappers.level

	return absoluteLevel
}
