package movie

func FindName(imdbID string) string {
	switch imdbID {
	case "tt4154796":
		return "Avenger: Endgame"
	case "tt825683":
		return "Black Panther"
	default:
		return "not found."
	}
}
