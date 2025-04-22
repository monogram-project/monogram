package mg

import (
	"fmt"
	"strconv"
)

type Iterator[T any] interface {
	Next() bool
	Value() T
}

type Element interface {
	Name() string
	FromTo() string
	GetOption(string) string
	SetOption(string, string)
	HasOption(string) bool
	Options() Iterator[string]
	Children() Iterator[Element]
	ChildrenCount() int
}

type NullIterator[T any] struct{}

func (it *NullIterator[T]) Next() bool {
	return false
}

func (it *NullIterator[T]) Value() T {
	var zero T
	return zero
}

func NewNullIterator[T any]() *NullIterator[T] {
	return &NullIterator[T]{}
}

type SingletonIterator[T any] struct {
	value T
	done  bool
}

func (it *SingletonIterator[T]) Next() bool {
	if it.done {
		return false
	}
	it.done = true
	return true
}

func (it *SingletonIterator[T]) Value() T {
	return it.value
}

func NewSingletonIterator[T any](value T) *SingletonIterator[T] {
	return &SingletonIterator[T]{value: value}
}

type SliceIterator[T any] struct {
	index int
	slice []T
}

func (it *SliceIterator[T]) Next() bool {
	it.index++
	return it.index < len(it.slice)
}

func (it *SliceIterator[T]) Value() T {
	return it.slice[it.index]
}

func NewSliceIterator[T any](slice []T) *SliceIterator[T] {
	return &SliceIterator[T]{slice: slice, index: -1}
}

func NewItemsIterator[T any](items ...T) *SliceIterator[T] {
	return &SliceIterator[T]{slice: items, index: -1}
}

//-- Interpolate ---------------------------------------------------------------

type InterpolateNode struct {
	Kind    string  `json:"kind"`
	Span    string  `json:"span"`
	Element Element `json:"element"`
}

func (n *InterpolateNode) Name() string {
	return NameInterpolate
}

func (n *InterpolateNode) GetOption(name string) string {
	if name == OptionKind {
		return n.Kind
	} else if name == OptionSpan {
		return n.Span
	}
	return ""
}

func (n *InterpolateNode) SetOption(name string, value string) {
	if name == OptionKind {
		n.Kind = value
	} else if name == OptionSpan {
		n.Span = value
	}
}

func (n *InterpolateNode) HasOption(name string) bool {
	return name == OptionKind || name == OptionSpan
}

func (n *InterpolateNode) Options() Iterator[string] {
	return NewItemsIterator(OptionKind, OptionSpan)
}

func (n *InterpolateNode) Children() Iterator[Element] {
	return NewSingletonIterator(n.Element)
}

func (n *InterpolateNode) ChildrenCount() int {
	return 1
}

func (n *InterpolateNode) FromTo() string {
	return n.Span
}

// -- JoinLines ----------------------------------------------------------------

type JoinLinesNode struct {
	Quote    string    `json:"quote"`
	Span     string    `json:"span"`
	Elements []Element `json:"elements"`
}

func (n *JoinLinesNode) Name() string {
	return NameJoin
}

func (n *JoinLinesNode) GetOption(name string) string {
	if name == OptionQuote {
		return n.Quote
	} else if name == OptionSpan {
		return n.Span
	}
	return ""
}

func (n *JoinLinesNode) SetOption(name string, value string) {
	if name == OptionQuote {
		n.Quote = value
	} else if name == OptionSpan {
		n.Span = value
	}
}

func (n *JoinLinesNode) HasOption(name string) bool {
	return name == OptionQuote || name == OptionSpan
}

func (n *JoinLinesNode) Options() Iterator[string] {
	return NewItemsIterator(OptionQuote, OptionSpan)
}

func (n *JoinLinesNode) Children() Iterator[Element] {
	return NewSliceIterator(n.Elements)
}

func (n *JoinLinesNode) ChildrenCount() int {
	return len(n.Elements)
}

func (n *JoinLinesNode) FromTo() string {
	return n.Span
}

//-- Join ----------------------------------------------------------------------

type JoinNode struct {
	Quote    string    `json:"quote"`
	Span     string    `json:"span"`
	Elements []Element `json:"elements"`
}

func (n *JoinNode) Name() string {
	return NameJoin
}

func (n *JoinNode) GetOption(name string) string {
	if name == OptionQuote {
		return n.Quote
	} else if name == OptionSpan {
		return n.Span
	}
	return ""
}

