package main

type InterfaceProvider struct //User Interface Provider
{

}

func (InterfaceProvider) provideObject() {
	return UserInterface
}
