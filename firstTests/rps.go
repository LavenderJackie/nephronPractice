import(
    "fmt"
    "math"
)

func promptUser () {
  fmt.Println("rock paper scissors SHOOT")
  return 2
}

func main () {
  fmt.Println("welcome to rock paper scissors, how many games max??")
  var m int
  fmt.Scanln(&m)

  score := 0
  cScore := 0
  n := m/2 + 1

  for cScore < n && score < n
    cRoll := Math.RandInt(3)
    roll := promptUser()
    fmt.Println(cRoll)
    fmt.Println(roll)
}
