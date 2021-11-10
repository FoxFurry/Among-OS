package command

type uname struct{}

func (e *uname) addParameterData(param byte) string {
	switch param {
	case 'r':
		return "kernel release 5.69 k-among"
	case 'p':
		return "Intel(R) Core(TM) i9-6969K CPU @ 6.969Ghz"
	case 's':
		return "Among OS"
	case 'a':
		return "Bruh"
	case 'i':
		return "SuS"
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
		"If no parameters is specified - kernel name will be shown (e.g. uname -s)"
}
