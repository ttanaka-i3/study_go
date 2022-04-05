package myLib

func Average(s []int) float64 {
  total := 0
  for i, _ := range s{
    total += s[i]
  }
  return float64(total) / float64(len(s))
}
