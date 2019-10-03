package main

import "testing"

func TestEncodeID(t *testing.T) {

	spec := GSpec{t}

	spec.Given("id=1234")

	spec.When("EncodeID(1234)")

	token, err := EncodeID("1234")

	spec.Then()

	spec.AssertAndFailNow(err == nil, "error is nil", err)

	spec.AssertAndFailNow(len(token) != 0, "len of token is not zero", len(token))

}

func TestDecodeToken(t *testing.T) {

	tokenString, err := EncodeID("1234")
	if err != nil {
		t.Fatalf("error while encoding token: %v", err)
	}

	spec := GSpec{t}

	spec.Given("token=" + tokenString)

	spec.When("DecodeToken(" + tokenString + ")")

	ID, err := DecodeToken(tokenString)

	spec.AssertAndFailNow(err == nil, "error is nil", err)

	spec.AssertAndFailNow(len(ID) > 0, "length of ID > 0", len(ID))

}
