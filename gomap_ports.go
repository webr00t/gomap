package gomap

var commonlist = map[int]string{
	21:    "ftp",
	22:    "ssh",
	23:    "telnet",
	25:    "smtp",
	53:    "dns",
	80:    "http",
	443:   "https",
	2082:  "cpanel",
	2083:  "cpanel",
	3306:  "mysql",
}

var detailedlist = map[int]string{
	// Custom Ports
	5550:  "vnc server",
	5938:  "teamviewer",
	25565: "minecraft",
	24800: "synergy",
	//
	1:     "Port Service Multiplexer",
	2:     "Management Utility",
	3:     "Compression Process",
	4:     "Unassigned",
	5:     "Remote Job Entry",
	7:     "Echo",

}
