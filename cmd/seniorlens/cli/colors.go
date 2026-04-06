package cli

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

func SetWhite(s string) string {
	return WHITE + s + WHITE
}

func SetGreen(s string) string {
	return GREEN + s + WHITE
}

func SetYellow(s string) string {
	return YELLOW + s + WHITE
}

func SetRed(s string) string {
	return RED + s + WHITE
}

func SetPurple(s string) string {
	return PURPLE + s + WHITE
}

func SetBlue(s string) string {
	return BLUE + s + WHITE
}