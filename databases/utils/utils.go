package utils

import (
	"fmt"
	"gopkg.in/myesui/uuid.v1"
)

type UUIDs struct {
	*uuid.UUID
}

func GetUUID(
	uuid_type int,
	seeding_string string) (
	*UUIDs,
	error) {

	var uuid_error error
	generated_uuid := new(UUIDs)

	uuid.SwitchFormatToUpper(uuid.FormatCanonicalCurly)

	generated_uuid.UUID = new(uuid.UUID)

	switch uuid_type {

	case 1:
		// Creating UUIDs Version 4

		*generated_uuid.UUID =
			uuid.NewV4()

	case 2:
		// Parsing UUIDs from string input
		generated_uuid.UUID, uuid_error = uuid.Parse(seeding_string)

		if uuid_error != nil {
			fmt.Printf("Something went wrong: %s", uuid_error)
			return generated_uuid, uuid_error
		}

	}
	return generated_uuid, nil

}

func Change_2d_interface_slice_to_string(two_d_interface [][]interface{}) [][]string {

	var return_string_slice [][]string
	var return_string_row []string

	for _, one_d_interface := range two_d_interface {

		for _, interface_element := range one_d_interface {

			switch interface_element.(type) {
			case int:
				str := fmt.Sprintf("%v", interface_element)
				return_string_row = append(return_string_row, str)
			case string:
				str := fmt.Sprintf("%v", interface_element)
				return_string_row = append(return_string_row, str)
			default:
				str := fmt.Sprintf("%v", interface_element)
				return_string_row = append(return_string_row, str)
			}

		}

		return_string_slice = append(return_string_slice, return_string_row)
		return_string_row = nil
	}

	return return_string_slice

}
