package main

import (
	"fmt"
	"github.com/ZeroOneAI/AISecu/cmd/scanner/imagepull"
	"github.com/ZeroOneAI/AISecu/cmd/scanner/imagepull/crane"
	"github.com/ZeroOneAI/AISecu/cmd/scanner/resultsender"
	"github.com/ZeroOneAI/AISecu/cmd/scanner/resultsender/localhost"
	"github.com/ZeroOneAI/AISecu/cmd/scanner/scanner"
	"github.com/ZeroOneAI/AISecu/cmd/scanner/scanner/trivy"
	"github.com/ZeroOneAI/AISecu/pkg/image"
	"os"
	"time"
)

func main() {
	var imageInfo image.InfoInterface = nil
	var imagePuller imagepull.Interface = nil
	var resultSender resultsender.Interface = nil
	var scan scanner.Interface = nil
	var err error = nil
	var resultFilePath = "result.json"

	imageInfo, err = image.NewInfo()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	imagePuller = crane.NewCrane()
	resultSender = localhost.NewLocalhost("3000")
	scan = trivy.NewTrivy()

	fmt.Println("hi")
	err = imagePuller.Pull(imageInfo)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
		return
	}

	fmt.Println("hi")
	err = scan.Scan(imageInfo, resultFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	fmt.Println("hi")
	err = resultSender.Send(resultFilePath)
	for i := 0; i < 10 && err != nil; i++ {
		fmt.Println(i, ":", err)
		time.Sleep(10 * time.Second)
		err = resultSender.Send(resultFilePath)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}
}
