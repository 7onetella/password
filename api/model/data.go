package model

import (
	"encoding/json"
)

// Password password record in search list
type Password struct {
	ID       string   `json:"-"`
	Title    string   `json:"title"`
	URL      string   `json:"url"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Notes    string   `json:"notes"`
	Tags     []string `json:"tags"`
}

// http://choly.ca/post/go-json-marshalling/

// MarshalJSON marshals json
func (p *Password) MarshalJSON() ([]byte, error) {
	type PasswordAlias Password
	return json.Marshal(&struct {
		Type       string         `json:"type"`
		ID         string         `json:"id"`
		Attributes *PasswordAlias `json:"attributes"`
	}{
		Type:       "passwords",
		ID:         p.ID,
		Attributes: (*PasswordAlias)(p),
	})
}

// UnmarshalJSON unmarshals json
func (p *Password) UnmarshalJSON(data []byte) error {
	type PasswordAlias Password
	aux := &struct {
		Type       string         `json:"type"`
		ID         string         `json:"id"`
		Attributes *PasswordAlias `json:"attributes"`
	}{
		Attributes: (*PasswordAlias)(p),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	p.ID = aux.ID

	return nil
}

// Metadata meta data in search list
type Metadata struct {
	Count int    `json:"count"`
	Size  int    `json:"size"`
	Next  string `json:"next"`
	Prev  string `json:"prev"`
}

// ListPasswordsInput input parameters for listing passwords
type ListPasswordsInput struct {
	Token    string
	Title    string
	Notes    string
	Tags     []string
	Username string
}

// ListPasswordsOutput list passwords output
type ListPasswordsOutput struct {
	Items    []Password `json:"data"`
	Metadata Metadata   `json:"links"`
	Token    string     `json:"token"`
}

// CreatePasswordOutput password create resonse
type CreatePasswordOutput struct {
	RID string `json:"rid"`
	ID  string `json:"id"`
}
