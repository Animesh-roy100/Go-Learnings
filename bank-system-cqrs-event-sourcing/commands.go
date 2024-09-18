package main

type Command interface{}

type CreateAccount struct {
	name string
}
