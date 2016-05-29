// Package menu provides a command-line menu interface.
package menu

import (
  "fmt"
  "strconv"
  "strings"
  "bufio"
  "os"
)

type Option struct {
  Choice string
  Desc   string
}

type Menu struct {
  options []Option
}

func New(options []Option) *Menu {
  return &Menu{
    options: options,
  }
}

func (m *Menu) Ask() string {
  fmt.Println("Options: ")
  for i, opt := range m.options {
    fmt.Printf("\t%d: %s\n", i, opt.Desc)
  }
  for {
    fmt.Print(" Choice > ")
    input := bufio.NewReader(os.Stdin)
    choice, err := input.ReadString('\n')
    if err != nil {
      return ""
    }
    c, err := strconv.Atoi(strings.TrimSpace(choice))
    if err != nil || c < 0 || c >= len(m.options) {
      fmt.Print("'%d' is not a valid choice.", choice)
    } else {
      return m.options[c].Choice
    }
  }
}
