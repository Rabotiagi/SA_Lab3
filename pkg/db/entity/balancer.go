package entity

import "fmt"

type Balancer struct {
	Id                int       `gorm:"primary_key, AUTO_INCREMENT"`
	ConnectedMachines []Machine `gorm:"foreignkey:BalancerId"`
}

type BalancerData struct {
	Id                 int
	UsedMachines       []int
	TotalMachinesCount int
}

func (b *Balancer) Print() {
	fmt.Printf("Balancer ID: %d \n", b.Id)
}
