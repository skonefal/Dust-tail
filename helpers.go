package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func createJsonArrayOfArrays(file string) {
	//	fd, err := os.Open(file)
	//	if err != nil {
	//		fmt.Println("os.open error" + err.Error())
	//		return
	//	}

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("os.readall error" + err.Error())
		return
	}

	//	ioutil.wr

	usage := strings.Replace(string(bytes), "][", "],[", -1)
	err = ioutil.WriteFile(file, []byte("["+usage+"]"), 0700)
	if err != nil {
		fmt.Println("ioutil.WriteFile error " + err.Error())
		return
	}

	//	_, err = fd.WriteString("[" + tempUsage + "]")

	//	fmt.Println("[" + tempUsage + "]")

}

func createResulsFilename(endpoint string) (string, error) {
	nodeNameArr := nodeRegexp.FindStringSubmatch(endpoint)
	nodeName := nodeNameArr[2]

	resultsFile := path.Join(EXPERIMENT_RESULTS_FOLDER, nodeName+"_"+experimentStartTime.String())

	// TODO cache filenames to reduce OS calls
	if _, err := os.Stat(resultsFile); err != nil {
		if os.IsNotExist(err) {
			os.Create(resultsFile)
		} else {
			fmt.Printf("Error while accessing folder | %s", err)
			return "", err
		}
	}

	return resultsFile, nil
}

func createResultsPath(dirPath string) error {
	if _, err := os.Stat(dirPath); err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dirPath, 0700)
			if err != nil {
				fmt.Printf("Error while creating path %s | %s", dirPath, err)
				return err
			}
		} else {
			fmt.Printf("Error while creating path %s | %s", dirPath, err)
			return err
		}
	}
	return nil
}
