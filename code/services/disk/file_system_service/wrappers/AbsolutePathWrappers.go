package wrappers

type AbsolutePathWrappers struct {
	PathWrappers
}

func (a *AbsolutePathWrappers) initialise(
	absolutePathString string) {

	a.PathWrappers.Initialise(absolutePathString)

}

func (a *AbsolutePathWrappers) AbsolutePathString() string {

	absolutePathString := a.PathString()

	return absolutePathString
}

func (a *AbsolutePathWrappers) AbsoluteLevel() int {

	absoluteLevel := a.Level()

	return absoluteLevel
}
