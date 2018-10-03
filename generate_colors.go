package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Helper to pass two values to the template condAssign
	// This template emits attr=val if val is non-empty
	funcMap := template.FuncMap{
		"assign": func(val string, attr string) map[string]string {
			return map[string]string{
				"val":  val,
				"attr": attr,
			}
		},
	}

	tmpl := template.New("schemes.tmpl")
	tmpl.Funcs(funcMap)
	template.Must(tmpl.ParseFiles("schemes.tmpl"))

	nft := acme()

	err := tmpl.Execute(os.Stdout, nft)
	if err != nil {
		log.Fatal(err)
	}
}

func withName(name string, vcr vimColorRow) vimColorRow {
	vcr.Name = name
	return vcr
}

func reverse(vcr vimColorRow) vimColorRow {
	vcr.BG, vcr.FG = vcr.FG, vcr.BG
	return vcr
}

type color struct {
	TermCode string
	GuiCode  string
}

type vimColorRow struct {
	Name string
	BG   color
	FG   color
	Attr string // only used for underline currently
}

type nofrilsTheme struct {
	Name            string
	Slug            string
	FadedRow        vimColorRow
	NoneRow         vimColorRow
	NormalRow       vimColorRow
	CommentRow      vimColorRow
	HeavyCommentRow vimColorRow
	HeavyStringRow  vimColorRow
	HeavyLineRow    vimColorRow
	Rows            []vimColorRow
}
