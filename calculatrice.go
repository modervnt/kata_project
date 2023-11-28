package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isRomanNumeral(s string) int {
	romans := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII",
		"IX", "X"}
	var found = 0
	for j := 0; j < len(romans); j++ {
		if s == romans[j] {
			found = 1
			break
		}
	}
	return found
}

func isArabicalNumeral(s string) int {
	arabical := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	var found = 0
	for j := 0; j < len(arabical); j++ {
		if s == arabical[j] {
			found = 1
			break
		}
	}
	return found
}

func axuena(a int, b int, o string) int {
	var result int
	if o == "+" {
		result = a + b
	} else if o == "-" {
		result = a - b
	} else if o == "*" {
		result = a * b
	} else if o == "/" {
		if b != 0 {
			result = a / b
		} else {
			fmt.Println("Error!")
		}

	}
	return result
}

func axuena1(a int, b int, o string) int {
	var result int
	if o == "+" {
		result = a + b
	} else if o == "-" {
		if a < b {
			fmt.Println("Error!")
		} else {
			result = a - b
		}
	} else if o == "*" {
		result = a * b
	} else if o == "/" {
		result = a / b
	}
	return result
}

func PourLesRomains(terme1 string, terme2 string, operator string) string {

	ChiffreRomain := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	// Recherche inverse pour obtenir les valeurs associées aux représentations romaines complètes
	var a, b int
	for roman, value := range ChiffreRomain {
		if roman == terme1 {
			a = value
		}
		if roman == terme2 {
			b = value
		}
	}

	inRoman := map[int]string{
		1:   "I",
		2:   "II",
		3:   "III",
		4:   "IV",
		5:   "V",
		6:   "VI",
		7:   "VII",
		8:   "VIII",
		9:   "IX",
		10:  "X",
		11:  "XI",
		12:  "XII",
		13:  "XIII",
		14:  "XIV",
		15:  "XV",
		16:  "XVI",
		17:  "XVII",
		18:  "XVIII",
		19:  "XIX",
		20:  "XX",
		21:  "XXI",
		22:  "XXII",
		23:  "XXIII",
		24:  "XXIV",
		25:  "XXV",
		26:  "XXVI",
		27:  "XXVII",
		28:  "XXVIII",
		29:  "XXIX",
		30:  "XXX",
		31:  "XXXI",
		32:  "XXXII",
		33:  "XXXIII",
		34:  "XXXIV",
		35:  "XXXV",
		36:  "XXXVI",
		37:  "XXXVII",
		38:  "XXXVIII",
		39:  "XXXIX",
		40:  "XL",
		41:  "XLI",
		42:  "XLII",
		43:  "XLIII",
		44:  "XLIV",
		45:  "XLV",
		46:  "XLVI",
		47:  "XLVII",
		48:  "XLVIII",
		49:  "XLIX",
		50:  "L",
		51:  "LI",
		52:  "LII",
		53:  "LIII",
		54:  "LIV",
		55:  "LV",
		56:  "LVI",
		57:  "LVII",
		58:  "LVIII",
		59:  "LIX",
		60:  "LX",
		61:  "LXI",
		62:  "LXII",
		63:  "LXIII",
		64:  "LXIV",
		65:  "LXV",
		66:  "LXVI",
		67:  "LXVII",
		68:  "LXVIII",
		69:  "LXIX",
		70:  "LXX",
		71:  "LXXI",
		72:  "LXXII",
		73:  "LXXIII",
		74:  "LXXIV",
		75:  "LXXV",
		76:  "LXXVI",
		77:  "LXXVII",
		78:  "LXXVIII",
		79:  "LXXIX",
		80:  "LXXX",
		81:  "LXXXI",
		82:  "LXXXII",
		83:  "LXXXIII",
		84:  "LXXXIV",
		85:  "LXXXV",
		86:  "LXXXVI",
		87:  "LXXXVII",
		88:  "LXXXVIII",
		89:  "LXXXIX",
		90:  "XC",
		91:  "XCI",
		92:  "XCII",
		93:  "XCIII",
		94:  "XCIV",
		95:  "XCV",
		96:  "XCVI",
		97:  "XCVII",
		98:  "XCVIII",
		99:  "XCIX",
		100: "C",
	}

	return (inRoman[axuena1(a, b, operator)])

}

func PourLesArabes(terme1 string, terme2 string, operator string) int {

	ChiffreArabe := map[string]int{

		"1":  1,
		"2":  2,
		"3":  3,
		"4":  4,
		"5":  5,
		"6":  6,
		"7":  7,
		"8":  8,
		"9":  9,
		"10": 10,
	}
	//if ChiffreArabe[terme1] < 11 && ChiffreArabe[terme2] < 11 && 1 <= ChiffreArabe[terme2] && 1 <= ChiffreArabe[terme1] {
	return (axuena(ChiffreArabe[terme1], ChiffreArabe[terme2], operator))
	/*else {
		fmt.Println("Error")
		return 0
	}*/
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Entrez votre operation:")
	operation, _ := reader.ReadString('\n')

	// Разделение строки на числа и оператор
	var parts = strings.Split(operation, " ")
	if len(parts) != 3 {
		fmt.Println("Ошибка: Неправильный формат выражения.")
		return
	}

	var position = make([]int, 2)
	var k = 0
	for i, _ := range operation {
		if operation[i] == ' ' {
			position[k] = i
			k = k + 1
		}
	}

	operator := parts[1]
	terme1 := parts[0]
	//terme2 := parts[2]
	var t2 = operation[position[1]:]
	terme2 := strings.TrimSpace(t2)

	//exclusion des caractères imbéciles

	if (terme1 != "1" && terme1 != "2" && terme1 != "3" && terme1 != "4" && terme1 != "5" && terme1 != "6" && terme1 != "7" && terme1 != "8" && terme1 != "9" && terme1 != "10" && terme1 != "I" && terme1 != "II" && terme1 != "III" && terme1 != "IV" && terme1 != "V" && terme1 != "VI" && terme1 != "VII" && terme1 != "VIII" && terme1 != "IX" && terme1 != "X") || (terme2 != "1" && terme2 != "2" && terme2 != "3" && terme2 != "4" && terme2 != "5" && terme2 != "6" && terme2 != "7" && terme2 != "8" && terme2 != "9" && terme2 != "10" && terme2 != "I" && terme2 != "II" && terme2 != "III" && terme2 != "IV" && terme2 != "V" && terme2 != "VI" && terme2 != "VII" && terme2 != "VIII" && terme2 != "IX" && terme2 != "X") {
		fmt.Println("Error!")
	} else {
		if isArabicalNumeral(terme1)*isArabicalNumeral(terme2) == 1 {
			if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
				println("Error!")
			} else {
				println(PourLesArabes(terme1, terme2, operator))
			}
		} else if isRomanNumeral(terme1)*isRomanNumeral(terme2) == 1 {
			if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
				println("Error!")
			} else {
				println(PourLesRomains(terme1, terme2, operator))
			}
		} else if isArabicalNumeral(terme1)*isRomanNumeral(terme2) == 1 || isArabicalNumeral(terme2)*isRomanNumeral(terme1) == 1 {
			println("Error!")
		}
	}
}
