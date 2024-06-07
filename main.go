package main

import (
	"fmt"
	"strings"
)

type Context struct {
	CheckNumber        []string
	RuleNumber         []string
	CorrectNumberCount int
	CorrectPlaceCount  int
}

type CheckExpression interface {
	Interpret(context Context) bool
}

type String struct {
	value string
}

func (number String) ConverStringToList() []string {
	strNumber := number.value
	var numberArray = []string{}
	chars := strings.Split(strNumber, "")
	for _, str := range chars {
		numberArray = append(numberArray, str)
	}
	return numberArray
}

func getRangeNumbers(from int, to int) []int {
	var array = []int{}
	for i := from; i < to; i++ {
		array = append(array, i)
	}
	return array
}

func fillPuzzleData(ctx *[]Context) {
	*ctx = append(*ctx, Context{RuleNumber: []string{"6", "9", "0"},
		CorrectNumberCount: 1, CorrectPlaceCount: 1})
	*ctx = append(*ctx, Context{RuleNumber: []string{"7", "4", "1"},
		CorrectNumberCount: 1, CorrectPlaceCount: 0})
	*ctx = append(*ctx, Context{RuleNumber: []string{"5", "0", "4"},
		CorrectNumberCount: 2, CorrectPlaceCount: 0})
	*ctx = append(*ctx, Context{RuleNumber: []string{"3", "8", "7"},
		CorrectNumberCount: 0, CorrectPlaceCount: 0})
	*ctx = append(*ctx, Context{RuleNumber: []string{"2", "1", "9"},
		CorrectNumberCount: 1, CorrectPlaceCount: 0})
}

/*func fillPuzzleData(ctx *[]Context) {
	*ctx = append(*ctx, Context{RuleNumber: []string{"6", "8", "2"},
		CorrectNumberCount: 1, CorrectPlaceCount: 1})
	*ctx = append(*ctx, Context{RuleNumber: []string{"6", "1", "4"},
		CorrectNumberCount: 1, CorrectPlaceCount: 0})
	*ctx = append(*ctx, Context{RuleNumber: []string{"2", "0", "6"},
		CorrectNumberCount: 2, CorrectPlaceCount: 0})
	*ctx = append(*ctx, Context{RuleNumber: []string{"7", "3", "8"},
		CorrectNumberCount: 0, CorrectPlaceCount: 0})
	*ctx = append(*ctx, Context{RuleNumber: []string{"7", "8", "0"},
		CorrectNumberCount: 1, CorrectPlaceCount: 0})
}*/

/*func fillPuzzleData(ctx *[]Context) {
	*ctx = append(*ctx, Context{RuleNumber: []string{"4", "2", "8"},
		CorrectNumberCount: 1, CorrectPlaceCount: 0})
	*ctx = append(*ctx, Context{RuleNumber: []string{"3", "6", "8"},
		CorrectNumberCount: 1, CorrectPlaceCount: 1})
	*ctx = append(*ctx, Context{RuleNumber: []string{"3", "2", "4"},
		CorrectNumberCount: 1, CorrectPlaceCount: 0})
	*ctx = append(*ctx, Context{RuleNumber: []string{"2", "4", "5"},
		CorrectNumberCount: 2, CorrectPlaceCount: 0})
}*/

type Pair[T, U, Z any] struct {
	First  T
	Second U
	Index  Z
}

func Zip[T, U, Z any](ts []T, us []U, zs []Z) []Pair[T, U, Z] {
	if len(ts) != len(us) {
		panic("slices have different length")
	}
	pairs := make([]Pair[T, U, Z], len(ts))
	for i := 0; i < len(ts); i++ {
		pairs[i] = Pair[T, U, Z]{ts[i], us[i], zs[i]}
	}
	return pairs
}

type RuleExpression struct{}

func (rex RuleExpression) Interpret(context Context) bool {
	var sameIndexCount = 0
	var sameCount = 0
	var index = getRangeNumbers(0, 3)

	for _, pair := range Zip(context.CheckNumber, context.RuleNumber, index) {
		if pair.First == pair.Second {
			sameIndexCount++
			sameCount++
		} else {
			var count = 0
			for index, rulNum := range context.RuleNumber {
				if rulNum == pair.First && index != pair.Index {
					count++
				}
			}
			if count > 0 {
				sameCount++
			}
		}
	}
	return sameIndexCount == context.CorrectPlaceCount && sameCount == context.CorrectNumberCount
}

var rex = RuleExpression{}

func CheckAllRules(number string) bool {
	for index, _ := range contextList {
		contextList[index].CheckNumber = String{number}.ConverStringToList()
		if !rex.Interpret(contextList[index]) {
			return false
		}
	}
	return true
}

var contextList = []Context{}

func main() {
	//contextList = make([]Context, 0)
	fillPuzzleData(&contextList)
	var resultList = []string{}
	for i := 1; i <= 999; i++ {
		str := fmt.Sprintf("%03d", i)
		if CheckAllRules(str) {
			resultList = append(resultList, str)
		}
	}
	for _, result := range resultList {
		fmt.Println(result)
	}
	//fmt.Println(len(resultList))
	if len(resultList) == 0 {

	}
}
