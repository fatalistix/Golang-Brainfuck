package format

type SquareBracketsFixedLengthWithArrowFormatter struct {
}

func (s SquareBracketsFixedLengthWithArrowFormatter) Format(status Status, output string) string {
	switch status {
	case Read:
		return "[~READ~] ~> " + output
	case Ok:
		return "[OUTPUT] ~> " + output
	case Error:
		return "[ERROR!] ~> " + output
	default:
		return "[UNKNWN] ~> " + output
	}
}
