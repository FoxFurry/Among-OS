package command

type uname struct{}

func (e *uname) addParameterData(param byte) string {
	switch param {
	case 'r':
		return "kernel release 4.2069 k-among FURRYES"
	case 'p':
		return "Intel(R) Core(TM) i9-69420K CPU @ 6.420hz"
	case 's':
		return "FOX_FURRY_OS"
	case 'a':
		return "Bruh"
	case 'i':
		return "SuS"
	case 'f':
		return "                                         x\nxxxx                                    xx\n x xx                                  xx x\n  x  x                                xx  x\n  xx  xx                             xx   x\n   xx  xx                           x      x\n    x   xxx                        x        x\n    xx    xx                      xx        xx\n     x     xx                    x           x\n     xx      xx                 xx           x\n      x       xxxxxxxxxxxxxxxxxxx             x\n     xx                                       x\n     x                                        xx\n    xx                                         xx\n   xx                                            x\n   x                                             xx\n  xx                                              x\n  x                                               x\n  x      xx                                       x\n  x      xxxx                     xxx            xx\n  x      x   x                  xx  x            x\n  x      x    xx              xx    x           xx\n  x      x      xx          xxx     x           x\n  x      xxxxxx xxx        xxxxxxxxxx          x\n  x                                           xx\n  xx                                         xx\n   xx                                      xxx\n    xx                                    xx\n     xx                                 xx\n      xx                             xxxx\n       xx                        xxxxx\n       xx                    xxxx\n        x                 xxx\n        xx              xxx\n         xx            xx\n          xx          xx\n           x xxxxx xxx\n           xxxxxxxxxx\n            xxxxxxx\n             xxx\n             xx"

	default:
		return ""
	}
}

func (e *uname) Execute(parameters []string) string {
	var result string

	alreadyUsedParams := make(map[byte]struct{})

	if len(parameters) == 0 {
		result += e.addParameterData('s')
	} else {
		for _, param := range parameters {
			if param[0] == '-' {
				for _, paramVal := range param[1:] {
					parameterByte := byte(paramVal)
					if _, ok := alreadyUsedParams[parameterByte]; !ok {
						alreadyUsedParams[parameterByte] = struct{}{}
						result += e.addParameterData(parameterByte) + " "
					}
				}
			} else {
				return "Parameters are expected to start with '-'"
			}
		}
	}

	return result
}

func (e *uname) Help() string {
	return "uname command, short for Unix Name, will print detailed information about your Linux system like the machine name, operating system, kernel, and so on\n" +
		"Parameters:\n" +
		"\t-r -- print kernel release version\n" +
		"\t-p -- print CPU information\n" +
		"\t-s -- print kernel name\n" +
		"\t-a -- print Bruh\n" +
		"\t-i -- print SuS\n" +
		"\t-f -- print full\n" +
		"If no parameters is specified - kernel name will be shown (e.g. uname -s)"
}
