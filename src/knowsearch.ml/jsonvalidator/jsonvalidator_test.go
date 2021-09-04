package jsonvalidator

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestCheckFPExistent(t *testing.T) {
	exist, err := CheckFP("./jsonvalidator.go")
	if err != nil {
		t.Error(err)
	}
	if !exist {
		t.Errorf("Existent File Not Detected")
	}
}

func TestCheckFPNonExistent(t *testing.T) {
	exist, _ := CheckFP("./nonexist.go")
	if exist {
		t.Errorf("Non Existent File Detected")
	}
}

func TestPrintFileLineError(t *testing.T) {
	testStr := `{
		"checkme" : 12,
		}	
	}`
	compareStr := `{
		"checkme" : 12,<==== CHECK HERE
		}	
	}`

	// the below lines of code capture what was printed out by the function
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	printFileLineError(testStr, 2)

	w.Close()

	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	funcOut := string(out)

	if strings.TrimSpace(funcOut) != strings.TrimSpace(compareStr) {
		t.Errorf("The Test String and Output dont match \n%s\n%s", funcOut, compareStr)
	}

}

func TestValidateJSONCorrect(t *testing.T) {
	testStr := `{
		"checkme" : 12,
	}`
	check := ValidateJSON(testStr, 0)
	if !check {
		t.Errorf("Correct JSON classified as incorrect")
	}
}

func TestValidateJSONBracketError(t *testing.T) {
	testStr := `{
		"checkme" : 12,
		}	
	}`
	check := ValidateJSON(testStr, 0)
	if check {
		t.Errorf("Incorrect JSON classified as Correct :: Bracket Error")
	}
}

func TestValidateJSONQuoteError(t *testing.T) {
	testStr := `{
		"checkme" : "12,
	}`
	check := ValidateJSON(testStr, 0)
	if !check {
		t.Errorf("Correct JSON classified as incorrect :: Quote Error")
	}
}
