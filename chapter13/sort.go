package main

import (
  "fmt"
  "sort"
)

type Person struct{
  Name string
  Age int
}

func (p Person) String() string {
  return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []Person
func (b ByName) Len() int{
  return len(b)
}

func (b ByName) Swap(i, j int){
  b[i], b[j] = b[j], b[i]
}

func (b ByName) Less(i, j int) bool {
  return b[i].Name < b[j].Name
}

func main(){
  kids := []Person{
    {"Jill", 9},
    {"Jack", 10},
  }

  fmt.Println(kids)

  sort.Sort(ByName(kids))

  fmt.Println(kids)


  people := []Person{
    {"Bob", 31},
    {"John", 42},
    {"Michel", 17},
    {"Jenny", 26},
  }

  fmt.Println(people)

  sort.Sort(ByAge(people))

  fmt.Println(people)
}

