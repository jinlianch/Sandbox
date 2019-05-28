package test

import (
	"reflect"
    "Sandbox/go/gocc_test/foo/ast"
    "Sandbox/go/gocc_test/foo/lexer"
    "Sandbox/go/gocc_test/foo/parser"
    "testing"
)

func TestWorld(t *testing.T) {
    input := []byte(`hello gocc`)
    lex := lexer.NewLexer(input)
    p := parser.NewParser()
    st, err := p.Parse(lex)
    t.Logf("after parse %v %v", reflect.TypeOf(st), st)
    if err != nil {
        panic(err)
    }
    w, ok := st.(*ast.World)
    if !ok {
        t.Fatalf("This is not a world")
    }
    t.Logf("w: %v", w)
    if w.Name != `gocc` {
        t.Fatalf("Wrong world %v", w.Name)
    }
}
