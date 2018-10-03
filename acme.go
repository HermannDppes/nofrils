package main

func acme() nofrilsTheme {
	cursorLineRow := vimColorRow{FG: colors["NONE"], BG: colors["White"], Attr: "NONE"}
	cursorRow := vimColorRow{FG: colors["Aqua"], BG: colors["FG"], Attr: "NONE"}
	diffAddRow := vimColorRow{FG: colors["Green1"], BG: colors["NONE"], Attr: "NONE"}
	diffChangeRow := vimColorRow{FG: colors["Orange4"], BG: colors["NONE"], Attr: "NONE"}
	diffDeleteRow := vimColorRow{FG: colors["Maroon"], BG: colors["NONE"], Attr: "NONE"}
	diffTextRow := vimColorRow{FG: colors["Navy"], BG: colors["NONE"], Attr: "NONE"}
	directoryRow := vimColorRow{FG: colors["DeepPink7"], BG: colors["NONE"], Attr: "NONE"}
	errorRow := vimColorRow{FG: colors["Red1"], BG: colors["White"], Attr: "NONE"}
	fadedRow := vimColorRow{FG: colors["SandyBrown"], BG: colors["NONE"], Attr: "NONE"}
	foldedRow := vimColorRow{FG: colors["Grey35"], BG: colors["NONE"], Attr: "NONE"}
	matchParenRow := vimColorRow{FG: colors["White"], BG: colors["Navy"], Attr: "NONE"}
	noneRow := vimColorRow{FG: colors["NONE"], BG: colors["NONE"], Attr: "NONE"}
	normalRow := vimColorRow{FG: colors["Black"], BG: colors["Cornsilk"], Attr: "NONE"}
	searchRow := vimColorRow{FG: colors["White"], BG: colors["Green1"], Attr: "NONE"}
	spellRow := vimColorRow{FG: colors["Purple2"], BG: colors["NONE"], Attr: "underline"}
	statusLineRow := vimColorRow{FG: colors["Black"], BG: colors["Plum2"], Attr: "NONE"}
	vertSplitRow := vimColorRow{FG: colors["Black"], BG: colors["LightCyan2"], Attr: "NONE"}
	visualSelectRow := vimColorRow{FG: colors["FG"], BG: colors["LightGoldenrod5"], Attr: "NONE"}
	heavyCommentRow := vimColorRow{FG: colors["DeepPink8"], BG: colors["NONE"], Attr: "NONE"}
	heavyStringRow := vimColorRow{FG: colors["Grey100"], BG: colors["NONE"], Attr: "NONE"}
	heavyLineRow := vimColorRow{FG: colors["DeepPink8"], BG: colors["NONE"], Attr: "NONE"}

	pmenuRow := reverse(normalRow)
	visualNOSRow := reverse(directoryRow)
	commentRow := reverse(visualSelectRow)

	nft := nofrilsTheme{
		Name:            "No Frils Acme",
		Slug:            "nofrils-acme",
		Background:      light,

		FadedRow:        fadedRow,
		NoneRow:         noneRow,
		NormalRow:       normalRow,
		CommentRow:      commentRow,
		HeavyCommentRow: heavyCommentRow,
		HeavyStringRow:  heavyStringRow,
		HeavyLineRow:    heavyLineRow,

		Rows: []vimColorRow{
			withName("Normal", normalRow),
			withName("ColorColumn", cursorLineRow),
			withName("MatchParen", matchParenRow),
			withName("PmenuSel", statusLineRow),
			withName("Visual", visualSelectRow),
			withName("VisualNOS", visualNOSRow),
			withName("Todo", diffAddRow),

			withName("Comment", commentRow),

			withName("FoldColumn", fadedRow),
			withName("LineNr", fadedRow),
			withName("NonText", fadedRow),

			withName("Folded", foldedRow),
			withName("SignColumn", foldedRow),
			withName("SpecialKey", foldedRow),

			withName("TabLineFill", vertSplitRow),
			withName("TabLine", vertSplitRow),
			withName("StatusLineNC", vertSplitRow),
			withName("VertSplit", vertSplitRow),

			withName("TabLineSel", statusLineRow),
			withName("StatusLine", statusLineRow),
			withName("SneakLabelTarget", statusLineRow),
			withName("SneakScope", statusLineRow),

			withName("CursorColumn", cursorLineRow),
			withName("CursorLineNr", cursorLineRow),
			withName("CursorLine", cursorLineRow),

			withName("CursorIM", cursorRow),
			withName("Cursor", cursorRow),

			withName("Directory", directoryRow),
			withName("ModeMsg", directoryRow),
			withName("MoreMsg", directoryRow),
			withName("Question", directoryRow),

			withName("Error", errorRow),
			withName("ErrorMsg", errorRow),
			withName("WarningMsg", errorRow),

			withName("IncSearch", searchRow),
			withName("Search", searchRow),

			withName("PmenuSbar", pmenuRow),
			withName("Pmenu", pmenuRow),
			withName("PmenuThumb", pmenuRow),
			withName("WildMenu", pmenuRow),

			withName("DiffAdd", diffAddRow),
			withName("DiffChange", diffChangeRow),
			withName("DiffDelete", diffDeleteRow),
			withName("DiffText", diffTextRow),

			withName("SpellBad", spellRow),
			withName("SpellCap", spellRow),
			withName("SpellLocal", spellRow),
			withName("SpellRare", spellRow),

			withName("SneakLabelMask", vertSplitRow),
			withName("SneakTarget", vertSplitRow),

			withName("Menu", noneRow),
			withName("Scrollbar", noneRow),
			withName("Tooltip", noneRow),
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
	return nft
}
