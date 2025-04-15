package controllers

import (
	"encoding/json"
	"go/parser"
	"go/token"
	"net/http"

	"github.com/sohibjon7731/go-compiler/dto"
)

func CheckSyntax(w http.ResponseWriter, r *http.Request) {
	var req dto.CodeRequest
	json.NewDecoder(r.Body).Decode(&req)

	fset := token.NewFileSet()
	_, err := parser.ParseFile(fset, "main.go", req.Code, parser.AllErrors)
	res := dto.CodeResponse{}
	if err != nil {
		res.Error = err.Error()
	}
	json.NewEncoder(w).Encode(res)
}
