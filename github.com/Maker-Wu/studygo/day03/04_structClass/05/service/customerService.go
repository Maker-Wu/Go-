package main

import (
	"github.com/Maker-Wu/studygo/day03/04_structClass/05/model"
	"github.com/Maker-Wu/studygo/day03/04_structClass/05/view"
)

type customerService struct {
	customers []*model.Customer
}

func main() {
	customerView := view.NewCustomerView()
	customerView.MainMenu()

}
