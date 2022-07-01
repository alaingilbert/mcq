package mc

import (
	"encoding/json"
	"github.com/alaingilbert/mcq/nbt"
	"strings"
)

type Sign struct {
	blockEntity
	text1       string
	text2       string
	text3       string
	text4       string
	glowingText bool
}

func SignFromNbt(node *nbt.TagNodeCompound) *Sign {
	s := new(Sign)
	s.blockEntity = *BlockEntityFromNbt(node)
	glowingTextRaw := node.Entries["GlowingText"]
	if glowingTextRaw != nil {
		s.glowingText = glowingTextRaw.(*nbt.TagNodeByte).Byte() == 1
	}
	type txt struct {
		Text string `json:"text"`
	}
	var t1, t2, t3, t4 txt
	_ = json.Unmarshal([]byte(node.Entries["Text1"].(*nbt.TagNodeString).String()), &t1)
	_ = json.Unmarshal([]byte(node.Entries["Text2"].(*nbt.TagNodeString).String()), &t2)
	_ = json.Unmarshal([]byte(node.Entries["Text3"].(*nbt.TagNodeString).String()), &t3)
	_ = json.Unmarshal([]byte(node.Entries["Text4"].(*nbt.TagNodeString).String()), &t4)
	s.text1 = t1.Text
	s.text2 = t2.Text
	s.text3 = t3.Text
	s.text4 = t4.Text
	return s
}

func (s Sign) InlineText() (out string) {
	arr := make([]string, 0)
	if s.text1 != "" {
		arr = append(arr, s.text1)
	}
	if s.text2 != "" {
		arr = append(arr, s.text2)
	}
	if s.text3 != "" {
		arr = append(arr, s.text3)
	}
	if s.text4 != "" {
		arr = append(arr, s.text4)
	}
	return strings.Join(arr, " ")
}

func (s Sign) HasText() bool {
	return s.text1 != "" || s.text2 != "" || s.text3 != "" || s.text4 != ""
}

func (s Sign) GlowingText() bool { return s.glowingText }
func (s Sign) Text1() string     { return s.text1 }
func (s Sign) Text2() string     { return s.text2 }
func (s Sign) Text3() string     { return s.text3 }
func (s Sign) Text4() string     { return s.text4 }
