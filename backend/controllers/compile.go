package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"os/exec"

	"github.com/sohibjon7731/go-compiler/dto"
)

func CompileHandler(w http.ResponseWriter, r *http.Request){
	var req dto.CodeRequest
	json.NewDecoder(r.Body).Decode(&req)

	tmpFile:= "/tmp/code.go"
	err:= os.WriteFile(tmpFile, []byte(req.Code), 0644)
	if err != nil {
		json.NewEncoder(w).Encode(dto.CodeResponse{Error: "Cannot write file"})
		return
	}

	cmd:= exec.Command("go", "run", tmpFile)
	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	
	cmd.Run()
	

	res:=dto.CodeResponse{}
	if errBuf.Len()>0 {
		res.Error = errBuf.String()
	}else{
		res.Output = outBuf.String()
	}

	json.NewEncoder(w).Encode(res)

}