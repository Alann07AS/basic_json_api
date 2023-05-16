package args

import "os"

func ReadArgs(argsNumber int) (arg string) {
	args := os.Args
	if len(args) == 1 {
		return
	}
	return args[argsNumber+1]
}
