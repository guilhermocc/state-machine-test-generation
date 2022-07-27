package devicestates

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDevice(t *testing.T) {
	t.Run("TestFirstPath", func(t *testing.T) {
		device := &Device{State: "OFF"}
		assert.Equal(t, "OFF", device.GetState())
		device.Home()
		assert.Equal(t, "LOCKED", device.GetState())
		device.Login("wrong", "wrong")
		assert.Equal(t, "LOCKED", device.GetState())
	})
	t.Run("TestFirstPath", func(t *testing.T) {
		device := &Device{State: "OFF"}
		assert.Equal(t, "OFF", device.GetState())
		device.Home()
		assert.Equal(t, "LOCKED", device.GetState())
		device.Login("login", "password")
		assert.Equal(t, "UNLOCKED", device.GetState())
		device.LockButton()
		assert.Equal(t, "LOCKED", device.GetState())
	})
	t.Run("TestFirstPath", func(t *testing.T) {
		device := &Device{State: "OFF"}
		assert.Equal(t, "OFF", device.GetState())
		device.Home()
		assert.Equal(t, "LOCKED", device.GetState())
		device.Login("login", "password")
		assert.Equal(t, "UNLOCKED", device.GetState())
		device.LongLockButton()
		assert.Equal(t, "OFF", device.GetState())
	})
	t.Run("TestFirstPath", func(t *testing.T) {
		device := &Device{State: "OFF"}
		assert.Equal(t, "OFF", device.GetState())
		device.Home()
		assert.Equal(t, "LOCKED", device.GetState())
		device.LongLockButton()
		assert.Equal(t, "OFF", device.GetState())
	})
}