func (n *JoinNode) SetOption(name string, value string) {
	if name == OptionQuote {
		n.Quote = value
	} else if name == OptionSpan {
		n.Span = value
	}
}

func (n *JoinNode) HasOption(name string) bool {
	return name == OptionQuote || name == OptionSpan
}

func (n *JoinNode) Options() Iterator[string] {
	return NewItemsIterator(OptionQuote, OptionSpan)
}

func (n *JoinNode) Children() Iterator[Element] {
	return NewSliceIterator(n.Elements)
}

func (n *JoinNode) ChildrenCount() int {
	return len(n.Elements)
}

func (n *JoinNode) FromTo() string {
	return n.Span
}

//-- Form ----------------------------------------------------------------------

type FormNode struct {
	Syntax string    `json:"syntax"`
	Span   string    `json:"span"`
	Parts  []Element `json:"parts"`
}

func (n *FormNode) Name() string {
	return NameForm
}

func (n *FormNode) GetOption(name string) string {
	if name == OptionSyntax {
		return n.Syntax
	} else if name == OptionSpan {
		return n.Span
	}
	return ""
}

func (n *FormNode) SetOption(name string, value string) {
	if name == OptionSyntax {
		n.Syntax = value
	} else if name == OptionSpan {
		n.Span = value
	}
}

func (n *FormNode) HasOption(name string) bool {
	return name == OptionSyntax || name == OptionSpan
}

func (n *FormNode) Options() Iterator[string] {
	return NewItemsIterator(OptionSyntax, OptionSpan)
}

func (n *FormNode) Children() Iterator[Element] {
	return NewSliceIterator(n.Parts)
}

func (n *FormNode) ChildrenCount() int {
	return len(n.Parts)
}

func (n *FormNode) FromTo() string {
	return n.Span
}

// -- Part ---------------------------------------------------------------------

type PartNode struct {
	Keyword string    `json:"keyword"`
	Span    string    `json:"span"`
	Exprs   []Element `json:"children"`
}

func (n *PartNode) Name() string {
	return NamePart
}

func (n *PartNode) GetOption(name string) string {
	if name == OptionKeyword {
		return n.Keyword
	} else if name == OptionSpan {
		return n.Span
	}
	return ""
}

func (n *PartNode) SetOption(name string, value string) {
	if name == OptionKeyword {
		n.Keyword = value
	} else if name == OptionSpan {
		n.Span = value
	}
}

func (n *PartNode) HasOption(name string) bool {
	return name == OptionKeyword || name == OptionSpan
}

func (n *PartNode) Options() Iterator[string] {
	return NewItemsIterator(OptionKeyword, OptionSpan)
}

func (n *PartNode) Children() Iterator[Element] {
	return NewSliceIterator(n.Exprs)
}

func (n *PartNode) ChildrenCount() int {
	return len(n.Exprs)
}

func (n *PartNode) FromTo() string {
	return n.Span
}

// -- Unit ---------------------------------------------------------------------

type UnitNode struct {
	Src   string    `json:"keyword"`
	Span  string    `json:"span"`
	Exprs []Element `json:"children"`
}

func (n *UnitNode) Name() string {
	return NameUnit
}

func (n *UnitNode) GetOption(name string) string {
	if name == OptionSrc {
		return n.Src
	} else if name == OptionSpan {
		return n.Span
	}
	return ""
}

func (n *UnitNode) SetOption(name string, value string) {
	if name == OptionSrc {
		n.Src = value
	} else if name == OptionSpan {
		n.Span = value
	}
}

func (n *UnitNode) HasOption(name string) bool {
	return name == OptionSrc || name == OptionSpan
}

func (n *UnitNode) Options() Iterator[string] {
	return NewItemsIterator(OptionSrc, OptionSpan)
}

func (n *UnitNode) Children() Iterator[Element] {
	return NewSliceIterator(n.Exprs)
}

func (n *UnitNode) ChildrenCount() int {
	return len(n.Exprs)
}

func (n *UnitNode) FromTo() string {
	return n.Span
}

// -- ApplyNode ----------------------------------------------------------------

type ApplyNode struct {
	Kind      string  `json:"kind"`
	Separator string  `json:"separator"`
	Span      string  `json:"span"`
	Func      Element `json:"func"`
	Args      Element `json:"args"`
}

func (n *ApplyNode) Name() string {
	return NameApply
}

