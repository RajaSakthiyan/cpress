package main

import "github.com/solusipse/cpress"

func main() {

    cpress.Initialize()
    cpress.Press_key(cpress.KEY_A)
    cpress.Press_key(cpress.KEY_B)
    cpress.Press_key(cpress.KEY_C)
    cpress.Hold_key(cpress.KEY_LEFTSHIFT)
    cpress.Press_key(cpress.KEY_D)
    cpress.Release_key(cpress.KEY_LEFTSHIFT)
    cpress.Press_key(cpress.KEY_E)
}