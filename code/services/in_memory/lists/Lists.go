package lists

import "container/list"

type Lists struct {
	*list.List
}

func (thisList *Lists) Iterate() (f func() (*list.Element, bool), hasNext bool) {
	lengthOfList := thisList.Len()

	loopIndex := lengthOfList

	hasNext = lengthOfList > 0
	element := thisList.Front()

	f = func() (*list.Element, bool) {
		if loopIndex == lengthOfList {
			loopIndex--
			return thisList.Front(), loopIndex > 0
		} else {
			loopIndex--
			element = element.Next()
			return element, loopIndex > 0
		}
	}
	return
}

func (thisList *Lists) AddStringSlice(sliceItems []string) {

	for _, sliceItem := range sliceItems {
		thisList.PushBack(sliceItem)
	}

}

func (thisList *Lists) RemoveElementString(elementName string) {

	for element := thisList.Front(); element != nil; element = element.Next() {
		if element.Value.(string) == elementName {
			thisList.Remove(element)
		}
	}

}

func (thisList *Lists) Contains(elementName string) bool {

	for element := thisList.Front(); element != nil; element = element.Next() {
		if element.Value.(string) == elementName {
			return true
		}
	}
	return false
}