func (n *ApplyNode) GetOption(name string) string {
	if name == OptionKind {
		return n.Kind
	} else if name == OptionSeparator {
		return n.Separator
	} else if name == OptionSpan {
		return n.Span
	}
	return ""
}

func (n *ApplyNode) SetOption(name string, value string) {
	if name == OptionKind {
		n.Kind = value
	} else if name == OptionSeparator {
		n.Separator = value
	} else if name == OptionSpan {
		n.Span = value
	}
}

func (n *ApplyNode) HasOption(name string) bool {
	return name == OptionKind || name == OptionSeparator || name == OptionSpan
}

func (n *ApplyNode) Options() Iterator[string] {
	return NewItemsIterator(OptionKind, OptionSeparator, OptionSpan)
}

func (n *ApplyNode) Children() Iterator[Element] {
	return NewItemsIterator(n.Func, n.Args)
}

func (n *ApplyNode) ChildrenCount() int {
	return 2
}

func (n *ApplyNode) FromTo() string {
	return n.Span
}

// -- ArgumentsNode ------------------------------------------------------------

type ArgumentsNode struct {
	Span  string    `json:"span"`
	Exprs []Element `json:"children"`
}

func (n *ArgumentsNode) Name() string {
	return NameArguments
}

func (n *ArgumentsNode) GetOption(name string) string {
	if name == OptionSpan {
		return n.Span
	}
	return ""
}

func (n *ArgumentsNode) SetOption(name string, value string) {
	if name == OptionSpan {
		n.Span = value
	}
}

func (n *ArgumentsNode) HasOption(name string) bool {
	return name == OptionSpan
}

func (n *ArgumentsNode) Options() Iterator[string] {
	return NewSingletonIterator(OptionSpan)
}

func (n *ArgumentsNode) Children() Iterator[Element] {
	return NewSliceIterator(n.Exprs)
}

func (n *ArgumentsNode) ChildrenCount() int {
	return len(n.Exprs)
}

func (n *ArgumentsNode) FromTo() string {
	return n.Span
}

// -- DelimitedNode ------------------------------------------------------------

type DelimitedNode struct {
	Kind      string    `json:"kind"`
	Separator string    `json:"separator"`
	Span      string    `json:"span"`
	Exprs     []Element `json:"children"`
}

func (n *DelimitedNode) Name() string {
	return NameDelimited
}

func (n *DelimitedNode) GetOption(name string) string {
	if name == OptionKind {
		return n.Kind
	} else if name == OptionSeparator {
		return n.Separator
	} else if name == OptionSpan {
		return n.Span
	}
	return ""
}

func (n *DelimitedNode) SetOption(name string, value string) {
	if name == OptionKind {
		n.Kind = value
	} else if name == OptionSeparator {
		n.Separator = value
	} else if name == OptionSpan {
		n.Span = value
	}
}

func (n *DelimitedNode) HasOption(name string) bool {
	return name == OptionKind || name == OptionSeparator || name == OptionSpan
}

func (n *DelimitedNode) Options() Iterator[string] {
	return NewItemsIterator(OptionKind, OptionSeparator, OptionSpan)
}

func (n *DelimitedNode) Children() Iterator[Element] {
	return NewSliceIterator(n.Exprs)
}

func (n *DelimitedNode) ChildrenCount() int {
	return len(n.Exprs)
}

func (n *DelimitedNode) FromTo() string {
	return n.Span
}

// -- GetNode ------------------------------------------------------------------

type GetNode struct {
	Property string `json:"name"`
	Span     string `json:"span"`
	Expr     Element
}

func (n *GetNode) Name() string {
	return n.Property
}

func (n *GetNode) GetOption(name string) string {
	if name == OptionName {
		return n.Property
	} else if name == OptionSpan {
		return n.Span
	}
	return ""
}

func (n *GetNode) SetOption(name string, value string) {
	if name == OptionName {
		n.Property = value
	} else if name == OptionSpan {
		n.Span = value
	}
}

func (n *GetNode) HasOption(name string) bool {
	return name == OptionName || name == OptionSpan
}

func (n *GetNode) Options() Iterator[string] {
	return NewItemsIterator(OptionName, OptionSpan)
}

func (n *GetNode) Children() Iterator[Element] {
	return NewSingletonIterator(n.Expr)
}

func (n *GetNode) ChildrenCount() int {
	return 1
}

func (n *GetNode) FromTo() string {
	return n.Span
}

// -- IdentifierNode -----------------------------------------------------------

