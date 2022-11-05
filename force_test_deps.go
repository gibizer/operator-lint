package main

import (
	// analysistest does not work with Go Modules yet so we need to list
	// all the dependencies we use in the test case
	_ "github.com/onsi/gomega"
)
