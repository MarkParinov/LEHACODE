package main

import (
	"errors"
	"strconv"
	"unicode/utf8"
)

type encoding_table_unit struct {
	value    string
	encoding string
}

type leha_byte struct {
	value string
}

func removeLastChar(str string) string {
	for len(str) > 0 {
		_, size := utf8.DecodeLastRuneInString(str)
		return str[:len(str)-size]
	}
	return str
}

func removeFirstChar(str string) string {
	for len(str) > 0 {
		_, size := utf8.DecodeLastRuneInString(str)
		return str[len(str)-size:]
	}
	return str
}

func generateEncodingTable() (slice []encoding_table_unit) {
	var symbols_avaible_in_encoding string = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890 $"
	var current_byte string
	for a := 0; a <= 1; a++ {
		current_byte += strconv.Itoa(a)
		for b := 0; b <= 1; b++ {
			current_byte += strconv.Itoa(b)
			for c := 0; c <= 1; c++ {
				current_byte += strconv.Itoa(c)
				for d := 0; d <= 1; d++ {
					current_byte += strconv.Itoa(d)
					for e := 0; e <= 1; e++ {
						current_byte += strconv.Itoa(e)
						for f := 0; f <= 1; f++ {
							current_byte += strconv.Itoa(f)
							slice = append(slice, encoding_table_unit{string(symbols_avaible_in_encoding[(a*32 + b*16 + c*8 + d*4 + e*2 + f)]), current_byte})
							current_byte = removeLastChar(current_byte)
						}
						current_byte = removeLastChar(current_byte)
					}
					current_byte = removeLastChar(current_byte)
				}
				current_byte = removeLastChar(current_byte)
			}
			current_byte = removeLastChar(current_byte)
		}
		current_byte = removeLastChar(current_byte)
	}
	return slice
}

var encoding_table []encoding_table_unit = generateEncodingTable()

func prepareStringForDecoding(str string) string {
	if len(str)%2 == 0 {
	} else {
		for i := 0; i < len(str); i++ {
			if len(str)%2 == 0 {
				break
			} else {
				str += "0"
			}
		}
	}
	return str
}

func lehaByteToChar(char leha_byte) string {
	var table_byte string

	for i := 0; i < 6; i++ {
		if string(char.value[i]) == "A" {
			table_byte += "1"
		} else if string(char.value[i]) == "0" {
			table_byte += "0"
		}
	}

	for i := 0; i < len(encoding_table); i++ {
		if table_byte == encoding_table[i].encoding {
			return encoding_table[i].value
		}
	}
	return ""
}

func charToLehaByte(char string) (leha_byte, error) {
	var output leha_byte
	var byte_to_encode string
	var char_is_valid bool = false

	for i := 0; i < len(encoding_table); i++ {
		if char == encoding_table[i].value {
			byte_to_encode = encoding_table[i].encoding
			char_is_valid = true
		}
	}
	if !char_is_valid {
		return leha_byte{""}, errors.New("Character '" + char + "' is not supported by LEHACODE.")
	}
	for n := 0; n < 6; n++ {
		if string(byte_to_encode[n]) == "1" {
			output.value += "A"
		} else if string(byte_to_encode[n]) == "0" {
			output.value += "0"
		}
	}
	return output, nil
}

func encodeLehaCode(str string) []leha_byte {
	var output []leha_byte
	for i := 0; i < len(str); i++ {
		cur_el := string(str[i])
		converted_el, err := charToLehaByte(cur_el)
		if err != nil {
			output = append(output, leha_byte{"NOT_SUPPORTED"})
		} else {
			output = append(output, converted_el)
		}
	}
	return output
}

func encodeToFullString(str string) string {
	var string_output string
	var output []leha_byte
	for i := 0; i < len(str); i++ {
		cur_el := string(str[i])
		converted_el, err := charToLehaByte(cur_el)
		if err != nil {
			output = append(output, leha_byte{"NOT_SUPPORTED"})
		} else {
			output = append(output, converted_el)
		}
	}

	for i := 0; i < len(output); i++ {
		string_output += string(output[i].value)
	}
	return string_output
}

func decodeLehaCode(str string) string {
	var output string
	var leha_bytes []leha_byte
	for len(str)%6 != 0 {
		str = "0" + str
	}

	i := -1

	for len(str) != 0 {
		i++
		leha_bytes = append(leha_bytes, leha_byte{})
		for n := 0; n < 6; n++ {
			leha_bytes[i].value += string(str[0])
			str = str[1:]
		}
	}

	for i := 0; i < len(leha_bytes); i++ {
		output += lehaByteToChar(leha_bytes[i])
	}
	return output
}
