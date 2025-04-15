package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/sohibjon7731/go-compiler/dto"
)

func FormatHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.CodeRequest
	json.NewDecoder(r.Body).Decode(&req)

	err := os.WriteFile("/tmp/code.go", []byte(req.Code), 0644)
	if err != nil {
		fmt.Println("Error Writing file")
	}
	cmd := exec.Command("gofmt", "/tmp/code.go")
	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	err = cmd.Run()

	res := dto.CodeResponse{}
	if err != nil || errBuf.Len() > 0 {
		res.Error = errBuf.String()
	} else {
		res.Output = outBuf.String()
	}

	json.NewEncoder(w).Encode(res)

}
