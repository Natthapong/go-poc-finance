package main

import (
	"fmt"
	"poc-finance/goalseek"
	"github.com/alpeb/go-finance/fin"
)

func myCommission(input float64) float64 {
	item_price := 5.0     
    comission  := 0.1    
    return item_price*input*(1-comission)
}

func myPMT(input float64) float64 {
	//fmt.Printf("input = %f\n", input)
	pmt, _ := fin.Payment(input, 60, -30000, 0, fin.PayEnd)
	//fmt.Printf("pmt = %f\n", pmt)
	return pmt
}

func myIRR(input float64) float64 {
	//fmt.Printf("input = %f\n", input)
	installment := (100000*(1+input))/10
	//fmt.Printf("installment = %f\n", installment)
	var values = []float64{-100000.00, installment, installment, installment, installment, installment, installment, installment, installment, installment, installment}
	irr, _ := fin.InternalRateOfReturn(values, 0.1)
	return irr
}

func main() {
	fmt.Println("Hello World Finance")

	result1 :=  goalseek.Run(myCommission, 1000, 25)
	fmt.Printf("result1 =  %#v\n",result1) 
	fmt.Printf("rate = %.6f%s \n", result1.ClosetValue*100, " %")

	result2 :=  goalseek.Run(myPMT, 2000, 25)
	fmt.Printf("result2 =  %#v\n",result2) 
	fmt.Printf("rate = %.6f%s \n", result2.ClosetValue*100, " %")

	result3 :=  goalseek.Run(myIRR, 0.05, 25)
	fmt.Printf("result3 =  %#v\n",result3) 
	fmt.Printf("installment = %f\n", (100000*(1+result3.ClosetValue))/10)

}