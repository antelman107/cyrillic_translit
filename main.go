package cyrillic_translit

var StandardDict = map[string]string{
	"а": "a", "б": "b", "в": "v", "г": "g", "д": "d", "е": "e",
	"ё": "yo", "ж": "zh", "з": "z", "и": "i", "й": "y", "к": "k",
	"л": "l", "м": "m", "н": "n", "о": "o", "п": "p", "р": "r",
	"с": "s", "т": "t", "у": "u", "ф": "f", "х": "kh", "ц": "ts",
	"ч": "ch", "ш": "sh", "щ": "sch", "ъ": "", "ы": "i",
	"ь": "", "э": "e", "ю": "yu", "я": "ya",
	"А": "A", "Б": "B", "В": "V", "Г": "G", "Д": "D", "Е": "E",
	"Ё": "YO", "Ж": "ZH", "З": "Z", "И": "I", "Й": "Y", "К": "K",
	"Л": "L", "М": "M", "Н": "N", "О": "O", "П": "P", "Р": "R",
	"С": "S", "Т": "T", "У": "U", "Ф": "F", "Х": "KH", "Ц": "TS",
	"Ч": "CH", "Ш": "SH", "Щ": "SCH", "Ъ": "", "Ы": "I",
	"Ь": "", "Э": "E", "Ю": "YU", "Я": "YA",
}

func Do(input string) (output string) {
	return DoDict(input, StandardDict)
}

func DoDict(input string, dict map[string]string) (output string) {
	output = ""
	for _, v := range input {
		replacement, ok := dict[string(v)]
		if ok {
			output += replacement
		} else {
			output += string(v)
		}
	}

	return
}
