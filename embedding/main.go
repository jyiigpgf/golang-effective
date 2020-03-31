package main

import "fmt"

type git struct {
	Name string
}

func (g *git)print()  {
	fmt.Println(g.Name)
}

type github struct {
	*git
}

func (g *github)print()  {
	fmt.Println(g.Name)
}

func main() {

	_git := &git{Name:"git"}
	g := &github{git:_git}
	g.print()
}
