package main

type Farmer struct {
	keychain_proxy struct{} // @TODO: Decide how to deal with this; which library to use for the keychain access

}

func (f *Farmer) New() {}
