package main

import "fmt"

func main(){
  elements := make(map[string]string)
  elements["H"] = "Hydrogen"
  elements["He"] = "Helium"
  elements["Li"] = "Lithium"
  elements["Be"] = "Beryllium"
  elements["B"] = "Boron"
  elements["C"] = "Carbon"
  elements["N"] = "Nitrogen"
  elements["O"] = "Oxygen"
  elements["F"] = "Fluorine"
  elements["Ne"] = "Neon"

  // name, ok := elements["Li"]
  // fmt.Println(name, ok)

  if name, ok := elements["Li"]; ok {
    fmt.Println(name, ok)
  }

  // for keys which don't exist, the zero type is returned
  // for strings it is the empty string
  // fmt.Println(elements["Un"])
}
