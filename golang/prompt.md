We are collaborating on a Go program, codename `monogram`. The `monogram` tool
translates from `monogram` notation into XML, JSON and other formats. The 
notation is designed to represent program-like texts. However it is just a
notation and not a programming language, although it does have an opinionated
grammar. Consequently it has no built-in variables, no built-in operators and
even the reserved words are dynamically discovered during the parse.

The program has several phases 

- an initial ingestion phase in which the input is tokenised, which is done.
- a pass to discover and mark the identifiers that are used as keywords, which is done.
- a parsing of the tokens to form an internal AST, which is our current focus.
- walking the AST to generate output, which is done.

I have already prototyped the parser as a recursive descent parser with 
operator precedence in a programming language called Pop-11. Unfortunately 
this is not a programming language you know well.

We have already implemented the Node structure in Go. Here it is:

```go
package main

type Node struct {
	Name     string            // The name of the node
	Options  map[string]string // Attributes (name-value pairs)
	Children []*Node           // Child nodes
}

func parseTokensToNodes(tokens []*Token) []*Node {
	// Dummy implementation for now: creates a node array based on dummy tokens
	return []*Node{
		{
			Name:    "node1",
			Options: map[string]string{"key1": "value1"},
			Children: []*Node{
				{
					Name:     "child1",
					Options:  map[string]string{"attribute": "data"},
					Children: nil,
				},
			},
		},
		{
			Name:     "node2",
			Options:  map[string]string{"key2": "value2"},
			Children: nil,
		},
	}
}

func parseToASTArray(input string) []*Node {
	// Step 1: Tokenize the input
	tokens := tokenizeInput(input)

	// Step 2: Parse the tokens into nodes
	nodes := parseTokensToNodes(tokens)

	return nodes
}

func parseToAST(input string, src string) *Node {
	// Get the array of nodes
	nodes := parseToASTArray(input)

	// Wrap the array in a "unit" node
	unitNode := &Node{
		Name:     "unit",
		Options:  map[string]string{"src": src},
		Children: nodes,
	}

	return unitNode
}
```

The core function is `read_expr_prec`, which consumes the input (the list
of tokens) and returns an tree of nodes. `read_expr_prec` reads a primary
expression off the input and then enters a loop in which it checks the 
precedence of the next symbol.

Here is the code for `read_primary_expr` in Pop-11. Even though you do
not know Pop-11 I think it is simple enough for you to use as a guide.

```pop11
define read_primary_expr();
    lvars q = proglist;
    lvars item = readitem();
    lvars tokentype = classify_item( item, proglist );
    returnunless( tokentype )( 
        if item.isstring then 
            lvars qm = string_quote( q );
            consNode( "string", $(quote=qm, value=item), [] )
        else
            consNode( "number", $(value=item), [] )
        endif 
    );
    false -> q;
    if inside_form and tokentype == "breaker" and unglue_option then
        lvars reclassified_tokentype = classify_item( item, [^unglue_option] );
        if reclassified_tokentype == tt_id then
            reclassified_tokentype -> tokentype;
            unglue_option :: proglist -> proglist
        endif
    endif;
    if tokentype == tt_open then
		lvars (sep, seq) = read_expr_seq_to( item.is_open_bracket, semi_comma, true );
        lvars dname = delimiter_name( item );
        consNode( "delimited", $(kind=dname, separator=sep.sep_as_word), seq )
    elseif tokentype == tt_start then
        read_form_expr( item )
    elseif tokentype == tt_id then
        consNode( "identifier", $(name=item), [] )
    elseif tokentype == tt_force then
        mishap( 'Misplaced macro indicator (' >< macro_mark >< ')', [] )
    elseif tokentype == tt_macro then
        pop11_need_nextreaditem( macro_mark ) -> _;
        if item.isword then
            lvars e = read_opt_expr_prec(max_precedence, true);
            if e then
                consNode( "form", syntax_prefix, [% consNode( "part", $(keyword=item), [^e] ) %] )
            else
                consNode( "form", syntax_prefix, [% consNode( "part", $(keyword=item), [] ) %] )
            endif
        else
            mishap( 'Identifier required following `' >< macro_mark >< '`', [^item] )
        endif
    else
        lvars p = precedence( item );
        if p then
            lvars e = read_expr_prec( p, false );
            consNode( "operator", $(name=item, syntax="prefix"), [ ^e ] )
        else
            mishap( 'Unexpected token at start of expression', [^item] )
        endif
    endif
enddefine;
```