package main

import (
        "log"
        "net"
        "os"

        "github.com/Adriankv1/is105sem03/mycrypt"
)

func main() {
        conn, err := net.Dial("tcp", "172.17.0.3:33419")
        if err != nil {
                log.Fatal(err)
        }

        kryptertMelding := mycrypt.Krypter([]rune(os.Args[1]), mycrypt.ALFSEM03, 4)
        log.Println("Kryptert melding: ", string(kryptertMelding))
        , err = conn.Write([]byte(string(kryptertMelding)))

        log.Println("os.Args[1] = ", os.Args[1])

        _, err = conn.Write([]byte(os.Args[1]))
        if err != nil {
                log.Fatal(err)
        }")
        if err != nil {
                log.Fatal(err)
        }

        kryptertMelding := mycrypt.Krypter([]rune(os.Args[1]), mycrypt.ALFSEM03, 4)
        log.Println("Kryptert melding: ", string(kryptertMelding))
        , err = conn.Write([]byte(string(kryptertMelding)))

        log.Println("os.Args[1] = ", os.Args[1])

        _, err = conn.Write([]byte(os.Args[1]))
        if err != nil {
                log.Fatal(err)
        }
}

