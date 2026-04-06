package cli

import (
	"fmt"
)

var (
	ClearScreen = "\033[2J\033[H"
	GREEN = "\033[32m"
	YELLOW = "\033[33m"
	RED = "\033[31m"
	PURPLE = "\033[34m"
	BLUE = "\033[36m"
	GREY = "\033[37m"
	WHITE = "\033[0m"
)

func setWhite(s string) string {
	return WHITE + s + WHITE
}

func setGreen(s string) string {
	return GREEN + s + WHITE
}

func setYellow(s string) string {
	return YELLOW + s + WHITE
}

func setRed(s string) string {
	return RED + s + WHITE
}

func setPurple(s string) string {
	return PURPLE + s + WHITE
}

func setBlue(s string) string {
	return BLUE + s + WHITE
}

func header(version string, analysisVersion string) {
	var asciiTop = YELLOW + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣀⣀⣀⣀⡀⠀⠀⠀⠀⠀⠀⠀   \n"+
				   YELLOW + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⠾⠛⢉⣉⣉⣉⡉⠛⠷⣦⣄⠀⠀⠀⠀   " + GREY + "███████╗███████╗███╗   ██╗██╗ ██████╗ ██████╗      ██╗     ███████╗███╗   ██╗███████╗\n"+
				   YELLOW + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⠋⣠⣴⣿⣿⣿⣿⣿⡿⣿⣶⣌⠹⣷⡀⠀⠀   " + GREY + "██╔════╝██╔════╝████╗  ██║██║██╔═══██╗██╔══██╗     ██║     ██╔════╝████╗  ██║██╔════╝\n"+
				   YELLOW + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⠁⣴⣿⣿⣿⣿⣿⣿⣿⣿⣆⠉⠻⣧⠘⣷⠀⠀   " + GREY + "███████╗█████╗  ██╔██╗ ██║██║██║   ██║██████╔╝     ██║     █████╗  ██╔██╗ ██║█████╗  \n"+
				   YELLOW + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣷⠀⢿⡆⠈⠛⠻⠟⠛⠉⠀⠀⠀⠀⠀⠀⣾⠃⠀   " + GREY + "╚════██║██╔══╝  ██║╚██╗██║██║██║   ██║██╔══██╗     ██║     ██╔══╝  ██║╚██╗██║╚════██╗\n"+
				   YELLOW + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⣧⡀⠻⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣼⠃⠀⠀   " + GREY + "███████║███████╗██║ ╚████║██║╚██████╔╝██║  ██║     ███████╗███████╗██║ ╚████║███████║\n"+
				   YELLOW + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢼⠿⣦⣄⠀⠀⠀⠀⠀⠀⠀⣀⣴⠟⠁⠀⠀⠀   " + GREY + "╚══════╝╚══════╝╚═╝  ╚═══╝╚═╝ ╚═════╝ ╚═╝  ╚═╝     ╚══════╝╚══════╝╚═╝  ╚═══╝╚══════╝\n"+
				   YELLOW + "⠀⠀⠀⠀⠀⠀⠀⠀⣠⣾⣿⣦⠀⠀⠈⠉⠛⠓⠲⠶⠖⠚⠋⠉⠀⠀⠀⠀⠀⠀   \n"+
				   YELLOW + "⠀⠀⠀⠀⠀⠀⣠⣾⣿⣿⠟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀   " + GREEN + "           STEP FROM JUNIOR TO SENIOR | Code Smells & Practices Detector\n"+
				   YELLOW + "     ⣾⣿⣿⠟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀                 " + setBlue("Version " + version) + " " + setWhite("|") + setPurple(" Analysis: " + analysisVersion)+ " " + setWhite("|") + setBlue(" Engine: SL-Core 2.1\n")+
				   YELLOW + "     ⠈⠛⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀"
		
	fmt.Println(ClearScreen)
	fmt.Println(asciiTop)
}
