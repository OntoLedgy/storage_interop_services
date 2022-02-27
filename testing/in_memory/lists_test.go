package in_memory

import (
	"container/list"
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/code/services/in_memory/lists"
	"testing"
)

func TestLists(t *testing.T) {

	tests := []struct {
		desc     string
		settings func() []string
		expected func() string
	}{
		{
			desc: "test list iterator",
			settings: func() []string {
				s := []string{
					"test1",
					"test2",
					"test3"}

				return s
			},
			expected: func() string {
				return fmt.Sprintf("select storage technology")
			},
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {

			testList := lists.Lists{
				list.New()}

			testList.AddStringSlice(test.settings())
			var currentElement *list.Element

			for nextElement, hasNextElement := testList.Iterate(); hasNextElement; {

				currentElement, hasNextElement = nextElement()

				if currentElement != nil {
					fmt.Println(currentElement.Value)

				}

			}

			//assert.Equal(t, test.expected(), actual.Error())
		})
	}
}
