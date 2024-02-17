package main

import (
	"errors"
	"flag"
	"github.com/Gamazic/stubgen/internal"
	"io"
	"log"
	"os"
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
			log.Fatalf("Error reading source file (arg --%s %s): %s", inpFileFlag, args.InpFile, err)
		}
	} else { // read from stdin
		srcGoCode, err = io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("Error reading from stdin: %s", err)
		}
	}
	if len(srcGoCode) == 0 {
		log.Fatal("Empty source code")
	}

	allStubsSrc, err := genStubsFromSrc(srcGoCode)
	if err != nil {
		log.Fatalf("Error generate: %s", err)
	}

	if args.OutFile != "" {
		f, err := os.Create(args.OutFile)
		if err != nil {
			log.Fatalf("Error open output file: %s", err)
		}
		defer f.Close()
		_, err = f.Write(allStubsSrc)
		if err != nil {
			log.Fatalf("Error write to output file: %s", err)
		}
	} else {
		os.Stdout.Write(allStubsSrc)
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

func genStubsFromSrc(src []byte) ([]byte, error) {
	astInfo, err := internal.GetAstModule("", src)
	if err != nil {
		return nil, err
	}
	if len(astInfo.InterfaceTypes) == 0 {
		return nil, errors.New("no interfaces were found")
	}
	packName := internal.ParsePackageName(astInfo.File)
	imports := internal.ParseImports(astInfo.Imports)
	interfaces := make([]internal.Interface, 0)
	for _, interfaceAst := range astInfo.InterfaceTypes {
		interfaces = append(interfaces, internal.ParseInterface(interfaceAst))
	}
	outSrc := internal.GenOutSrcFromParsed(packName, imports, interfaces)
	return outSrc, nil
}
