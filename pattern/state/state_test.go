package state

import "testing"

func TestInitialState(t *testing.T) {

	initalState := &Brouillon{}

	document := Document{state: initalState}

	if document.state.Render() != "Render in brouillon" {
		t.Errorf("Expected Render in brouillon")
	}

}

func TestChangeStateDocument(t *testing.T) {

	initalState := &Brouillon{}

	document := Document{state: initalState}

	document.ChangeState(&EnCours{})

	if document.state.Render() != "Render in en cours" {
		t.Errorf("Expected Render in en cours")
	}

}

func TestChangeDocumentState(t *testing.T) {

	initalState := &Brouillon{}

	document := NewDocument("doc1")

	initalState.ChangeDocument(document)

	document2 := NewDocument("doc2")

	initalState.ChangeDocument(document2)

	if initalState.document.Name != "doc2" {
		t.Errorf("Expected doc2")
	}

}