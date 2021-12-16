package Hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type HangManData struct {
	WordToFind     string
	ModifiedWord   []rune
	Attempts       int
	NotAgain       []rune
	DisplayInAscii bool
}

func GamingLoop() {
	var WordTF string
	var Struct HangManData
	FilesName := []string{"./Game/main/words.txt", "./Game/main/words1.txt", "./Game/main/words2.txt"}
	WordTF = FindRandomWord(FilesName, 0)
	Struct = HangManData{WordTF, ChangeWord(WordTF), 10, []rune{'0'}, false}
	fmt.Print("To change the display font, write 'ASCII'.\nIf you want to save your progress, write 'STOP'.\nGood Luck, ")
	PtStruct := &Struct
	_ = PtStruct
}

func FindRandomWord(fileName []string, n int) string { //Permet de récupérer un mot aléatoirement dans un fichier
	content, _ := os.Open(fileName[n])
	rand.Seed(time.Now().UnixNano())
	randomLine := rand.Intn(CountLines(fileName, n))
	if randomLine == 0 {
		randomLine = 1
	}
	compt := 0
	defer content.Close()
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		compt++
		if compt == randomLine {
			return scanner.Text()
		}
	}
	return "unreachable point"
}

func CountLines(FName []string, n int) int { //Compte le nombre de ligne dans le fichier
	content, _ := os.Open(FName[n])
	defer content.Close()
	scanner := bufio.NewScanner(content)
	lineNbrs := 1
	for scanner.Scan() {
		lineNbrs++
	}
	return lineNbrs
}

func ChangeWord(s string) []rune { //Créer un tableau avec le mot modifier (avec underscore)
	rand.Seed(time.Now().UnixNano())
	n := len(s)/2 - 1
	Tabs := []rune(s)
	var ChangeTab []rune
	var resetTab []rune
	for i := 0; i < len(s); i++ {
		ChangeTab = append(ChangeTab, '_')
		resetTab = append(resetTab, '_')
	}
	if n != 0 {
		for i := 0; i < n; i++ {
			index := rand.Intn(len(s))
			ChangeTab[index] = Tabs[index]
		}
	}
	valid := ValidChangeWord(s, ChangeTab, n)
	if valid == false {
		ChangeTab = resetTab
		for i := 0; i < n; i++ {
			index := rand.Intn(len(s))
			ChangeTab[index] = Tabs[index]
		}
	}
	return ChangeTab
}

func ValidChangeWord(s string, tab []rune, n int) bool { //Valide le tableau créer précédemment
	comptLetters := 0
	for i := 0; i < len(tab); i++ {
		if tab[i] != '_' {
			comptLetters++
		}
	}
	if comptLetters != n {
		return false
	} else {
		return true
	}
}

// func Interface(PtStruct *HangManData) {
// 	fmt.Println("You have", PtStruct.Attempts, "attempts.")
// 	Dico := AsciiDisctionary()
// 	if PtStruct.DisplayInAscii == true {
// 		DisplayAscii(PtStruct, Dico)
// 	} else {
// 		DisplayToFindWord(PtStruct.ModifiedWord)
// 	}
// 	fmt.Print("Choose: ")

// 	reader := bufio.NewReader(os.Stdin)
// 	line, _, _ := reader.ReadLine()
// 	TabWord := []rune(PtStruct.WordToFind)
// 	GGWP := false
// 	TiretDu8 := true

// 	} else if len(line) == 0 {
// 		fmt.Println("Nothing in input, please try again")
// 		fmt.Println("******************************")
// 		Interface(PtStruct)
// 	} else if len(line) == 1 { //Une lettre en entrée
// 		InputLetter(PtStruct, line, TabWord, GGWP, TiretDu8)
// 	} else if len(line) == len(PtStruct.WordToFind) { //Un mot en entrée
// 		InputWord(PtStruct, line, TabWord, GGWP)
// 	} else { //Nombre de charactères différents du mot à trouver (comme si erreur lors de l'entrée d'un mot)
// 		WrongWordGuess(PtStruct)
// 	}
// }

