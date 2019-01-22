package readme

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/openset/leetcode/internal/base"
	"github.com/openset/leetcode/internal/leetcode"
)

var CmdReadme = &base.Command{
	Run:       runReadme,
	UsageLine: "readme",
	Short:     "build README.md file",
	Long:      "build README.md file.",
}

func runReadme(cmd *base.Command, args []string) {
	if len(args) == 1 && args[0] == "page" {
		buildCmd = "page"
		fileName = "index.md"
		leetcode.LockStr = " &hearts;"
	}
	if len(args) > 1 {
		cmd.Usage()
	}
	var buf bytes.Buffer
	buf.WriteString(base.AuthInfo(buildCmd))
	buf.WriteString(defaultStr)
	writeProblems(&buf)
	base.FilePutContents(fileName, buf.Bytes())
}

func writeProblems(buf *bytes.Buffer) {
	problems := leetcode.ProblemsAll()
	problemsSet := make(map[int]string)
	maxId := 0
	for _, problem := range problems.StatStatusPairs {
		id := problem.Stat.FrontendQuestionId
		stat := problem.Stat
		title := strings.TrimSpace(problem.Stat.QuestionTitle)
		slug := problem.Stat.QuestionTitleSlug
		levelName := problem.Difficulty.LevelName()
		format := "| <span id=\"%d\">%d</span> | [%s](https://leetcode.com/problems/%s%s)%s | [%s](https://github.com/openset/leetcode/tree/master/problems/%s) | %s |\n"
		problemsSet[id] = fmt.Sprintf(format, id, id, title, slug, stat.TranslationTitle(), problem.PaidOnly.Str(), stat.Lang(), slug, levelName)
		if id > maxId {
			maxId = id
		}
	}
	// table
	step, long := 50, 300
	buf.WriteString("<table><thead>\n")
	for i := 0; i < maxId; i += long {
		buf.WriteString("<tr>\n")
		for j := 0; j < long/step; j++ {
			buf.WriteString(fmt.Sprintf("\t<th align=\"center\"><a href=\"#%d\">[%d-%d]</a></th>\n", 1+i+j*step, 1+i+j*step, i+j*step+step))
		}
		buf.WriteString("</tr>\n")
	}
	buf.WriteString("</thead></table>\n\n")
	// list
	buf.WriteString("| # | Title | Solution | Difficulty |\n")
	buf.WriteString("| :-: | - | - | :-: |\n")
	for i := maxId; i > 0; i-- {
		if row, ok := problemsSet[i]; ok {
			buf.WriteString(row)
		}
	}
}

var (
	buildCmd = "readme"
	fileName = "README.md"
)

var defaultStr = `
# [LeetCode](https://openset.github.io/leetcode)
LeetCode Problems' Solutions
[[话题分类](https://github.com/openset/leetcode/blob/master/tag/README.md)]

[![Build Status](https://travis-ci.org/openset/leetcode.svg?branch=master)](https://travis-ci.org/openset/leetcode)
[![codecov](https://codecov.io/gh/openset/leetcode/branch/master/graph/badge.svg)](https://codecov.io/gh/openset/leetcode)
[![Go Report Card](https://goreportcard.com/badge/github.com/openset/leetcode)](https://goreportcard.com/report/github.com/openset/leetcode)
[![GitHub contributors](https://img.shields.io/github/contributors/openset/leetcode.svg)](https://github.com/openset/leetcode/graphs/contributors)
[![license](https://img.shields.io/github/license/openset/leetcode.svg)](https://github.com/openset/leetcode/blob/master/LICENSE)
[![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/openset/leetcode.svg?colorB=green)](https://github.com/openset/leetcode/archive/master.zip)

`
