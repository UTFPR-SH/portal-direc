// Copyright 2018 Portal Direc's authors. All rights reserved.
//
// Authors: Rafael Campos Nunes <rafaelnunes@alunos.utfpr.edu.br>
//          ...                 <email>
//
// Use of this source code is governed by a Apache License. The license can be
// found (LICENSE) at the root of the project.

// DOCUMENTATION
// ---------------------------------------------------------------------------
// The entry point to the entire system. It interprets flags and start the
// server at a desired <port> value, here called of address, or if no port
// value is assigned the default one is 8080.

package main

import (
	"flag"
)

var (
	addr string
)

func init() {
	flag.StringVar(&addr, "p", "8080", "Port address used to start the server")
}

func main() {

	flag.Parse()

	//Init the server at port address
	Init(addr)
}