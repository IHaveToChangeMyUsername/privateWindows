package util

type DoorState struct {
	doorOpen bool
}

func (d *DoorState) IsDoorOpen() bool {
	return d.doorOpen
}

func (d *DoorState) SetDoorOpen(open bool) {
	d.doorOpen = open
}
