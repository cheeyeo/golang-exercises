package main

import "fmt"

func main(){
  elements := map[string]string{
    "H": "Hydrogen",
    "He": "Helium",
     "Li": "Lithium",
     "Be": "Beryllium",
     "B": "Boron",
     "C": "Carbon",
     "N": "Nitrogen",
     "O": "Oxygen",
     "F": "Fluorine",
     "Ne": "Neon",
  }

  // name, ok := elements["Li"]
  // fmt.Println(name, ok)

  if name, ok := elements["Li"]; ok {
    fmt.Println(name, ok)
  }

  // for keys which don't exist, the zero type is returned
  // for strings it is the empty string
  // fmt.Println(elements["Un"])
}
