package mg

type Node struct {
	Name     string            // The name of the node
	Options  map[string]string // Attributes (name-value pairs)
	Children []*Node           // Child nodes
}

const NameForm = "form"
const NamePart = "part"
const NameUnit = "unit"
const NameApply = "apply"
const NameArguments = "arguments"
const NameDelimited = "delimited"
const NameGet = "get"
const NameIdentifier = "identifier"
const NameInvoke = "invoke"
const NameNumber = "number"
const NameOperator = "operator"
const NameString = "string"
const NameJoin = "join"
const NameJoinLines = "joinlines"
const NameInterpolate = "interpolation"

const OptionValue = "value"
const OptionsDecimalValue = "decimal"
const OptionName = "name"
const OptionKind = "kind"
const OptionSeparator = "separator"
const OptionKeyword = "keyword"
const OptionSpan = "span"
const OptionSyntax = "syntax"
const OptionQuote = "quote"
const OptionSrc = "src"

const ValueInfix = "infix"
const ValuePrefix = "prefix"
const ValueSurround = "surround"
const ValueComma = "comma"
const ValueSemicolon = "semicolon"
const ValueUndefined = "undefined"

type FormBuilder struct {
	node         *Node
	startForm    LineCol // The span of the form
	startPart    LineCol // The span of the current part
	includeSpans bool    // Whether to include spans in the form
}

func NewFormBuilder(partName string, lc LineCol, includeSpans bool) *FormBuilder {
	return &FormBuilder{
		node: &Node{
			Name: NameForm,
			Options: map[string]string{
				OptionSyntax: ValuePrefix,
			},
			Children: []*Node{
				{
					Name: NamePart,
					Options: map[string]string{
						OptionKeyword: partName,
					},
					Children: []*Node{},
				},
			},
		},
		startForm:    lc,
		startPart:    lc,
		includeSpans: includeSpans,
	}
}

func (b *FormBuilder) AddChild(child *Node) {
	parts := b.node.Children
	lastpart := parts[len(parts)-1]
	lastpart.Children = append(lastpart.Children, child)
}

func (b *FormBuilder) endPartSpan(endPart LineCol) {
	if b.includeSpans {
		parts := b.node.Children
		lastpart := parts[len(parts)-1]
		lastpart.Options[OptionSpan] = b.startPart.SpanString(endPart)
	}
}

func (b *FormBuilder) BeginNextPart(partName string, endOldPart LineCol, startNewPart LineCol) {

	// Set the span of the last part
	b.endPartSpan(endOldPart)

	// Create a new part
	b.node.Children = append(b.node.Children, &Node{
		Name: NamePart,
		Options: map[string]string{
			OptionKeyword: partName,
		},
		Children: []*Node{},
	})

	// Set the start of the new part
	b.startPart = startNewPart

}

func (b *FormBuilder) Build(endForm LineCol) *Node {
	if b.includeSpans {
		b.endPartSpan(endForm)
		b.node.Options[OptionSpan] = b.startForm.SpanString(endForm)
	}
	return b.node
}
