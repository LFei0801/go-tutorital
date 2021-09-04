package main

type Account struct {
	name string
}

// ChangeName 更改账户的名字
func (a *Account) ChangeName(name string){
	a.name = name
}