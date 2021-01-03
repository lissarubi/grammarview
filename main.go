package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
  file := os.Args[1]
  colorRed := "\033[31m"
  colorWhite := "\033[37m"

  data, err := ioutil.ReadFile(file)
  if err != nil {
    fmt.Println("File reading error", err)
    return
  }

  text := string(data)
  fmt.Println(text)

  textSplited := strings.Split(text, "\n")

  data, err = ioutil.ReadFile("dic/words.txt")
  if err != nil {
    fmt.Println("File reading error", err)
    return
  }
  dic := string(data)

  dicSplited := strings.Split(dic, "\n")

  for i := 0; i < len(textSplited); i++{
    textLine := textSplited[i]
    textLineSplited := strings.Split(textLine, " ")

    wrongWords := []string{}

    for j := 0; j < len(textLineSplited); j++{


      currentLine := sanitazeLine(textLineSplited[j])
      checked := false

      for _, v := range dicSplited {
        if v == currentLine {
          checked = true
        }
      }

     if (checked == false){
       wrongWords = append(wrongWords, currentLine)
     }
   }
    for c := 0; c < len(wrongWords); c++{
      textLine = strings.ReplaceAll(textLine, wrongWords[c], string(colorRed) + wrongWords[c] + string(colorWhite))
    }
    fmt.Println(textLine)
  }
}

func sanitazeLine(line string) string{
      line = strings.ReplaceAll(line , "!", "")
      line = strings.ReplaceAll(line , "#", "")
      line = strings.ReplaceAll(line , "#", "")
      line = strings.ReplaceAll(line , "$", "")
      line = strings.ReplaceAll(line , "$", "")
      line = strings.ReplaceAll(line , "%", "")
      line = strings.ReplaceAll(line , "&", "")
      line = strings.ReplaceAll(line , "'", "")
      line = strings.ReplaceAll(line , "(", "")
      line = strings.ReplaceAll(line , ")", "")
      line = strings.ReplaceAll(line , "*", "")
      line = strings.ReplaceAll(line , "+", "")
      line = strings.ReplaceAll(line , ",", "")
      line = strings.ReplaceAll(line , "-", "")
      line = strings.ReplaceAll(line , ".", "")
      line = strings.ReplaceAll(line , "0", "")
      line = strings.ReplaceAll(line , "1", "")
      line = strings.ReplaceAll(line , "2", "")
      line = strings.ReplaceAll(line , "3", "")
      line = strings.ReplaceAll(line , "4", "")
      line = strings.ReplaceAll(line , "5", "")
      line = strings.ReplaceAll(line , "6", "")
      line = strings.ReplaceAll(line , "7", "")
      line = strings.ReplaceAll(line , "8", "")
      line = strings.ReplaceAll(line , "9", "")
      line = strings.ReplaceAll(line , ":", "")
      line = strings.ReplaceAll(line , "<", "")
      line = strings.ReplaceAll(line , "=", "")
      line = strings.ReplaceAll(line , ">", "")
      line = strings.ReplaceAll(line , "?", "")
      line = strings.ReplaceAll(line , "@", "")
      line = strings.ReplaceAll(line , "[", "")
      line = strings.ReplaceAll(line , "\"", "")
      line = strings.ReplaceAll(line , "]", "")
      line = strings.ReplaceAll(line , "`", "")
      line = strings.ReplaceAll(line , "{", "")
      line = strings.ReplaceAll(line , "}", "")


      return strings.ToLower(line)
}
