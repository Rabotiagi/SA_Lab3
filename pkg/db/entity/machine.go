package entity

import "fmt"

type Machine struct {
	Id         int `gorm:"primary_key, AUTO_INCREMENT"`
	State      int
	BalancerId int `gorm:"column:balancer_id"`
}

type MachineData struct {
	State int `json:"state"`
}

func (m *Machine) Print() {
	fmt.Printf("Machine Id: %d| State: %d | Balancer ID: %d \n", m.Id, m.State, m.BalancerId)
}
