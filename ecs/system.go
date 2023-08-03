package ecs

type WorldState uint8
const CONTINUE WorldState = 0
const QUIT WorldState = 1
const SWITCH WorldState = 2

type System func(*World) (WorldState, *World)
