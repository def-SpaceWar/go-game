package main

type Entity uint32

func (e Entity) GetComponent(compType *string) (*Component, error) {
    return currentWorld.GetComponent(e, compType)
}
