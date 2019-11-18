package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	NUMBER   = 0
	TEXT     = 1
	PRINT    = 3
	ADD      = 4
	SUBTRACT = 5
)

type LEX_ITEM struct {
	ITEM_TYPE  int
	ITEM_VALUE string
	ITEM_NAME  string
}

func LEX_TEXT(text string) []LEX_ITEM {
	var instructions []LEX_ITEM
	lines := strings.Split(strings.ReplaceAll(strings.ReplaceAll(text, "\n", ""), "\t", ""), ";")

	for _, line := range lines {
		split_line := strings.Split(line, ",")
		switch split_line[0] {
		case "<comment>":
			continue
		case "<number>":
			numberItem := LEX_ITEM{NUMBER, split_line[2], split_line[1]}
			instructions = append(instructions, numberItem)
		case "<text>":
			textItem := LEX_ITEM{TEXT, split_line[2], split_line[1]}
			instructions = append(instructions, textItem)
		case "<print>":
			printItem := LEX_ITEM{PRINT, split_line[1], ""}
			instructions = append(instructions, printItem)
		case "<add>":
			addItem := LEX_ITEM{ADD, split_line[2], split_line[1]}
			instructions = append(instructions, addItem)
		case "<subtract>":
			subtractItem := LEX_ITEM{SUBTRACT, split_line[2], split_line[1]}
			instructions = append(instructions, subtractItem)
		case "":
			continue
		default:
			fmt.Println("Failed to parse \"" + line + "\".")
			os.Exit(1)
		}
	}
	return instructions
}

func PARSE_INSTRUCTIONS(instructions []LEX_ITEM) {
	var variables []LEX_ITEM
	for _, instruction := range instructions {
		if instruction.ITEM_TYPE == NUMBER || instruction.ITEM_TYPE == TEXT {
			variables = append(variables, instruction)
		} else if instruction.ITEM_TYPE == PRINT {
			value_slice := []rune(instruction.ITEM_VALUE)

			if value_slice[0] == '"' && value_slice[len(value_slice)-1] == '"' {
				fmt.Println(strings.ReplaceAll(instruction.ITEM_VALUE, "\"", ""))
			} else {
				for _, variable := range variables {
					if variable.ITEM_NAME == instruction.ITEM_VALUE {
						fmt.Println(strings.ReplaceAll(variable.ITEM_VALUE, "\"", ""))
						break
					}
				}
			}
		} else if instruction.ITEM_TYPE == ADD {
			value, err := strconv.ParseInt(instruction.ITEM_VALUE, 10, 64)
			if err != nil {
				// If it's a variable
				var variable_value string
				for _, vari := range variables {
					if vari.ITEM_NAME == instruction.ITEM_VALUE {
						variable_value = vari.ITEM_VALUE
					}
				}
				value, _ = strconv.ParseInt(variable_value, 10, 64)
				for iteration, vari := range variables {
					if vari.ITEM_NAME == instruction.ITEM_NAME {
						original_value, _ := strconv.ParseInt(vari.ITEM_VALUE, 10, 64)
						new_value := original_value + value
						vari.ITEM_VALUE = strconv.FormatInt(new_value, 10)
						variables[iteration] = vari
					}
				}

			} else {
				// If it's a number
				value, _ = strconv.ParseInt(instruction.ITEM_VALUE, 10, 64)
				for iteration, vari := range variables {
					if vari.ITEM_NAME == instruction.ITEM_NAME {
						original_value, _ := strconv.ParseInt(vari.ITEM_VALUE, 10, 64)
						new_value := original_value + value
						vari.ITEM_VALUE = strconv.FormatInt(new_value, 10)
						variables[iteration] = vari
					}
				}
			}
		} else if instruction.ITEM_TYPE == SUBTRACT {
			value, err := strconv.ParseInt(instruction.ITEM_VALUE, 10, 64)
			if err != nil {
				// If it's a variable
				var variable_value string
				for _, vari := range variables {
					if vari.ITEM_NAME == instruction.ITEM_VALUE {
						variable_value = vari.ITEM_VALUE
					}
				}
				value, _ = strconv.ParseInt(variable_value, 10, 64)
				for iteration, vari := range variables {
					if vari.ITEM_NAME == instruction.ITEM_NAME {
						original_value, _ := strconv.ParseInt(vari.ITEM_VALUE, 10, 64)
						new_value := original_value - value
						vari.ITEM_VALUE = strconv.FormatInt(new_value, 10)
						variables[iteration] = vari
					}
				}

			} else {
				// If it's a number
				value, _ = strconv.ParseInt(instruction.ITEM_VALUE, 10, 64)
				for iteration, vari := range variables {
					if vari.ITEM_NAME == instruction.ITEM_NAME {
						original_value, _ := strconv.ParseInt(vari.ITEM_VALUE, 10, 64)
						new_value := original_value - value
						vari.ITEM_VALUE = strconv.FormatInt(new_value, 10)
						variables[iteration] = vari
					}
				}
			}
		}
	}
}
