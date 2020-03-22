package explore

import (
	"fmt"
)

// Explore .
func Explore() {
	// byte == uint8
	// rune == int32

	s1 := "English"
	s2 := "文言文"

	fmt.Println([]byte(s1))
	fmt.Println([]rune(s1))

	fmt.Println([]byte(s2))
	fmt.Println([]rune(s2))

	fmt.Println(s1[:2])
	fmt.Println(string([]rune(s2)[:2]))

	// strings.Compare(a, b)
	// strings.Contains(s, substr)
	// strings.ContainsAny(s, chars)
	// strings.ContainsRune(s, r)
	// strings.Count(s, substr)
	// strings.EqualFold(s, t)
	// strings.Fields(s)
	// strings.FieldsFunc(s, f)
	// strings.HasPrefix(s, prefix)
	// strings.HasSuffix(s, suffix)
	// strings.Index(s, substr)
	// strings.IndexAny(s, chars)
	// strings.IndexByte(s, c)
	// strings.IndexFunc(s, f)
	// strings.IndexRune(s, r)
	// strings.Join(a [], sep)
	// strings.LastIndex(s, substr)
	// strings.LastIndexAny(s, chars)
	// strings.LastIndexByte(s, c)
	// strings.LastIndexFunc(s, f)
	// strings.Map(mapping, s)
	// strings.Repeat(s, count)
	// strings.Replace(s, old, new, n)
	// strings.ReplaceAll(s, old, new)
	// strings.Split(s, sep)
	// strings.SplitAfter(s, sep)
	// strings.SplitAfterN(s, sep, n)
	// strings.SplitN(s, sep, n)
	// strings.Title(s)
	// strings.ToLower(s)
	// strings.ToLowerSpecial(c, s)
	// strings.ToTitle(s)
	// strings.ToTitleSpecial(c, s)
	// strings.ToUpper(s)
	// strings.ToUpperSpecial(c, s)
	// strings.ToValidUTF8(s, replacement)
	// strings.Trim(s, cutset)
	// strings.TrimFunc(s, f)
	// strings.TrimLeft(s, cutset)
	// strings.TrimLeftFunc(s, f)
	// strings.TrimPrefix(s, prefix)
	// strings.TrimRight(s, cutset)
	// strings.TrimRightFunc(s, f)
	// strings.TrimSpace(s)
	// strings.TrimSuffix(s, suffix)

	// var b = strings.Builder
	// b.Cap()
	// b.Grow(n)
	// b.Len()
	// b.Reset()
	// b.String()
	// b.Write(p)
	// b.WriteByte(c)
	// b.WriteRune(r)
	// b.WriteString(s)

	// var reader = strings.Reader
	// NewReader(s string) *Reader
	// reader.Len()
	// reader.Read(b)
	// reader.ReadAt(b, off)
	// reader.ReadByte()
	// reader.ReadRune()
	// reader.Reset(s)
	// reader.Seek(offset, whence)
	// reader.Size()
	// reader.UnreadByte()
	// reader.UnreadRune()
	// reader.WriteTo(w)

	// var replacer = strings.Replacer
	// NewReplacer(oldnew) *Replacer
	// replacer.Replace(s)
	// replacer.WriteString(w, s)
}
