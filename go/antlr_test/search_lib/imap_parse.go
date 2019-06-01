package main

import (
	"fmt"
	"bytes"
	goimap "github.com/emersion/go-imap"
)

func testImapParse(){
	expected := `(1:42 UID 743:938 ` +
			`SINCE "5-Nov-1984" BEFORE "21-Nov-1997" SENTSINCE "5-Nov-1984" SENTBEFORE "21-Nov-1997" ` +
			`FROM root@protonmail.com BODY "hey there" TEXT DILLE ` +
			`ANSWERED DELETED KEYWORD cc UNKEYWORD microsoft ` +
			`LARGER 4242 SMALLER 4342 ` +
			`NOT (SENTON "21-Nov-1997" HEADER Content-Type text/csv) ` +
			`OR (ON "5-Nov-1984" DRAFT FLAGGED UNANSWERED UNDELETED OLD) (UNDRAFT UNFLAGGED UNSEEN))`
	expected = `(1:42 UID 743:938 AND ` +
			`SINCE "5-Nov-1984" BEFORE "21-Nov-1997" SENTSINCE "5-Nov-1984" SENTBEFORE "21-Nov-1997" ` +
			`OR (ON "5-Nov-1984" DRAFT FLAGGED UNANSWERED UNDELETED OLD) (UNDRAFT UNFLAGGED UNSEEN))`
	// expected = `subject:test after:2019-01-02 and test or -from:jinlian and from:ping or to:ping $hello`
	fmt.Println("Start testImapParse")
	criteria := new(goimap.SearchCriteria)

	b := bytes.NewBuffer([]byte(expected))
	r := goimap.NewReader(b)
	fields, _ := r.ReadFields()
	fmt.Printf("All fields: %v\n", fields)
	fmt.Printf("All fields[0]: %v\n", fields[0])
	for _, f := range fields[0].([]interface{}) {
	// for _, f := range fields {
		fmt.Printf("Field: %v\n", f)
	}

	if err := criteria.ParseWithCharset(fields[0].([]interface{}), nil); err != nil {
		fmt.Printf("Cannot parse search criteria for error: %v \n", err)
	}
	fmt.Printf("Criteria %v", criteria)
}

func main() {
	testImapParse()
}