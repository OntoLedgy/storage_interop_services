package lists

import "container/list"

type Lists struct {
	list.List
}

func (list *Lists) AddStringSlice(sliceItems []string) {

	for _, sliceItem := range sliceItems {
		list.PushBack(sliceItem)
	}

}

func (list *Lists) RemoveElementString(elementName string) {

	for element := list.Front(); element != nil; element = element.Next() {
		if element.Value.(string) == elementName {
			list.Remove(element)
		}
	}

}

func (list *Lists) Contains(elementName string) bool {

	for element := list.Front(); element != nil; element = element.Next() {
		if element.Value.(string) == elementName {
			return true
		}
	}
	return false
}
