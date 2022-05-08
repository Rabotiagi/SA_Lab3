package entity

import "fmt"

type Balancer struct {
	Id                int       `gorm:"primary_key, AUTO_INCREMENT"`
	ConnectedMachines []Machine `gorm:"foreignkey:BalancerId"`
}

func (b *Balancer) Print() {
	fmt.Printf("Balancer ID: %d \n", b.Id)
}
