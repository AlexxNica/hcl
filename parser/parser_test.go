package parser

import "testing"

func TestAssignStatment(t *testing.T) {
	src := `ami = "${var.foo}"`
	p := New([]byte(src))
	p.enableTrace = true
	n, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}

	if n.String() != src {
		t.Errorf("AssignStatement is not parsed correctly\n\twant: '%s'\n\tgot : '%s'", src, n.String())
	}

	if n.Pos().Line != 1 {
		t.Errorf("AssignStatement position is wrong\n\twant: '%d'\n\tgot : '%d'", 1, n.Pos().Line)
	}

}