type IdentifierNode struct {
	Id   string `json:"name"`
	Span string `json:"span"`
}

func (n *IdentifierNode) Name() string {
	return NameIdentifier
}

func (n *IdentifierNode) GetOption(name string) string {
	if name == OptionName {
		return n.Id
	} else if name == OptionSpan {
		return n.Span
	}
	return ""
}

func (n *IdentifierNode) SetOption(name string, value string) {
	if name == OptionName {
		n.Id = value
	} else if name == OptionSpan {
		n.Span = value
	}
}

func (n *IdentifierNode) HasOption(name string) bool {
	return name == OptionName || name == OptionSpan
}

func (n *IdentifierNode) Options() Iterator[string] {
	return NewItemsIterator(OptionName, OptionSpan)
}

func (n *IdentifierNode) Children() Iterator[Element] {
	return NewNullIterator[Element]()
}

func (n *IdentifierNode) ChildrenCount() int {
	return 0
}

func (n *IdentifierNode) FromTo() string {
	return n.Span
}

// -- InvokeNode ---------------------------------------------------------------

type InvokeNode struct {
	Kind      string  `json:"kind"`
	Separator string  `json:"separator"`
	Property  string  `json:"name"`
	Span      string  `json:"span"`
	LHS       Element `json:"lhs"`
	RHS       Element `json:"rhs"`
}

func (n *InvokeNode) Name() string {
	return NameInvoke
}

func (n *InvokeNode) GetOption(name string) string {
	if name == OptionKind {
		return n.Kind
	} else if name == OptionSeparator {
		return n.Separator
	} else if name == OptionName {
		return n.Property
	} else if name == OptionSpan {
		return n.Span
	}
	return ""
}

func (n *InvokeNode) SetOption(name string, value string) {
	if name == OptionKind {
		n.Kind = value
	} else if name == OptionSeparator {
		n.Separator = value
	} else if name == OptionName {
		n.Property = value
	} else if name == OptionSpan {
		n.Span = value
	}
}

func (n *InvokeNode) HasOption(name string) bool {
	return name == OptionKind || name == OptionSeparator || name == OptionName || name == OptionSpan
}

func (n *InvokeNode) Options() Iterator[string] {
	return NewItemsIterator(OptionKind, OptionSeparator, OptionName, OptionSpan)
}

func (n *InvokeNode) Children() Iterator[Element] {
	return NewItemsIterator(n.LHS, n.RHS)
}

func (n *InvokeNode) ChildrenCount() int {
	return 2
}

func (n *InvokeNode) FromTo() string {
	return n.Span
}

// -- NumberNode ---------------------------------------------------------------

type NumberNode struct {
	Value float64 `json:"value"`
	Span  string  `json:"span"`
}

func (n *NumberNode) Name() string {
	return NameNumber
}

func (n *NumberNode) GetOption(name string) string {
	if name == OptionValue {
		return fmt.Sprintf("%f", n.Value)
	} else if name == OptionSpan {
		return n.Span
	}
	return ""
}

func (n *NumberNode) SetOption(name string, value string) {
	if name == OptionValue {
		fmt.Sscanf(value, "%f", &n.Value)
	} else if name == OptionSpan {
		n.Span = value
	}
}

func (n *NumberNode) HasOption(name string) bool {
	return name == OptionValue || name == OptionSpan
}

func (n *NumberNode) Options() Iterator[string] {
	return NewItemsIterator(OptionValue, OptionSpan)
}

func (n *NumberNode) Children() Iterator[Element] {
	return NewNullIterator[Element]()
}

func (n *NumberNode) ChildrenCount() int {
	return 0
}

func (n *NumberNode) FromTo() string {
	return n.Span
}

// -- OperatorNode -------------------------------------------------------------

type InfixOperatorNode struct {
	Op   string  `json:"name"`
	Span string  `json:"span"`
	LHS  Element `json:"lhs"`
	RHS  Element `json:"rhs"`
}

func (n *InfixOperatorNode) Name() string {
	return NameOperator
}

func (n *InfixOperatorNode) GetOption(name string) string {
	if name == OptionName {
		return n.Op
	} else if name == OptionSyntax {
		return ValueInfix
	} else if name == OptionSpan {
		return n.Span
	}
	return ""
}

func (n *InfixOperatorNode) SetOption(name string, value string) {
	if name == OptionName {
		n.Op = value
	} else if name == OptionSpan {
		n.Span = value
	}
}

