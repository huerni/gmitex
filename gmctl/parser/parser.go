package parser

import (
	"errors"
	"github.com/emicklei/proto"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrGoPackage = errors.New(`option go_package = "" field is not filled in`)

type ProtoParser struct {
}

func NewProtoParser() *ProtoParser {
	return &ProtoParser{}
}

func (p *ProtoParser) Parse(src string) (Proto, error) {
	var ret Proto

	abs, err := filepath.Abs(src)
	if err != nil {
		return Proto{}, err
	}

	r, err := os.Open(abs)
	if err != nil {
		return ret, err
	}
	defer r.Close()

	parser := proto.NewParser(r)
	definition, err := parser.Parse()
	if err != nil {
		return ret, err
	}

	var serviceList []Service

	proto.Walk(definition,
		proto.WithImport(func(i *proto.Import) {
			ret.Import = append(ret.Import, Import{Import: i})
		}),
		proto.WithPackage(func(p *proto.Package) {
			ret.Package = Package{Package: p}
		}),
		proto.WithMessage(func(message *proto.Message) {
			ret.Message = append(ret.Message, Message{message})
		}),
		proto.WithService(func(service *proto.Service) {
			serv := Service{Service: service}
			elements := service.Elements
			for _, el := range elements {
				v, _ := el.(*proto.RPC)
				if v == nil {
					continue
				}
				serv.RPC = append(serv.RPC, &RPC{RPC: v})
			}
			serviceList = append(serviceList, serv)
		}),
		proto.WithOption(func(option *proto.Option) {
			if option.Name == "go_package" {
				ret.GoPackage = option.Constant.Source
			}
		}),
	)

	if len(ret.GoPackage) == 0 {
		if ret.Package.Package == nil {
			return ret, ErrGoPackage
		}
		ret.GoPackage = ret.Package.Name
	}

	ret.PbPackage = GoSanitized(filepath.Base(ret.GoPackage))
	ret.Src = abs
	ret.Name = filepath.Base(abs)
	ret.Service = serviceList

	return ret, nil
}

func GoSanitized(s string) string {
	s = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return r
		}
		return '_'
	}, s)

	r, _ := utf8.DecodeRuneInString(s)
	if token.Lookup(s).IsKeyword() || !unicode.IsLetter(r) {
		return "_" + s
	}
	return s
}

func GetComment(comment *proto.Comment) string {
	if comment == nil {
		return ""
	}
	return "// " + strings.TrimSpace(comment.Message())
}

func CamelCase(s string) string {
	if s == "" {
		return ""
	}
	t := make([]byte, 0, 32)
	i := 0
	if s[0] == '_' {
		// Need a capital letter; drop the '_'.
		t = append(t, 'X')
		i++
	}
	// Invariant: if the next letter is lower case, it must be converted
	// to upper case.
	// That is, we process a word at a time, where words are marked by _ or
	// upper case letter. Digits are treated as words.
	for ; i < len(s); i++ {
		c := s[i]
		if c == '_' && i+1 < len(s) && isASCIILower(s[i+1]) {
			continue // Skip the underscore in s.
		}
		if isASCIIDigit(c) {
			t = append(t, c)
			continue
		}
		// Assume we have a letter now - if not, it's a bogus identifier.
		// The next word is a sequence of characters that must start upper case.
		if isASCIILower(c) {
			c ^= ' ' // Make it a capital letter.
		}
		t = append(t, c) // Guaranteed not lower case.
		// Accept lower case sequence that follows.
		for i+1 < len(s) && isASCIILower(s[i+1]) {
			i++
			t = append(t, s[i])
		}
	}
	return string(t)
}

func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

// Is c an ASCII digit?
func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
