package main

func acme() nofrilsTheme {
	cursorLineRow := vimColorRow{FG: colors["NONE"], BG: colors["White"], Attr: none}
	cursorRow := vimColorRow{FG: colors["Aqua"], BG: colors["FG"], Attr: none}
	diffAddRow := vimColorRow{FG: colors["Green1"], BG: colors["NONE"], Attr: none}
	diffChangeRow := vimColorRow{FG: colors["Orange4"], BG: colors["NONE"], Attr: none}
	diffDeleteRow := vimColorRow{FG: colors["Maroon"], BG: colors["NONE"], Attr: none}
	diffTextRow := vimColorRow{FG: colors["Navy"], BG: colors["NONE"], Attr: none}
	directoryRow := vimColorRow{FG: colors["DeepPink7"], BG: colors["NONE"], Attr: none}
	errorRow := vimColorRow{FG: colors["Red1"], BG: colors["White"], Attr: none}
	fadedRow := vimColorRow{FG: colors["DarkGoldenrod"], BG: colors["NONE"], Attr: none}
	foldedRow := vimColorRow{FG: colors["Grey35"], BG: colors["NONE"], Attr: none}
	matchParenRow := vimColorRow{FG: colors["White"], BG: colors["Navy"], Attr: none}
	noneRow := vimColorRow{FG: colors["NONE"], BG: colors["NONE"], Attr: none}
	normalRow := vimColorRow{FG: colors["Black"], BG: colors["Cornsilk"], Attr: none}
	searchRow := vimColorRow{FG: colors["White"], BG: colors["Green1"], Attr: none}
	spellRow := vimColorRow{FG: colors["Purple2"], BG: colors["NONE"], Attr: underline}
	statusLineRow := vimColorRow{FG: colors["Black"], BG: colors["Plum2"], Attr: none}
	vertSplitRow := vimColorRow{FG: colors["Black"], BG: colors["LightCyan2"], Attr: none}
	visualSelectRow := vimColorRow{FG: colors["FG"], BG: colors["LightGoldenrod5"], Attr: none}
	heavyCommentRow := vimColorRow{FG: colors["DeepPink8"], BG: colors["NONE"], Attr: none}
	heavyStringRow := vimColorRow{FG: colors["NONE"], BG: colors["White"], Attr: none}
	heavyLineRow := vimColorRow{FG: colors["DeepPink8"], BG: colors["NONE"], Attr: none}
	reversedRow := vimColorRow{FG: colors["NONE"], BG: colors["NONE"], Attr: reverse}
	lineNrRow := vimColorRow{FG: colors["DarkGoldenrod"], BG: colors["TermBG"], Attr: none}

	visualNOSRow := directoryRow.reversed()
	commentRow := visualSelectRow.reversed()

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
			withName("Comment", commentRow),
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