func (n *InfixOperatorNode) HasOption(name string) bool {
	return name == OptionName || name == OptionSyntax || name == OptionSpan
}

func (n *InfixOperatorNode) Options() Iterator[string] {
	return NewItemsIterator(OptionName, OptionSyntax, OptionSpan)
}

func (n *InfixOperatorNode) Children() Iterator[Element] {
	return NewItemsIterator(n.LHS, n.RHS)
}

func (n *InfixOperatorNode) ChildrenCount() int {
	return 2
}

func (n *InfixOperatorNode) FromTo() string {
	return n.Span
}

type PrefixOperatorNode struct {
	Op   string  `json:"name"`
	Span string  `json:"span"`
	Arg  Element `json:"arg"`
}

func (n *PrefixOperatorNode) Name() string {
	return NameOperator
}

func (n *PrefixOperatorNode) GetOption(name string) string {
	if name == OptionName {
		return n.Op
	} else if name == OptionSyntax {
		return ValuePrefix
	} else if name == OptionSpan {
		return n.Span
	}
	return ""
}

func (n *PrefixOperatorNode) SetOption(name string, value string) {
	if name == OptionName {
		n.Op = value
	} else if name == OptionSpan {
		n.Span = value
	}
}

func (n *PrefixOperatorNode) HasOption(name string) bool {
	return name == OptionName || name == OptionSyntax || name == OptionSpan
}

func (n *PrefixOperatorNode) Options() Iterator[string] {
	return NewItemsIterator(OptionName, OptionSyntax, OptionSpan)
}

func (n *PrefixOperatorNode) Children() Iterator[Element] {
	return NewSingletonIterator(n.Arg)
}

func (n *PrefixOperatorNode) ChildrenCount() int {
	return 1
}

func (n *PrefixOperatorNode) FromTo() string {
	return n.Span
}

// -- StringNode ---------------------------------------------------------------

type StringNode struct {
	Value string `json:"value"`
	Quote string `json:"quote"`
	Span  string `json:"span"`
}

func (n *StringNode) Name() string {
	return NameString
}

func (n *StringNode) GetOption(name string) string {
	if name == OptionValue {
		return n.Value
	} else if name == OptionSyntax {
		return n.Quote
	} else if name == OptionSpan {
		return n.Span
	}
	return ""
}

func (n *StringNode) SetOption(name string, value string) {
	if name == OptionValue {
		n.Value = value
	} else if name == OptionSyntax {
		n.Quote = value
	} else if name == OptionSpan {
		n.Span = value
	}
}

func (n *StringNode) HasOption(name string) bool {
	return name == OptionValue || name == OptionSyntax || name == OptionSpan
}

func (n *StringNode) Options() Iterator[string] {
	return NewItemsIterator(OptionValue, OptionSyntax, OptionSpan)
}

func (n *StringNode) Children() Iterator[Element] {
	return NewNullIterator[Element]()
}

func (n *StringNode) ChildrenCount() int {
	return 0
}

func (n *StringNode) FromTo() string {
	return n.Span
}

// -----------------------------------------------------------------------------
// Convert general *Node to Element
// -----------------------------------------------------------------------------

func AllToElement(nodes []*Node) ([]Element, error) {
	var elements []Element
	for _, node := range nodes {
		e, err := node.ToElement()
		if err != nil {
			return nil, err
		}
		elements = append(elements, e)
	}
	return elements, nil
}

