package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/Gamazic/stubgen/internal"
)

const (
	inpFileFlag = "inp-file"
	outFileFlag = "out-file"
	debugFlag   = "debug"
)

var (
	flagInputFile  = flag.String(inpFileFlag, "", "path to file with interfaces")
	flagOutputFile = flag.String(outFileFlag, "", "path to file to save stubs")
	flagDebug      = flag.Bool(debugFlag, false, "print internal error details")
)

func main() {
	flag.Parse()
	defer handlePanic()

	args, err := parseArgs()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	var srcGoCode []byte
	if args.InpFile != "" { // read from file
		srcGoCode, err = os.ReadFile(args.InpFile)
		if err != nil {
			log.Fatalf("Error reading source file (--%s %s): %s", inpFileFlag, args.InpFile, err)
		}
	} else { // read from stdin
		srcGoCode, err = io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("Error reading: %s", err)
		}
		if len(srcGoCode) == 0 {
			log.Fatal("Empty stdin input module")
		}
	}

	if len(srcGoCode) == 0 {
		log.Fatal("Empty source code file")
	}
	allStubsSrc, err := genStubsFromSrc(string(srcGoCode))
	if err != nil {
		log.Fatalf("Error generate: %s", err)
	}

	if args.OutFile != "" {
		f, err := os.Create(args.OutFile)
		if err != nil {
			log.Fatalf("Error open output file: %s", err)
		}
		defer f.Close()
		_, err = f.WriteString(allStubsSrc)
		if err != nil {
			log.Fatalf("Error write to output file: %s", err)
		}
	} else {
		fmt.Println(allStubsSrc)
	}

}

func handlePanic() {
	if !*flagDebug {
		if r := recover(); r != nil {
			log.Fatalf("Internal error is occurred, "+
				"please report the problem to the github issue. "+
				"For the details set flag --%s.\nError: %s", debugFlag, r)
		}
	}
}

type appArgs struct {
	InpFile string
	OutFile string
}

func parseArgs() (appArgs, error) {
	var inpF, outF string
	if flagInputFile != nil {
		inpF = *flagInputFile
	}
	if flagOutputFile != nil {
		outF = *flagOutputFile
	}
	return appArgs{
		InpFile: inpF,
		OutFile: outF,
	}, nil
}

func genStubsFromSrc(src string) (string, error) {
	interfacesAst, err := internal.GetAstInterfaces("", src)
	if err != nil {
		return "", err
	}
	if len(interfacesAst) == 0 {
		return "", errors.New("no interfaces were found")
	}
	stubSrc := make([]string, 0)
	for _, interfaceAst := range interfacesAst {
		parsedInterface := internal.ParseInterface(interfaceAst)
		interfaceSrc := internal.GenStubFromInterface(parsedInterface)
		stubSrc = append(stubSrc, interfaceSrc)
	}
	allStubsSrc := strings.Join(stubSrc, "\n")
	return allStubsSrc, nil
}
