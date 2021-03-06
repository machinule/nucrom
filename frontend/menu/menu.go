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
    trimmed := strings.TrimSpace(choice)
    c, err := strconv.Atoi(trimmed)
    if err != nil {
      fmt.Printf("'%s' is not a valid choice. Please choose a number.", trimmed)
    } else if c < 0 || c >= len(m.options) {
      fmt.Printf("'%d' is not a valid choice. Please choose one of the options.", c)
    } else {
      return m.options[c].Choice
    }
  }
}