func (node *Node) ToElement() (Element, error) {
	switch node.Name {
	case NameForm:
		kids, err := AllToElement(node.Children)
		if err != nil {
			return nil, err
		}
		return &FormNode{
			Syntax: node.Options[OptionSyntax],
			Span:   node.Options[OptionSpan],
			Parts:  kids,
		}, nil
	case NamePart:
		kids, err := AllToElement(node.Children)
		if err != nil {
			return nil, err
		}
		return &PartNode{
			Keyword: node.Options[OptionKeyword],
			Span:    node.Options[OptionSpan],
			Exprs:   kids,
		}, nil
	case NameUnit:
		kids, err := AllToElement(node.Children)
		if err != nil {
			return nil, err
		}
		return &UnitNode{
			Src:   node.Options[OptionSrc],
			Span:  node.Options[OptionSpan],
			Exprs: kids,
		}, nil
	case NameApply:
		if len(node.Children) != 2 {
			return nil, fmt.Errorf("apply node must have exactly two children")
		}
		e0, err := node.Children[0].ToElement()
		if err != nil {
			return nil, err
		}
		e1, err := node.Children[1].ToElement()
		if err != nil {
			return nil, err
		}
		return &ApplyNode{
			Kind:      node.Options[OptionKind],
			Separator: node.Options[OptionSeparator],
			Span:      node.Options[OptionSpan],
			Func:      e0,
			Args:      e1,
		}, nil
	case NameArguments:
		kids, err := AllToElement(node.Children)
		if err != nil {
			return nil, err
		}
		return &ArgumentsNode{
			Span:  node.Options[OptionSpan],
			Exprs: kids,
		}, nil
	case NameDelimited:
		kids, err := AllToElement(node.Children)
		if err != nil {
			return nil, err
		}
		return &DelimitedNode{
			Kind:      node.Options[OptionKind],
			Separator: node.Options[OptionSeparator],
			Span:      node.Options[OptionSpan],
			Exprs:     kids,
		}, nil
	case NameGet:
		if len(node.Children) != 1 {
			return nil, fmt.Errorf("get node must have exactly one child")
		}
		e, err := node.Children[0].ToElement()
		if err != nil {
			return nil, err
		}
		return &GetNode{
			Property: node.Options[OptionName],
			Span:     node.Options[OptionSpan],
			Expr:     e,
		}, nil
	case NameIdentifier:
		return &IdentifierNode{
			Id:   node.Options[OptionName],
			Span: node.Options[OptionSpan],
		}, nil
	case NameInvoke:
		if len(node.Children) != 2 {
			return nil, fmt.Errorf("invoke node must have exactly two children")
		}
		e0, err := node.Children[0].ToElement()
		if err != nil {
			return nil, err
		}
		e1, err := node.Children[1].ToElement()
		if err != nil {
			return nil, err
		}
		return &InvokeNode{
			Kind:      node.Options[OptionKind],
			Separator: node.Options[OptionSeparator],
			Property:  node.Options[OptionName],
			Span:      node.Options[OptionSpan],
			LHS:       e0,
			RHS:       e1,
		}, nil
	case NameNumber:
		value, err := strconv.ParseFloat(node.Options[OptionValue], 64)
		if err != nil {
			value = 0.0
		}
		return &NumberNode{
			Value: value,
			Span:  node.Options[OptionSpan],
		}, nil
	case NameOperator:
		if len(node.Children) < 1 {
			return nil, fmt.Errorf("operator node must have at least one child")
		}
		e0, err := node.Children[0].ToElement()
		if err != nil {
			return nil, err
		}
		if node.Options[OptionSyntax] == ValueInfix {
			if len(node.Children) != 2 {
				return nil, fmt.Errorf("infix operator node must have exactly two children")
			}
			e1, err := node.Children[1].ToElement()
			if err != nil {
				return nil, err
			}
			return &InfixOperatorNode{
				Op:   node.Options[OptionName],
				Span: node.Options[OptionSpan],
				LHS:  e0,
				RHS:  e1,
			}, nil
		}
		if len(node.Children) != 1 {
			return nil, fmt.Errorf("prefix operator node must have exactly one child")
		}
		return &PrefixOperatorNode{
			Op:   node.Options[OptionName],
			Span: node.Options[OptionSpan],
			Arg:  e0,
		}, nil
	case NameString:
		return &StringNode{
			Value: node.Options[OptionValue],
			Quote: node.Options[OptionSyntax],
			Span:  node.Options[OptionSpan],
		}, nil
	case NameJoin:
		kids, err := AllToElement(node.Children)
		if err != nil {
			return nil, err
		}
		return &JoinNode{
			Quote:    node.Options[OptionQuote],
			Span:     node.Options[OptionSpan],
			Elements: kids,
		}, nil
	case NameJoinLines:
		kids, err := AllToElement(node.Children)
		if err != nil {
			return nil, err
		}
		return &JoinLinesNode{
			Quote:    node.Options[OptionQuote],
			Span:     node.Options[OptionSpan],
			Elements: kids,
		}, nil
	case NameInterpolate:
		e, err := node.Children[0].ToElement()
		if err != nil {
			return nil, err
		}
		return &InterpolateNode{
			Kind:    node.Options[OptionKind],
			Span:    node.Options[OptionSpan],
			Element: e,
		}, nil
	default:
		return nil, fmt.Errorf("unknown node type: %s", node.Name)
	}
}
