package main

func col_row(fg string, bg string) vimColorRow {
	return vimColorRow{FG: colors[fg], BG: colors[bg], Attr: none}
}

func acme() nofrilsTheme {
	cursorLineRow   := col_row("NONE",          "White")
	cursorRow       := col_row("Aqua",          "FG")
	diffAddRow      := col_row("Green1",        "NONE")
	diffChangeRow   := col_row("Orange4",       "NONE")
	diffDeleteRow   := col_row("Maroon",        "NONE")
	diffTextRow     := col_row("Navy",          "NONE")
	directoryRow    := col_row("DeepPink7",     "NONE")
	errorRow        := col_row("Red1",          "White")
	fadedRow        := col_row("DarkGoldenrod", "NONE")
	foldedRow       := col_row("Grey35",        "NONE")
	matchParenRow   := col_row("White",         "Navy")
	noneRow         := col_row("NONE",          "NONE")
	normalRow       := col_row("Black",         "Cornsilk")
	searchRow       := col_row("White",         "Green1")
	spellRow        := col_row("Purple2",       "NONE").underline()
	statusLineRow   := col_row("Black",         "Plum2")
	vertSplitRow    := col_row("Black",         "LightCyan2")
	visualSelectRow := col_row("FG",            "LightGoldenrod5")
	heavyCommentRow := col_row("DeepPink8",     "NONE")
	heavyStringRow  := col_row("NONE",          "White")
	heavyLineRow    := col_row("DeepPink8",     "NONE")
	reversedRow     := col_row("NONE",          "NONE").reverse()
	lineNrRow       := col_row("DarkGoldenrod", "TermBG")

	visualNOSRow := directoryRow.reversed()

	baseline := vimColorGroup{
		Name: "Baseline",
		Rows: []vimColorRow{
			withName("Normal", normalRow),
		},
	}

	faded := vimColorGroup{
		Name: "Faded",
		Rows: []vimColorRow{
			withName("ColorColumn", cursorLineRow),
			withName("Comment", fadedRow),
			withName("FoldColumn", fadedRow),
			withName("Folded", foldedRow),
			withName("LineNr", fadedRow),
			withName("NonText", fadedRow),
			withName("SignColumn", foldedRow),
			withName("SpecialKey", foldedRow),
			withName("StatusLineNC", vertSplitRow),
			withName("VertSplit", vertSplitRow),
		},
	}

	highlight := vimColorGroup{
		Name: "Highlight",
		Rows: []vimColorRow{
			withName("CursorColumn", cursorLineRow),
			withName("CursorIM", cursorRow),
			withName("CursorLineNr", cursorLineRow),
			withName("CursorLine", cursorLineRow),
			withName("Cursor", cursorRow),
			withName("Directory", directoryRow),
			withName("ErrorMsg", errorRow),
			withName("Error", errorRow),
			withName("IncSearch", searchRow),
			withName("MatchParen", matchParenRow),
			withName("ModeMsg", directoryRow),
			withName("MoreMsg", directoryRow),
			withName("PmenuSel", statusLineRow),
			withName("Question", directoryRow),
			withName("Search", searchRow),
			withName("StatusLine", statusLineRow),
			withName("Todo", diffAddRow),
			withName("WarningMsg", errorRow),
			withName("WildMenu", statusLineRow.reversed()),
			withName("Visual", visualSelectRow),
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

	vim_features := vimColorGroup{
		Name: "Vim Features",
		Rows: []vimColorRow{
			withName("Menu", noneRow),
			withName("Scrollbar", noneRow),
			withName("TabLineFill", vertSplitRow),
			withName("TabLine", vertSplitRow),
			withName("Tooltip", noneRow),
		},
	}

	syntax_highsepiaing := vimColorGroup{
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
			withName("SneakLabelMask", vertSplitRow),
			withName("SneakTarget", vertSplitRow),
			withName("SneakLabelTarget", statusLineRow),
			withName("SneakScope", statusLineRow),
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
			vim_features,
			syntax_highsepiaing,
			sneak,
		},
	}
	return nft
}
