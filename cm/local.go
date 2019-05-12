package cm

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

// RunLocalCommand run local system command
func RunLocalCommand(command string, arg ...string) (string, string, error) {
	outStr, errStr := "", ""
	cmd := exec.Command(command, arg...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	outStr, errStr = string(stdout.Bytes()), string(stderr.Bytes())
	if len(outStr) > 0 {
		fmt.Println(outStr)
	}
	if len(errStr) > 0 {
		fmt.Println(errStr)
	}
	return outStr, errStr, nil
}

// DownloadFile download a file from an url to the local filesystem
func DownloadFile(filepath string, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

// DownloadFileInMemory download file as bytes
func DownloadFileInMemory(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytesResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bytesResp, nil
}