// func InputLetter(PtStruct *HangManData, line []byte, TabWord []rune, GGWP, TiretDu8 bool) {
// 	SameLetter := false //Permet de vérifier si la lettre a déjà été entrée
// 	var StrNotAgain string
// 	if PtStruct.NotAgain[0] == '0' { //La rune de 0 permet de vérifier que le tableau est "vide"; s'il est réellement vide, il y a des erreurs dans les boucles suivantes
// 		PtStruct.NotAgain[0] = rune(line[0])
// 	} else {
// 		for i := 0; i < len(PtStruct.NotAgain); i++ {
// 			if rune(line[0]) == PtStruct.NotAgain[i] {
// 				SameLetter = true
// 				break
// 			}
// 		}
// 		if SameLetter == true {
// 			StrNotAgain = string(PtStruct.NotAgain)
// 		} else {
// 			PtStruct.NotAgain = append(PtStruct.NotAgain, rune(line[0]))
// 		}
// 	}

// 	for i := 0; i < len(TabWord); i++ { //Verifie que la lettre entrée appartient au mot et qu'elle remplace un underscore dans le tableau du mot modifié
// 		if rune(line[0]) == TabWord[i] && PtStruct.ModifiedWord[i] == '_' {
// 			PtStruct.ModifiedWord[i] = rune(line[0])
// 			GGWP = true
// 		}
// 	}

// 	if GGWP == true { //Bonne lettre en entrée
// 		for i := 0; i < len(PtStruct.ModifiedWord); i++ { //Vérifie que toutes les lettres ont été trouvées
// 			if PtStruct.ModifiedWord[i] == '_' {
// 				TiretDu8 = true
// 				break
// 			} else {
// 				TiretDu8 = false
// 			}
// 		}
// 	} else if GGWP == false && SameLetter == true {
// 		fmt.Println(StrNotAgain) // Affiche toutes les lettres déjà entrées
// 	} else if GGWP == false && PtStruct.Attempts > 1 { //Essais encore disponibles
// 		PtStruct.Attempts--
// 	} else { //Aucun essai restant
// 		PtStruct.Attempts--
// 	}
// 	if TiretDu8 == false { //Plus aucune lettre à trouver
// 		PtStruct.Attempts = 0
// 	}
// }

// func InputWord(PtStruct *HangManData, line []byte, TabWord []rune, GGWP bool) {
// 	GGWP = true
// 	for i := 0; i < len(TabWord); i++ { //Compare le mot en entrée avec le mot à trouver
// 		if rune(line[i]) != TabWord[i] {
// 			GGWP = false
// 			break
// 		}
// 	}

// 	if GGWP == true { //Mot en entrée correct
// 		EndGame(PtStruct, true)
// 		PtStruct.Attempts = 0
// 	} else if GGWP == false { //Mot en entrée incorrect
// 		WrongWordGuess(PtStruct)
// 	}
// }

// func WrongWordGuess(PtStruct *HangManData) { //Différentes possibilités quand plusieurs essais sont décrémentés
// 	if PtStruct.Attempts == 2 {
// 		PtStruct.Attempts--
// 		DisplayJose(PtStruct.Attempts)
// 		EndGame(PtStruct, false)
// 		PtStruct.Attempts--
// 	} else if PtStruct.Attempts >= 2 {
// 		PtStruct.Attempts--
// 		DisplayJose(PtStruct.Attempts)
// 		fmt.Println("******************************")
// 		PtStruct.Attempts--
// 	} else {
// 		DisplayJose(PtStruct.Attempts)
// 		PtStruct.Attempts--
// 		EndGame(PtStruct, false)
// 	}
// }

// func DisplayToFindWord(t []rune) {
// 	for i := 0; i < len(t); i++ {
// 		if i == len(t)-1 {
// 			fmt.Println(string(t[i]))
// 		} else {
// 			fmt.Print(string(t[i]))
// 			fmt.Print(" ")
// 		}
// 	}
// }

// func DisplayJose(nb int) {
// 	count := 10
// 	draw, _ := os.Open("hangman.txt")
// 	defer draw.Close()
// 	reader := bufio.NewReader(draw)
// 	buf := make([]byte, 71)
// 	for {
// 		n, err := reader.Read(buf)
// 		if err == io.EOF {
// 			break
// 		}
// 		if count == nb {
// 			fmt.Print(string(buf[0:n]))
// 		}
// 		count--
// 	}
// }
