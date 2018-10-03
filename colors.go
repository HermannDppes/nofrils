package main

var colors = map[string]color{
	"Aqua":              color{"14", "#00FFFF"},
	"Aquamarine1":       color{"122", "#87FFD7"},
	"Aquamarine2":       color{"86", "#5FFFD7"},
	"Aquamarine3":       color{"79", "#5FD7AF"},
	"Black":             color{"black", "#000000"},
	"Blue1":             color{"12", "#0000FF"},
	"Blue2":             color{"21", "#0000FF"},
	"Blue3":             color{"19", "#0000AF"},
	"Blue4":             color{"20", "#0000D7"},
	"BlueViolet":        color{"57", "#5F00FF"},
	"CadetBlue1":        color{"72", "#5FAF87"},
	"CadetBlue2":        color{"73", "#5FAFAF"},
	"Chartreuse1":       color{"118", "#87FF00"},
	"Chartreuse2":       color{"112", "#87D700"},
	"Chartreuse3":       color{"82", "#5FFF00"},
	"Chartreuse4":       color{"70", "#5FAF00"},
	"Chartreuse5":       color{"76", "#5FD700"},
	"Chartreuse6":       color{"64", "#5F8700"},
	"CornflowerBlue":    color{"69", "#5F87FF"},
	"Cornsilk":          color{"230", "#FFFFD7"},
	"Cyan1":             color{"51", "#00FFFF"},
	"Cyan2":             color{"50", "#00FFD7"},
	"Cyan3":             color{"43", "#00D7AF"},
	"DarkBlue":          color{"18", "#000087"},
	"DarkCyan":          color{"36", "#00AF87"},
	"DarkGoldenrod":     color{"136", "#AF8700"},
	"DarkGreen":         color{"22", "#005F00"},
	"DarkKhaki":         color{"143", "#AFAF5F"},
	"DarkMagenta1":      color{"90", "#870087"},
	"DarkMagenta2":      color{"91", "#8700AF"},
	"DarkOliveGreen1":   color{"191", "#D7FF5F"},
	"DarkOliveGreen2":   color{"192", "#D7FF87"},
	"DarkOliveGreen3":   color{"155", "#AFFF5F"},
	"DarkOliveGreen4":   color{"107", "#87AF5F"},
	"DarkOliveGreen5":   color{"113", "#87D75F"},
	"DarkOliveGreen6":   color{"149", "#AFD75F"},
	"DarkOrange1":       color{"208", "#FF8700"},
	"DarkOrange2":       color{"130", "#AF5F00"},
	"DarkOrange3":       color{"166", "#D75F00"},
	"DarkRed1":          color{"52", "#5F0000"},
	"DarkRed2":          color{"88", "#870000"},
	"DarkSeaGreen1":     color{"108", "#87AF87"},
	"DarkSeaGreen2":     color{"158", "#AFFFD7"},
	"DarkSeaGreen3":     color{"193", "#D7FFAF"},
	"DarkSeaGreen4":     color{"151", "#AFD7AF"},
	"DarkSeaGreen5":     color{"157", "#AFFFAF"},
	"DarkSeaGreen6":     color{"115", "#87D7AF"},
	"DarkSeaGreen7":     color{"150", "#AFD787"},
	"DarkSeaGreen8":     color{"65", "#5F875F"},
	"DarkSeaGreen9":     color{"71", "#5FAF5F"},
	"DarkSlateGray1":    color{"123", "#87FFFF"},
	"DarkSlateGray2":    color{"87", "#5FFFFF"},
	"DarkSlateGray3":    color{"116", "#87D7D7"},
	"DarkTurquoise":     color{"44", "#00D7D7"},
	"DarkViolet1":       color{"128", "#AF00D7"},
	"DarkViolet2":       color{"92", "#8700D7"},
	"DeepPink1":         color{"198", "#FF0087"},
	"DeepPink2":         color{"199", "#FF00AF"},
	"DeepPink3":         color{"197", "#FF005F"},
	"DeepPink4":         color{"161", "#D7005F"},
	"DeepPink5":         color{"162", "#D70087"},
	"DeepPink6":         color{"125", "#AF005F"},
	"DeepPink7":         color{"53", "#5F005F"},
	"DeepPink8":         color{"89", "#87005F"},
	"DeepSkyBlue1":      color{"39", "#00AFFF"},
	"DeepSkyBlue2":      color{"38", "#00AFD7"},
	"DeepSkyBlue3":      color{"31", "#0087AF"},
	"DeepSkyBlue4":      color{"32", "#0087D7"},
	"DeepSkyBlue5":      color{"23", "#005F5F"},
	"DeepSkyBlue6":      color{"24", "#005F87"},
	"DeepSkyBlue7":      color{"25", "#005FAF"},
	"DodgerBlue1":       color{"33", "#0087FF"},
	"DodgerBlue2":       color{"27", "#005FFF"},
	"DodgerBlue3":       color{"26", "#005FD7"},
	"Fuchsia":           color{"13", "#FF00FF"},
	"Gold1":             color{"220", "#FFD700"},
	"Gold2":             color{"142", "#AFAF00"},
	"Gold3":             color{"178", "#D7AF00"},
	"Green1":            color{"2", "#008000"},
	"Green2":            color{"46", "#00FF00"},
	"Green3":            color{"34", "#00AF00"},
	"Green4":            color{"40", "#00D700"},
	"Green5":            color{"28", "#008700"},
	"GreenYellow":       color{"154", "#AFFF00"},
	"Grey1":             color{"8", "#808080"},
	"Grey0":             color{"16", "#000000"},
	"Grey100":           color{"231", "#FFFFFF"},
	"Grey11":            color{"234", "#1C1C1C"},
	"Grey15":            color{"235", "#262626"},
	"Grey19":            color{"236", "#303030"},
	"Grey23":            color{"237", "#3A3A3A"},
	"Grey27":            color{"238", "#444444"},
	"Grey3":             color{"232", "#080808"},
	"Grey30":            color{"239", "#4E4E4E"},
	"Grey35":            color{"240", "#585858"},
	"Grey37":            color{"59", "#5F5F5F"},
	"Grey39":            color{"241", "#626262"},
	"Grey42":            color{"242", "#6C6C6C"},
	"Grey46":            color{"243", "#767676"},
	"Grey50":            color{"244", "#808080"},
	"Grey53":            color{"102", "#878787"},
	"Grey54":            color{"245", "#8A8A8A"},
	"Grey58":            color{"246", "#949494"},
	"Grey62":            color{"247", "#9E9E9E"},
	"Grey63":            color{"139", "#AF87AF"},
	"Grey66":            color{"248", "#A8A8A8"},
	"Grey69":            color{"145", "#AFAFAF"},
	"Grey7":             color{"233", "#121212"},
	"Grey70":            color{"249", "#B2B2B2"},
	"Grey74":            color{"250", "#BCBCBC"},
	"Grey78":            color{"251", "#C6C6C6"},
	"Grey82":            color{"252", "#D0D0D0"},
	"Grey84":            color{"188", "#D7D7D7"},
	"Grey85":            color{"253", "#DADADA"},
	"Grey89":            color{"254", "#E4E4E4"},
	"Grey93":            color{"255", "#EEEEEE"},
	"Honeydew":          color{"194", "#D7FFD7"},
	"HotPink1":          color{"205", "#FF5FAF"},
	"HotPink2":          color{"206", "#FF5FD7"},
	"HotPink3":          color{"169", "#D75FAF"},
	"HotPink4":          color{"132", "#AF5F87"},
	"HotPink5":          color{"168", "#D75F87"},
	"IndianRed1":        color{"131", "#AF5F5F"},
	"IndianRed2":        color{"167", "#D75F5F"},
	"IndianRed3":        color{"203", "#FF5F5F"},
	"IndianRed4":        color{"204", "#FF5F87"},
	"Khaki1":            color{"228", "#FFFF87"},
	"Khaki2":            color{"185", "#D7D75F"},
	"LightCoral":        color{"210", "#FF8787"},
	"LightCyan2":        color{"195", "#D7FFFF"},
	"LightCyan3":        color{"152", "#AFD7D7"},
	"LightGoldenrod1":   color{"227", "#FFFF5F"},
	"LightGoldenrod2":   color{"186", "#D7D787"},
	"LightGoldenrod3":   color{"221", "#FFD75F"},
	"LightGoldenrod4":   color{"222", "#FFD787"},
	"LightGoldenrod5":   color{"179", "#D7AF5F"},
	"LightGreen1":       color{"119", "#87FF5F"},
	"LightGreen2":       color{"120", "#87FF87"},
	"LightPink1":        color{"217", "#FFAFAF"},
	"LightPink2":        color{"174", "#D78787"},
	"LightPink3":        color{"95", "#875F5F"},
	"LightSalmon1":      color{"216", "#FFAF87"},
	"LightSalmon2":      color{"137", "#AF875F"},
	"LightSalmon3":      color{"173", "#D7875F"},
	"LightSeaGreen":     color{"37", "#00AFAF"},
	"LightSkyBlue1":     color{"153", "#AFD7FF"},
	"LightSkyBlue2":     color{"109", "#87AFAF"},
	"LightSkyBlue3":     color{"110", "#87AFD7"},
	"LightSlateBlue":    color{"105", "#8787FF"},
	"LightSlateGrey":    color{"103", "#8787AF"},
	"LightSteelBlue1":   color{"147", "#AFAFFF"},
	"LightSteelBlue2":   color{"189", "#D7D7FF"},
	"LightSteelBlue3":   color{"146", "#AFAFD7"},
	"LightYellow":       color{"187", "#D7D7AF"},
	"Lime":              color{"10", "#00FF00"},
	"Magenta1":          color{"201", "#FF00FF"},
	"Magenta2":          color{"165", "#D700FF"},
	"Magenta3":          color{"200", "#FF00D7"},
	"Magenta4":          color{"127", "#AF00AF"},
	"Magenta5":          color{"163", "#D700AF"},
	"Magenta6":          color{"164", "#D700D7"},
	"Maroon":            color{"1", "#800000"},
	"MediumOrchid1":     color{"134", "#AF5FD7"},
	"MediumOrchid2":     color{"171", "#D75FFF"},
	"MediumOrchid3":     color{"207", "#FF5FFF"},
	"MediumOrchid4":     color{"133", "#AF5FAF"},
	"MediumPurple1":     color{"104", "#8787D7"},
	"MediumPurple2":     color{"141", "#AF87FF"},
	"MediumPurple3":     color{"135", "#AF5FFF"},
	"MediumPurple4":     color{"140", "#AF87D7"},
	"MediumPurple5":     color{"97", "#875FAF"},
	"MediumPurple6":     color{"98", "#875FD7"},
	"MediumPurple7":     color{"60", "#5F5F87"},
	"MediumSpringGreen": color{"49", "#00FFAF"},
	"MediumTurquoise":   color{"80", "#5FD7D7"},
	"MediumVioletRed":   color{"126", "#AF0087"},
	"MistyRose1":        color{"224", "#FFD7D7"},
	"MistyRose2":        color{"181", "#D7AFAF"},
	"NavajoWhite1":      color{"223", "#FFD7AF"},
	"NavajoWhite2":      color{"144", "#AFAF87"},
	"Navy":              color{"4", "#000080"},
	"NavyBlue":          color{"17", "#00005F"},
	"Olive":             color{"3", "#808000"},
	"Orange1":           color{"214", "#FFAF00"},
	"Orange2":           color{"172", "#D78700"},
	"Orange3":           color{"58", "#5F5F00"},
	"Orange4":           color{"94", "#875F00"},
	"OrangeRed":         color{"202", "#FF5F00"},
	"Orchid1":           color{"170", "#D75FD7"},
	"Orchid2":           color{"213", "#FF87FF"},
	"Orchid3":           color{"212", "#FF87D7"},
	"PaleGreen1":        color{"121", "#87FFAF"},
	"PaleGreen2":        color{"156", "#AFFF87"},
	"PaleGreen3":        color{"114", "#87D787"},
	"PaleGreen4":        color{"77", "#5FD75F"},
	"PaleTurquoise1":    color{"159", "#AFFFFF"},
	"PaleTurquoise2":    color{"66", "#5F8787"},
	"PaleVioletRed":     color{"211", "#FF87AF"},
	"Pink1":             color{"218", "#FFAFD7"},
	"Pink2":             color{"175", "#D787AF"},
	"Plum1":             color{"219", "#FFAFFF"},
	"Plum2":             color{"183", "#D7AFFF"},
	"Plum3":             color{"176", "#D787D7"},
	"Plum4":             color{"96", "#875F87"},
	"Purple1":           color{"129", "#AF00FF"},
	"Purple2":           color{"5", "#CD00CD"},
	"Purple3":           color{"93", "#8700FF"},
	"Purple4":           color{"56", "#5F00D7"},
	"Purple5":           color{"54", "#5F0087"},
	"Purple6":           color{"55", "#5F00AF"},
	"Red1":              color{"9", "#FF0000"},
	"Red2":              color{"196", "#FF0000"},
	"Red3":              color{"124", "#AF0000"},
	"Red4":              color{"160", "#D70000"},
	"RosyBrown":         color{"138", "#AF8787"},
	"RoyalBlue":         color{"63", "#5F5FFF"},
	"Salmon":            color{"209", "#FF875F"},
	"SandyBrown":        color{"215", "#FFAF5F"},
	"SeaGreen1":         color{"84", "#5FFF87"},
	"SeaGreen2":         color{"85", "#5FFFAF"},
	"SeaGreen3":         color{"83", "#5FFF5F"},
	"SeaGreen4":         color{"78", "#5FD787"},
	"Silver":            color{"7", "#C0C0C0"},
	"SkyBlue1":          color{"117", "#87D7FF"},
	"SkyBlue2":          color{"111", "#87AFFF"},
	"SkyBlue3":          color{"74", "#5FAFD7"},
	"SlateBlue1":        color{"99", "#875FFF"},
	"SlateBlue2":        color{"61", "#5F5FAF"},
	"SlateBlue3":        color{"62", "#5F5FD7"},
	"SpringGreen1":      color{"48", "#00FF87"},
	"SpringGreen2":      color{"42", "#00D787"},
	"SpringGreen3":      color{"47", "#00FF5F"},
	"SpringGreen4":      color{"35", "#00AF5F"},
	"SpringGreen5":      color{"41", "#00D75F"},
	"SpringGreen6":      color{"29", "#00875F"},
	"SteelBlue1":        color{"67", "#5F87AF"},
	"SteelBlue2":        color{"75", "#5FAFFF"},
	"SteelBlue3":        color{"81", "#5FD7FF"},
	"SteelBlue4":        color{"68", "#5F87D7"},
	"Tan":               color{"180", "#D7AF87"},
	"Teal":              color{"6", "#008080"},
	"Thistle1":          color{"225", "#FFD7FF"},
	"Thistle2":          color{"182", "#D7AFD7"},
	"Turquoise1":        color{"45", "#00D7FF"},
	"Turquoise2":        color{"30", "#008787"},
	"Violet":            color{"177", "#D787FF"},
	"Wheat1":            color{"229", "#FFFFAF"},
	"Wheat2":            color{"101", "#87875F"},
	"White":             color{"white", "white"},
	"Yellow1":           color{"11", "#FFFF00"},
	"Yellow2":           color{"226", "#FFFF00"},
	"Yellow3":           color{"190", "#D7FF00"},
	"Yellow4":           color{"148", "#AFD700"},
	"Yellow5":           color{"184", "#D7D700"},
	"Yellow6":           color{"100", "#878700"},
	"Yellow7":           color{"106", "#87AF00"},
	"FG":                color{"fg", "fg"},
	"BG":                color{"bg", "bg"},
	"NONE":              color{"NONE", "NONE"},
}
