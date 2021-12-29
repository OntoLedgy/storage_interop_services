package storage

import "encoding/json"

func Extract_columns_from_2d_slices(original_slice [][]string, columns_to_extract []int) [][]string {

	length := len(columns_to_extract)
	breadth := len(original_slice)

	extracted_slice := Make_dynamic_2d_string_slice(length, breadth)

	for row_index, original_slice_row := range original_slice {

		for column_index, original_slice_cell := range original_slice_row {

			for extracted_column_index, column_number := range columns_to_extract {

				if column_index == column_number {

					extracted_slice[row_index][extracted_column_index] = original_slice_cell
				}
			}
		}
	}

	return extracted_slice
}

func Make_dynamic_2d_string_slice(cols int, rows int) [][]string {

	var mat = make([][]string, rows)

	for i := range mat {

		mat[i] = make([]string, cols)
	}

	return mat
}

func Make_dynamic_2d_interface_slice(cols int, rows int) [][]interface{} {

	var mat = make([][]interface{}, rows)

	for i := range mat {

		mat[i] = make([]interface{}, cols)
	}

	return mat
}

func Add_single_value_column_to_2d_slice(original_slice [][]string, value string) [][]string { //#TODO - change to interface

	converted_slice := Copy_2d_slices(original_slice)

	for row_index, _ := range original_slice {

		converted_slice[row_index] = append(original_slice[row_index], value)

	}
	return converted_slice
}

func Copy_2d_slices(original_slice [][]string) [][]string {

	copy_slice := make([][]string, len(original_slice))
	for i := range original_slice {
		copy_slice[i] = make([]string, len(original_slice[i]))
		copy(copy_slice[i], original_slice[i])
	}
	return copy_slice
}

func Convert_2d_interface_to_string(interface_to_convert [][]interface{}) [][]string {

	var converted_string [][]string

	converted_string = Make_dynamic_2d_string_slice(len(interface_to_convert[0]), len(interface_to_convert))

	for row_index, interface_to_convert_row := range interface_to_convert {

		for column_index, interface_to_convert_cell := range interface_to_convert_row {

			converted_string[row_index][column_index] = interface_to_convert_cell.(string)
		}

	}

	return converted_string
}

func Convert_2d_string_to_interface(string_to_convert [][]string) [][]interface{} {

	converted_interface := Make_dynamic_2d_interface_slice(len(string_to_convert[0]), len(string_to_convert))

	for row_index, string_to_convert_row := range string_to_convert {

		for column_index, string_to_convert_cell := range string_to_convert_row {
			converted_interface[row_index][column_index] = string_to_convert_cell
		}

	}

	return converted_interface
}

func Pretty_print(json_data [][]interface{}) []byte {

	pretty_printed_json, _ := json.MarshalIndent(
		json_data,
		"",
		"	") //#TODO add pretty_printer to general utilities

	return pretty_printed_json

} //TODO - Stage 1 - move to json utilities
