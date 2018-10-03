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
		"assignStr": func(val string, attr string) map[string]string {
			return map[string]string{
				"val":  val,
				"attr": attr,
			}
		},
		"assignAttr": func(val vimAttribute, attr string) map[string]string {
			return map[string]string{
				"val":  val.String(),
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

func (vcr vimColorRow) reversed() vimColorRow {
	vcr.BG, vcr.FG = vcr.FG, vcr.BG
	return vcr
}

// Go does not have enums, so we will have to do this in an icky way
type background struct {
	dark bool
}

var light = background{dark: false}
var dark  = background{dark: true}

func (bg background) String() string {
	if bg.dark {
		return "dark"
	} else {
		return "light"
	}
}

type color struct {
	TermCode string
	GuiCode  string
}

// Sets over enums should always be emulated as structs with bools –.–'
type vimAttribute struct {
	underline bool
	reverse   bool
}

var none = vimAttribute{underline: false, reverse: false}
var underline = vimAttribute{underline: true, reverse: false}
var reverse = vimAttribute{underline: false, reverse: true}

func append(str string, more string) string {
	if str == "" {
		return more
	} else {
		return str + "," + more
	}
}

func (attr vimAttribute) String() string {
	res := ""
	if attr.underline {
		res = append(res, "underline")
	}
	if attr.reverse {
		res = append(res, "reverse")
	}
	if res == "" {
		res = "NONE"
	}
	return res
}

type vimColorRow struct {
	Name string
	BG   color
	FG   color
	Attr vimAttribute
}

type vimColorGroup struct {
	Name string
	Rows []vimColorRow
}

type nofrilsTheme struct {
	Name            string
	Slug            string
	Background      background
	FadedRow        vimColorRow
	NoneRow         vimColorRow
	NormalRow       vimColorRow
	CommentRow      vimColorRow
	HeavyCommentRow vimColorRow
	HeavyStringRow  vimColorRow
	HeavyLineRow    vimColorRow
	LineNrRow       vimColorRow
	Groups          []vimColorGroup
}
