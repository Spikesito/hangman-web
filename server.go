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
	End            bool
}

func HandleHomePage(rw http.ResponseWriter, r *http.Request, d *HangWebData) {
	tmp, _ := template.ParseFiles("./statics/hangman.html")
	tmp.Execute(rw, d)
}

func main() {
	fmt.Printf("Starting server at port 8080\n")

	FilesName := []string{"./Game/main/words.txt", "./Game/main/words1.txt", "./Game/main/words2.txt"}
	WordTF := hangman.FindRandomWord(FilesName, 0)
	WordTFRune := []rune(WordTF)
	ModifWordRune := hangman.ChangeWord(WordTF)
	ModifWordStr := string(ModifWordRune)

	Struct := HangWebData{
		Attempts:      10,
		ModifWordRune: ModifWordRune,
		ModifWordStr:  ModifWordStr,
		WordTF:        WordTF,
		WordTFRune:    WordTFRune,
		NotAgainWeb:   []rune{'0'},
	}
	Pts := &Struct

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		HandleHomePage(rw, r, &Struct)
	})

	http.HandleFunc("/hangman", func(rw http.ResponseWriter, r *http.Request) {
		Pts.Input = r.FormValue("letter")
		WordOrLetter(Pts)
		RuneToStr(Pts)
		if Pts.Attempts == 0 {
			Pts.End = true
		}
		http.Redirect(rw, r, "/", http.StatusFound)
	})

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	pts.Letter = r.FormValue("letter")
	// 	tmpl.Execute(w, pts)
	// 	// pts.WordTF = "bon_o__"
	// 	// tmpl.Execute(w, pts)
	// })

	http.ListenAndServe(":8080", nil)

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
	} else if GGWP == false {
		Pts.Attempts -= 2
	}
}
