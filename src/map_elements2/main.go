package main

import "fmt"

func main() {
  elements := map[string]string {
    "H": "Hydrogen",
    "He": "Helium",
    "Li": "Lithium",
    "Be": "Beryllium",
    "B": "Boron",
    "C": "Carbon",
    "N": "Nitrogen",
    "O": "Oxygen",
    "F": "Flourine",
    "Ne": "Neon",
  }


  fmt.Println(elements["Li"])

  name, ok := elements["Un"]
  fmt.Println(name, ok)
}
