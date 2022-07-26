package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/thientran2020/financial-cli/models"
)

const (
	ColorOff = "\033[0m"

	// Regular Colors
	Black  = "\033[0;30m"
	Red    = "\033[0;31m"
	Green  = "\033[0;32m"
	Yellow = "\033[0;33m"
	Blue   = "\033[0;34m"
	Purple = "\033[0;35m"
	Cyan   = "\033[0;36m"
	White  = "\033[0;37m"

	// Bold Colors
	BBlack  = "\033[1;30m"
	BRed    = "\033[1;31m"
	BGreen  = "\033[1;32m"
	BYellow = "\033[1;33m"
	BBlue   = "\033[1;34m"
	BPurple = "\033[1;35m"
	BCyan   = "\033[1;36m"
	BWhite  = "\033[1;37m"

	// Underline Colors
	UBlack  = "\033[4;30m"
	URed    = "\033[4;31m"
	UGreen  = "\033[4;32m"
	UYellow = "\033[4;33m"
	UBlue   = "\033[4;34m"
	UPurple = "\033[4;35m"
	UCyan   = "\033[4;36m"
	UWhite  = "\033[4;37m"

	UnderlineCommandColor = UGreen + "%s" + ColorOff
	BoldCommandColor      = BGreen + "%s" + ColorOff
	RedCommandColor       = Red + "%s" + ColorOff
	CheckMark             = "\u2713"
)

func ConfirmYesNoPromt(label string) bool {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}
	_, err := prompt.Run()
	return err == nil
}

func MultiSelect(label string, items []string) (string, error) {
	prompt := promptui.Select{
		Label:  label,
		Items:  items,
		Size:   len(items),
		Stdout: &BellSkipper{},
	}
	_, result, err := prompt.Run()

	if err != nil {
		return "", err
	}
	return result, nil
}

func NumberEnter(label string) (int64, error) {
	validate := func(input string) error {
		number, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			return errors.New("Invalid number")
		}
		if number < 0 {
			return errors.New("Negative number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}

	stringNum, err := prompt.Run()
	if err != nil {
		return -1, err
	}

	result, _ := strconv.ParseInt(stringNum, 10, 64)
	return result, nil
}

func PromptEnter(label string) (string, error) {
	validate := func(input string) error {
		if len(input) == 0 {
			return errors.New("Invalid input")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}

	result, err := prompt.Run()

	return result, err
}

func PrintCustomizedMessage(message string, color string, newline bool) {
	message = strings.ReplaceAll(message, ColorOff, "")
	if newline {
		fmt.Printf("%s%s\n", color, message)
	} else {
		fmt.Printf("%s%s", color, message)
	}
}

func PrintSingleRecord(record models.Record, color string) {
	date := fmt.Sprintf("    %s/%s/%s   ", getStringDate(record.Month), getStringDate(record.Day), strconv.Itoa(record.Year))
	description := fmt.Sprintf(" %-35s", record.Description)
	costString := fmt.Sprintf(" $%-6s", strconv.Itoa(record.Cost))
	category := fmt.Sprintf(" %-18s", record.Category)
	dash := "|-----|-----|-----|------------------------------------|--------|-------------------|"

	message := "\n" + dash + "\n" + fmt.Sprintf("|%s|%s|%s|%s|", date, description, costString, category) + "\n" + dash + "\n"
	PrintCustomizedMessage(message, color, true)
}

func getStringDate(number int) string {
	if number < 10 {
		return "0" + strconv.Itoa(number)
	}
	return strconv.Itoa(number)
}

type BellSkipper struct{}

func (bs *BellSkipper) Write(b []byte) (int, error) {
	const charBell = 7
	if len(b) == 1 && b[0] == charBell {
		return 0, nil
	}
	return os.Stderr.Write(b)
}

func (bs *BellSkipper) Close() error {
	return os.Stderr.Close()
}
