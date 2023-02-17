package utils

import (
	"fmt"
	"testing"
)

func TestSetStructValue(t *testing.T) {
	type MobileReq struct {
		Mobile   string `json:"mobile"`
		Areacode int    `json:"areacode" default:"86"`
	}
	var req MobileReq
	if err := SetStructValue(&req); err != nil {
		panic(err)
	}
	fmt.Printf("%#+v\n", req)

	mobileReq := MobileReq{Mobile: "18842878899"}
	if err := SetStructValue(&mobileReq); err != nil {
		panic(err)
	}
	fmt.Printf("%#+v\n", mobileReq)

	mobileSlice := []*MobileReq{
		&MobileReq{
			Mobile: "18842878899",
		},
	}
	if err := SetStructValue(&mobileSlice); err != nil {
		panic(err)
	}
	fmt.Printf("%#+v\n", mobileSlice[0])

	mobileMap := map[string]*MobileReq{
		"liming": &MobileReq{
			Mobile: "18842878899",
		},
	}
	if err := SetStructValue(&mobileMap); err != nil {
		panic(err)
	}
	fmt.Printf("%#+v\n", mobileMap["liming"])
}

func TestVerifyStructFieldRequired(t *testing.T) {
	type MobileReq struct {
		Mobile   string `json:"mobile" required:"true"`
		Areacode int    `json:"areacode" default:"86"`
	}
	var req MobileReq
	if err := SetStructValue(&req, WithVerifyStructFieldRequired()); err != nil {
		fmt.Println(err)
	}

	mobileReq := MobileReq{Mobile: "18842878899"}
	if err := SetStructValue(&mobileReq, WithVerifyStructFieldRequired()); err != nil {
		panic(err)
	}
	fmt.Printf("%#+v\n", mobileReq)

	mobileSlice := []*MobileReq{
		&MobileReq{},
	}
	if err := SetStructValue(&mobileSlice, WithVerifyStructFieldRequired()); err != nil {
		fmt.Println(err)
	}

	mobileMap := map[string]*MobileReq{
		"liming": &MobileReq{},
	}
	if err := SetStructValue(&mobileMap, WithVerifyStructFieldRequired()); err != nil {
		fmt.Println(err)
	}
}

func TestVerifyStructFieldRange(t *testing.T) {
	type MobileReq struct {
		Mobile   string `json:"mobile"`
		Areacode int    `json:"areacode" range:"86:9999"`
	}
	var req MobileReq
	if err := SetStructValue(&req, WithVerifyStructFieldRange()); err != nil {
		fmt.Println(err)
	}

	mobileReq := MobileReq{Mobile: "18842878899", Areacode: 86}
	if err := SetStructValue(&mobileReq, WithVerifyStructFieldRange()); err != nil {
		panic(err)
	}
	fmt.Printf("%#+v\n", mobileReq)
}

func TestVerifyStructFieldOptions(t *testing.T) {
	type MobileReq struct {
		Mobile   string `json:"mobile"`
		Areacode int    `json:"areacode" options:"1,86,321"`
	}
	var req MobileReq
	if err := SetStructValue(&req, WithVerifyStructFieldOptions()); err != nil {
		fmt.Println(err)
	}

	mobileReq := MobileReq{Mobile: "18842878899", Areacode: 86}
	if err := SetStructValue(&mobileReq, WithVerifyStructFieldOptions()); err != nil {
		panic(err)
	}
	fmt.Printf("%#+v\n", mobileReq)
}
