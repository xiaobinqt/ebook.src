package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	mlib "dev/book/go_yuyanbianchen/music/lib"
)

var lib *mlib.MusicManager
var id int = 1
var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist,
				e.Source, e.Type)
		}
	case "add":
		if len(tokens) == 6 {
			id++
			lib.Add(&mlib.MusicEntry{
				ID:     strconv.Itoa(id),
				Name:   tokens[2],
				Artist: tokens[3],
				Source: tokens[4],
				Type:   tokens[5],
			})
		} else {
			fmt.Println("USAGE: lib add <name><artist><source><type>")
		}
	case "remove":
		if len(tokens) == 3 {
			lib.RemoveByName(tokens[2])
		} else {
			fmt.Println("USAGE: lib remove <name>")
		}
	default:
		fmt.Println("未知命令 ", tokens[1])
	}

}

func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}

	e := lib.Find(tokens[1])
	if e != nil {
		fmt.Println("音乐没找到 ", tokens[1])
		return
	}

	mlib.Play(e.Source, e.Type)
}

// 主函数...
func main() {
	fmt.Println(`
  Enter following commands to control the player:
  lib list -- View the existing music lib
  lib add <name><artist><source><type> -- Add a music to the music lib
  lib remove <name> -- Remove the specified music from the lib
  play <name> -- Play the specified music
  
`)

	lib = mlib.NewMusicManager()
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command -> ")

		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)

		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")

		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("未知的命令 ", tokens[0])
		}
	}
}
