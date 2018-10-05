package main

func acme() nofrilsTheme {
	colorRow        := colRow("NONE",          "Grey85")
	cursorLineRow   := colRow("NONE",          "Khaki1")
	cursorLineNrRow := colRow("TermBlack",     "Khaki1")
	cursorRow       := colRow("FG",            "Navy")
	diffAddRow      := colRow("Green1",        "NONE")
	diffChangeRow   := colRow("Orange4",       "NONE")
	diffDeleteRow   := colRow("Maroon",        "NONE")
	diffTextRow     := colRow("Navy",          "NONE")
	directoryRow    := colRow("DeepPink7",     "NONE")
	errorRow        := colRow("Red1",          "White")
	fadedRow        := colRow("DarkGoldenrod", "NONE")
	foldedRow       := colRow("Grey35",        "NONE")
	specialKeyRow   := colRow("Grey35",        "BG")
	matchParenRow   := colRow("White",         "Navy")
	noneRow         := colRow("NONE",          "NONE")
	normalRow       := colRow("Black",         "Cornsilk")
	searchRow       := colRow("White",         "Green1")
	spellRow        := colRow("Purple2",       "NONE").underline()
	statusLineRow   := colRow("FG",            "Plum2")
	pmenuSelRow     := colRow("FG",            "Fuchsia")
	sneak2Row       := colRow("Black",         "Plum2")
	tabLineFillRow  := colRow("FG",            "DarkGoldenrod")
	vertSplitRow    := colRow("FG",            "LightCyan2")
	sneak1Row       := colRow("Black",         "LightCyan2")
	heavyCommentRow := colRow("DeepPink8",     "NONE")
	heavyStringRow  := colRow("NONE",          "White")
	heavyLineRow    := colRow("DeepPink8",     "NONE")
	reversedRow     := colRow("NONE",          "NONE").reverse()
	lineNrRow       := colRow("DarkGoldenrod", "TermBG")
	visualRow       := colRow("TermFG",        "LightGoldenrod4")
	visualNOSRow    := visualRow.underline()
	wildMenuRow     := colRow("Plum2",         "FG")

	baseline := vimColorGroup{
		Name: "Baseline",
		Rows: []vimColorRow{
			withName("Normal", normalRow),
		},
	}

	faded := vimColorGroup{
		Name: "Faded",
		Rows: []vimColorRow{
			withName("ColorColumn", colorRow),
			withName("Comment", fadedRow),
			withName("FoldColumn", fadedRow),
			withName("Folded", foldedRow),
			withName("LineNr", lineNrRow),
			withName("NonText", fadedRow),
			withName("SignColumn", specialKeyRow),
			withName("SpecialKey", specialKeyRow),
			withName("StatusLineNC", vertSplitRow),
			withName("VertSplit", vertSplitRow),
		},
	}

	highlight := vimColorGroup{
		Name: "Highlight",
		Rows: []vimColorRow{
			withName("CursorColumn", cursorLineRow),
			withName("CursorIM", cursorRow),
			withName("CursorLineNr", cursorLineNrRow),
			withName("CursorLine", cursorLineRow),
			withName("Cursor", cursorRow),
			withName("Directory", directoryRow),
			withName("ErrorMsg", errorRow),
			withName("Error", errorRow),
			withName("IncSearch", searchRow),
			withName("MatchParen", matchParenRow),
			withName("ModeMsg", directoryRow),
			withName("MoreMsg", directoryRow),
			withName("PmenuSel", pmenuSelRow),
			withName("Question", directoryRow),
			withName("Search", searchRow),
			withName("StatusLine", statusLineRow),
			withName("Todo", diffAddRow),
			withName("WarningMsg", errorRow),
			withName("WildMenu", wildMenuRow),
			withName("Visual", visualRow),
			withName("VisualNOS", visualNOSRow),
		},
	}

	reversed := vimColorGroup{
		Name: "Reversed",
		Rows: []vimColorRow{
			withName("PmenuSbar", reversedRow),
			withName("Pmenu", reversedRow),
			withName("PmenuThumb", reversedRow),
			withName("TabLineSel", reversedRow),
		},
	}

	diff := vimColorGroup{
		Name: "Diff",
		Rows: []vimColorRow{
			withName("DiffAdd", diffAddRow),
			withName("DiffChange", diffChangeRow),
			withName("DiffDelete", diffDeleteRow),
			withName("DiffText", diffTextRow),
		},
	}

	spell := vimColorGroup{
		Name: "Spell",
		Rows: []vimColorRow{
			withName("SpellBad", spellRow),
			withName("SpellCap", spellRow),
			withName("SpellLocal", spellRow),
			withName("SpellRare", spellRow),
		},
	}

	vimFeatures := vimColorGroup{
		Name: "Vim Features",
		Rows: []vimColorRow{
			withName("Menu", noneRow),
			withName("Scrollbar", noneRow),
			withName("TabLineFill", tabLineFillRow),
			withName("TabLine", noneRow),
			withName("Tooltip", noneRow),
		},
	}

	syntaxHighsepiaing := vimColorGroup{
		Name: "Syntax Highsepiaing (or lack there of)",
		Rows: []vimColorRow{
			withName("Boolean", noneRow),
			withName("Character", noneRow),
			withName("Conceal", noneRow),
			withName("Conditional", noneRow),
			withName("Constant", noneRow),
			withName("Debug", noneRow),
			withName("Define", noneRow),
			withName("Delimiter", noneRow),
			withName("Directive", noneRow),
			withName("Exception", noneRow),
			withName("Float", noneRow),
			withName("Format", noneRow),
			withName("Function", noneRow),
			withName("Identifier", noneRow),
			withName("Ignore", noneRow),
			withName("Include", noneRow),
			withName("Keyword", noneRow),
			withName("Label", noneRow),
			withName("Macro", noneRow),
			withName("Number", noneRow),
			withName("Operator", noneRow),
			withName("PreCondit", noneRow),
			withName("PreProc", noneRow),
			withName("Repeat", noneRow),
			withName("SpecialChar", noneRow),
			withName("SpecialComment", noneRow),
			withName("Special", noneRow),
			withName("Statement", noneRow),
			withName("StorageClass", noneRow),
			withName("String", noneRow),
			withName("Structure", noneRow),
			withName("Tag", noneRow),
			withName("Title", noneRow),
			withName("Typedef", noneRow),
			withName("Type", noneRow),
			withName("Underlined", noneRow),
		},
	}

	sneak := vimColorGroup{
		Name: "Sneak",
		Rows: []vimColorRow{
			withName("SneakLabelMask", sneak1Row),
			withName("SneakTarget", sneak1Row),
			withName("SneakLabelTarget", sneak2Row),
			withName("SneakScope", sneak2Row),
		},
	}

	nft := nofrilsTheme{
		Name:            "No Frils Acme",
		Slug:            "nofrils-acme",
		Background:      light,

		FadedRow:        fadedRow,
		NoneRow:         noneRow,
		NormalRow:       normalRow,
		HeavyCommentRow: heavyCommentRow,
		HeavyStringRow:  heavyStringRow,
		HeavyLineRow:    heavyLineRow,
		LineNrRow:       lineNrRow,

		Groups: []vimColorGroup{
			baseline,
			faded,
			highlight,
			reversed,
			diff,
			spell,
			vimFeatures,
			syntaxHighsepiaing,
			sneak,
		},
	}
	return nft
}
