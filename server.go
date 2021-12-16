package main

import (
	hangman "Hangman/Game"
	"fmt"
	"net/http"
	"text/template"
)

type HangWebData struct {
	Input          string
	Attempts       int
	WordTFRune     []rune
	WordTF         string
	ModifWordRune  []rune
	ModifWordStr   string
	NotAgainWeb    []rune
	StrNotAgainWeb string
	Win            bool
	MsgEnd         string
	SetDifficulty  bool
	Difficulty     string
}

func HandleHomePage(rw http.ResponseWriter, r *http.Request, d *HangWebData) {
	tmp, _ := template.ParseFiles("hangman.html")
	tmp.Execute(rw, d)
}

func HandleEndPage(rw http.ResponseWriter, r *http.Request, d *HangWebData) {
	tmpE, _ := template.ParseFiles("endgame.html")
	tmpE.Execute(rw, d)
}

func HandleLevelPage(rw http.ResponseWriter, r *http.Request, d *HangWebData) {
	tmpL, _ := template.ParseFiles("level.html")
	tmpL.Execute(rw, d)
}

func main() {
	fmt.Printf("Starting server at port 8080\n")

	var Struct HangWebData
	Pts := &Struct

	fp := http.FileServer(http.Dir("./asset/"))
	http.Handle("/asset/", http.StripPrefix("/asset/", fp))

	fs := http.FileServer(http.Dir("./statics/"))
	http.Handle("/statics/", http.StripPrefix("/statics/", fs))

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		if Pts.SetDifficulty == false {
			http.Redirect(rw, r, "/level", http.StatusFound)
		} else {
			HandleHomePage(rw, r, &Struct)
		}
	})

	http.HandleFunc("/endgame", func(rw http.ResponseWriter, r *http.Request) {
		HandleEndPage(rw, r, &Struct)
	})

	http.HandleFunc("/level", func(rw http.ResponseWriter, r *http.Request) {
		input := r.FormValue("level")
		if input == "" {
			HandleLevelPage(rw, r, &Struct)
		} else {
			Pts.Difficulty = r.FormValue("level")
			Pts.SetDifficulty = true
			InitializeStruct(Pts)
			http.Redirect(rw, r, "/", http.StatusFound)
		}
	})

	http.HandleFunc("/hangman", func(rw http.ResponseWriter, r *http.Request) {
		Pts.Input = r.FormValue("letter")
		WordOrLetter(Pts)
		RuneToStr(Pts)
		if Pts.Attempts > 0 {
			http.Redirect(rw, r, "/", http.StatusFound)
		} else {
			if TiretDu8Left(Pts) == false {
				Pts.Win = true
			} else {
				Pts.Win = false
			}
			Pts.SetDifficulty = false
			End(Pts)
			http.Redirect(rw, r, "/endgame", http.StatusFound)
		}
	})

	http.ListenAndServe(":8080", nil)
}

func InitializeStruct(Pts *HangWebData) {
	Pts.Attempts = 10
	FilesName := []string{"Game/main/words.txt", "Game/main/words1.txt", "Game/main/words2.txt"}
	Pts.WordTF = hangman.FindRandomWord(FilesName, ChooseFile(Pts))
	Pts.WordTFRune = []rune(Pts.WordTF)
	Pts.ModifWordRune = hangman.ChangeWord(Pts.WordTF)
	Pts.ModifWordStr = string(Pts.ModifWordRune)
	Pts.NotAgainWeb = []rune{'0'}
}

func ChooseFile(Pts *HangWebData) int {
	if Pts.Difficulty == "0" {
		return 0
	} else if Pts.Difficulty == "1" {
		return 1
	} else if Pts.Difficulty == "2" {
		return 2
	}
	return 0
}

func RuneToStr(Pts *HangWebData) {
	Pts.ModifWordStr = string(Pts.ModifWordRune)
	Pts.StrNotAgainWeb = string(Pts.NotAgainWeb)
}

func WordOrLetter(Pts *HangWebData) {
	GGWP := false
	TiretDu8 := true
	TabInput := []rune(Pts.Input)
	if len(Pts.Input) == 1 {
		InputLetter(Pts, TabInput, GGWP, TiretDu8)
	} else {
		InputWord(Pts, TabInput, GGWP, TiretDu8)
	}
}

func InputLetter(Pts *HangWebData, TabInput []rune, GGWP, TiretDu8 bool) {
	SameLetter := MemoriseLetter(Pts, TabInput)
	for i := 0; i < len(Pts.WordTFRune); i++ {
		if TabInput[0] == Pts.WordTFRune[i] && Pts.ModifWordRune[i] == '_' {
			Pts.ModifWordRune[i] = TabInput[0]
			GGWP = true
		}
	}

	if GGWP == true {
		for i := 0; i < len(Pts.ModifWordRune); i++ {
			if Pts.ModifWordRune[i] == '_' {
				TiretDu8 = true
				break
			} else {
				TiretDu8 = false
			}
		}
	} else if GGWP == false && SameLetter == true {
		fmt.Println(Pts.StrNotAgainWeb)
	} else if GGWP == false && Pts.Attempts > 1 {
		Pts.Attempts--
	} else {
		Pts.Attempts--
	}
	if TiretDu8 == false {
		Pts.Attempts = 0
		End(Pts)
	}
}

func MemoriseLetter(Pts *HangWebData, input []rune) bool {
	SameLetter := false
	if Pts.NotAgainWeb[0] == '0' {
		Pts.NotAgainWeb[0] = input[0]
	} else {
		for i := 0; i < len(Pts.NotAgainWeb); i++ {
			if input[0] == Pts.NotAgainWeb[i] {
				SameLetter = true
				break
			}
		}
		if SameLetter == true {
			Pts.StrNotAgainWeb = string(Pts.NotAgainWeb)
		} else {
			Pts.NotAgainWeb = append(Pts.NotAgainWeb, input[0])
		}
	}
	return SameLetter
}

func InputWord(Pts *HangWebData, TabInput []rune, GGWP, TiretDu8 bool) {
	GGWP = true
	for i := 0; i < len(Pts.WordTFRune); i++ {
		if TabInput[i] != Pts.WordTFRune[i] {
			GGWP = false
			break
		}
	}
	if GGWP == true {
		Pts.Attempts = 0
		Pts.ModifWordRune = Pts.WordTFRune
		Pts.Win = true
	} else if GGWP == false {
		Pts.Attempts -= 2
	}
}

func End(Pts *HangWebData) {
	if Pts.Win == true {
		Pts.MsgEnd = "GOOD JOB, YOU WON"
	} else if Pts.Win == false {
		Pts.MsgEnd = "YOU LOSE !!!"
	}
}

func TiretDu8Left(Pts *HangWebData) bool {
	for i := 0; i < len(Pts.ModifWordRune); i++ {
		if Pts.ModifWordRune[i] == '_' {
			return true
		}
	}
	return false
}
