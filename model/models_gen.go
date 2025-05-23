// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type BasicInfo struct {
	ID             string  `json:"id"`
	FirstName      string  `json:"firstName"`
	LastName       string  `json:"lastName"`
	AdditionalName *string `json:"AdditionalName,omitempty"`
	Pronoun        *string `json:"pronoun,omitempty"`
	Headline       *string `json:"headline,omitempty"`
}

type CreateBasicInfoInput struct {
	FirstName      string  `json:"firstName"`
	LastName       string  `json:"lastName"`
	AdditionalName *string `json:"AdditionalName,omitempty"`
	Pronoun        *string `json:"pronoun,omitempty"`
	Headline       *string `json:"headline,omitempty"`
}

type Mutation struct {
}

type Query struct {
}

type UpdateBasicInfoInput struct {
	FirstName      *string `json:"firstName,omitempty"`
	LastName       *string `json:"lastName,omitempty"`
	AdditionalName *string `json:"AdditionalName,omitempty"`
	Pronoun        *string `json:"pronoun,omitempty"`
	Headline       *string `json:"headline,omitempty"`
}
