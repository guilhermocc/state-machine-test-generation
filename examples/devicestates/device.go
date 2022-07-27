package devicestates

type Device struct {
	State string
}

func (d *Device) GetState() string {
	return d.State
}

func (d *Device) Home() {
	if d.State == "OFF" {
		d.State = "LOCKED"
	}
}

func (d *Device) Login(user, pass string) {
	var correctPass bool
	if user == "login" && pass == "password" {
		correctPass = true
	} else {
		correctPass = false
	}

	if correctPass {
		d.State = "UNLOCKED"
	} else {
		d.State = "LOCKED"
	}
}

func (d *Device) LockButton() {
	d.State = "LOCKED"
}

func (d *Device) LongLockButton() {
	d.State = "OFF"
}
