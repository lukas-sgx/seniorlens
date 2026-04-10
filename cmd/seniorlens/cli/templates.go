package cli

import (
	"fmt"
	"github.com/lukas-sgx/seniorlens/internal/ui"
)

func header(version string, analysisVersion string) {
	var asciiTop = ui.YELLOW + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣀⣀⣀⣀⡀⠀⠀⠀⠀⠀⠀⠀   \n"+
				   ui.YELLOW + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⠾⠛⢉⣉⣉⣉⡉⠛⠷⣦⣄⠀⠀⠀⠀   " + ui.GREY + "███████╗███████╗███╗   ██╗██╗ ██████╗ ██████╗      ██╗     ███████╗███╗   ██╗███████╗\n"+
				   ui.YELLOW + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⠋⣠⣴⣿⣿⣿⣿⣿⡿⣿⣶⣌⠹⣷⡀⠀⠀   " + ui.GREY + "██╔════╝██╔════╝████╗  ██║██║██╔═══██╗██╔══██╗     ██║     ██╔════╝████╗  ██║██╔════╝\n"+
				   ui.YELLOW + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⠁⣴⣿⣿⣿⣿⣿⣿⣿⣿⣆⠉⠻⣧⠘⣷⠀⠀   " + ui.GREY + "███████╗█████╗  ██╔██╗ ██║██║██║   ██║██████╔╝     ██║     █████╗  ██╔██╗ ██║█████╗  \n"+
				   ui.YELLOW + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣷⠀⢿⡆⠈⠛⠻⠟⠛⠉⠀⠀⠀⠀⠀⠀⣾⠃⠀   " + ui.GREY + "╚════██║██╔══╝  ██║╚██╗██║██║██║   ██║██╔══██╗     ██║     ██╔══╝  ██║╚██╗██║╚════██╗\n"+
				   ui.YELLOW + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⣧⡀⠻⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣼⠃⠀⠀   " + ui.GREY + "███████║███████╗██║ ╚████║██║╚██████╔╝██║  ██║     ███████╗███████╗██║ ╚████║███████║\n"+
				   ui.YELLOW + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢼⠿⣦⣄⠀⠀⠀⠀⠀⠀⠀⣀⣴⠟⠁⠀⠀⠀   " + ui.GREY + "╚══════╝╚══════╝╚═╝  ╚═══╝╚═╝ ╚═════╝ ╚═╝  ╚═╝     ╚══════╝╚══════╝╚═╝  ╚═══╝╚══════╝\n"+
				   ui.YELLOW + "⠀⠀⠀⠀⠀⠀⠀⠀⣠⣾⣿⣦⠀⠀⠈⠉⠛⠓⠲⠶⠖⠚⠋⠉⠀⠀⠀⠀⠀⠀   \n"+
				   ui.YELLOW + "⠀⠀⠀⠀⠀⠀⣠⣾⣿⣿⠟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀   " + ui.GREEN + "           STEP FROM JUNIOR TO SENIOR | Code Smells & Practices Detector\n"+
				   ui.YELLOW + "     ⣾⣿⣿⠟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀                 " + ui.SetBlue("Version " + version) + " " + ui.SetWhite("|") + ui.SetPurple(" Analysis: " + analysisVersion)+ " " + ui.SetWhite("|") + ui.SetBlue(" Engine: SL-Core 2.1\n")+
				   ui.YELLOW + "     ⠈⠛⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀" + ui.WHITE
		
	fmt.Println(ui.ClearScreen)
	fmt.Println(asciiTop)
}
