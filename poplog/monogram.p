;;; A prototype of monogram in Pop-11.

compile_mode :pop11 +strict;

;;; Tokenisation:                               Has precedence?     Classify
;;;     0. Comments - # ...                     N/A                 N/A
;;;     1. Literals - strings, numbers          false               false
;;;     2. Separators - comma, semi-colon       false               "sep"
;;;        [a] Expr separator - comma
;;;        [b] Stmnt separator - semi-colon
;;;     3. Open Brackets - ( { [                true                "open"
;;;     4. Close Brackets - ] } )               false               "close"
;;;     5. Start Form - XXX                     false               "start"
;;;     6. End Form - endXXX                    false               "end"
;;;     7. Force - !                            false           1   "force"
;;;     8. Form breakers - foo:                 false               "breaker"
;;;     9. Form breakers - else-if              false               "breaker1"
;;;    10. Label - : ?                          false               "label"
;;;    11. Signs - + / %                        true                "sign"
;;;        [a] Method invoke - dot
;;;        [b] The rest
;;;        [c] N.B. Note the special role of Force and Labels.
;;;    12. Identifiers                          false               "id"

vars last_char_read = false;
vars last_nonwhitespace_read = false;
vars unglue_option = false;
vars allow_newline_option = false;
vars inferred_form_starts = [];
vars inferred_forced_words = [];
vars inside_form = false;

constant semi = ";";
constant comma = ",";
constant semi_comma = [ ^semi ^comma ];
constant colon = ":";
constant macro_mark = "!";


define is_open_bracket( word );
    returnif( word == "(" )( ")" );
    returnif( word == "{" )( "}" );
    returnif( word == "[" )( "]" );
    return( false )
enddefine;

define is_close_bracket( word );
    returnif( word == ")" )( true );
    returnif( word == "}" )( true );
    returnif( word == "]" )( true );
    return( false )
enddefine;

define delimiter_name( word );
    returnif( word == "(" )( "parentheses" );
    returnif( word == "{" )( "braces" );
    returnif( word == "[" )( "brackets" );
    return( false )
enddefine;

define is_sign( word );
    lvars ch;
    for ch in_vectorclass word.fast_word_string do
        nextif( ch == `.` );
        lvars n = item_chartype(ch);
        nextif( n == 3 or n == 10 or n == 11 );
        return( false );
    endfor;
    return( word /== macro_mark and not( inside_form and word == colon ) );
enddefine;

define is_form_end( word );
    isstartstring( 'end', word )
enddefine;

vars procedure whitespace_after_item, string_quote;

define classify_item( item, subsequent_items );
    returnunless( item.isword )( false );
    lvars L = datalength( item );
    returnif( L == 0 )( false );
    if L == 1 then
        lvars ch_first = subscrw( 1, item );
        returnif( ch_first == subscrw(1, macro_mark) )( "force" );
        returnif( inside_form and (ch_first == subscrw(1, colon)) )( "label" );
        returnif( locchar( ch_first, 1, '({[' ) )( "open" );
        returnif( locchar( ch_first, 1, ']})' ) )( "close" );
        returnif( ch_first == `,` or ch_first == `;` )( "sep" );
    endif;
    returnif( item.is_form_end )( "end" );
    returnif( item.is_sign )( "sign" );
    returnif( fast_lmember( item, inferred_form_starts ) )( "start" );
        
    returnif( fast_lmember( item, inferred_forced_words ) )( "macro" );

    lvars rest1 = subsequent_items;
    lvars next_item = not(rest1.null) and rest1.hd;

    returnunless( next_item )( "id" );
    returnif( inside_form and (next_item == colon) )( "breaker" );

    lvars peek1 =
        not(rest1.whitespace_after_item) and
        next_item == "-";
    
    returnunless( peek1 )( "id" );

    lvars rest2 = rest1.tl;
    if not( rest2.null ) and fast_lmember( rest2.hd, inferred_form_starts ) then
        "breaker1"
    else
        "id"
    endif
enddefine;


;;; Precedence rules.
;;; 10 . ( [ {
;;; 20 * / %
;;; 29 ++ --
;;; 30 + -
;;; 39 << >>
;;; 40 < <=	> >=
;;; 49 &&
;;; 50 &
;;; 59 ||
;;; 50 |
;;; 59 == =/=
;;; 60 =

constant max_precedence = 999;

define precedence( item );
    returnunless( item.isword )( false );
    returnunless( item.is_sign or item.is_open_bracket )( false );
    returnif( inside_form and (item == colon) )( false );
    lvars n = datalength( item );
    if n > 0 then
        lvars ch = subscrw( 1, item );
        lvars L = locchar( ch, 1, '.({[*/%+-<&|?:=' );
        if L then
            lvars prec = 10 * L;
            if n >= 2 and locchar(ch, 2, item ) then
                prec - 1 -> prec
            endif;
            prec
        else
            max_precedence
        endif
    else
        max_precedence
    endif
enddefine;

defclass Node {
    nodeName,
    nodeAttrs,
    nodeChildren
};

constant null_attrs = $();

vars procedure 
    read_expr, read_expr_prec, read_expr_allow_newline, newline_on_item, 
    read_opt_expr_prec;


define read_form_expr(opening_word);
    dlocal inside_form = true;
    lvars closing_keywords = [% "end", "end" <> opening_word %];
    lvars current_part = [];
    lvars current_keyword = opening_word;
    lvars procedure read = if allow_newline_option then read_expr_allow_newline else read_expr endif;
    lvars contents = [%
        lvars first_expr_in_part = true;
        lvars prev_expr_terminated = true;
        until pop11_try_nextreaditem( closing_keywords ) do
            if proglist.null then
                mishap( 'Unexpected end of file while reading form', [^opening_word])
            else
                lvars item1 = proglist.hd;
                lvars tokentype1 = classify_item( item1, proglist.tl );
                if first_expr_in_part and tokentype1 == "breaker" and unglue_option then
                    [^item1 ^unglue_option ^^(proglist.tl)] -> proglist;
                    "id" -> tokentype1;
                elseif not(first_expr_in_part) and not(prev_expr_terminated) and tokentype1 == "label" and unglue_option then
                    [^unglue_option ^^proglist] -> proglist;
                    unglue_option -> item1;
                    "breaker" -> tokentype1;
                endif;
                if tokentype1 == "sep" or tokentype1 == "sign" or tokentype1 == "close" or tokentype1 == "label" then
                    mishap( 'Unexpected item at start of expression (in ' >< opening_word >< ')', [^item1] )
                elseif tokentype1 == "end" then
                    mishap( 'Mismatched closing keyword', [^item1] )
                elseif tokentype1 == "breaker" then
                    consNode( "part", $(keyword = current_keyword), current_part );
                    [] -> current_part;
                    item1 -> current_keyword;
                    ;;; Skip the `:` or `?`.
                    lvars label_item = proglist.tl.dest -> proglist;
                    false -> first_expr_in_part;
                    true -> prev_expr_terminated;
                elseif tokentype1 == "breaker1" then
                    ;;; Skip the `-` `if`
                    lvars (t1, t2) = proglist.tl.dest.dest -> proglist;
                    lvars kw = consword( item1 >< t1 >< t2 );
                    if t2 /== opening_word then
                        mishap( 'Mismatched breaker found', [FOUND ^kw INSIDE ^opening_word])
                    endif;
                    consNode( "part", $(keyword = current_keyword), current_part );
                    [] -> current_part;
                    kw -> current_keyword;
                    true -> first_expr_in_part;
                    true -> prev_expr_terminated;
                else
                    if not( prev_expr_terminated ) then
                        mishap( 'Semi-colon or line-break expected', [^item1] )
                    endif;
                    current_part <> [% read() %] -> current_part;
                    pop11_try_nextreaditem( semi ) -> prev_expr_terminated;
                    prev_expr_terminated or (allow_newline_option and proglist.newline_on_item) -> prev_expr_terminated;
                    false -> first_expr_in_part;
                endif
            endif
        enduntil;
        consNode( "part", $(keyword = current_keyword), current_part );
    %];
    consNode( "form", null_attrs, contents)
enddefine;

define read_expr_seq_to( closing_delimiters, breakers, allow_newline );
    dlocal inside_form = false;
    lvars b = false;
	lvars items = [%
        if pop11_try_nextreaditem( closing_delimiters ) then
            ;;; Done.
        else
            repeat
                read_expr();
                quitif( pop11_try_nextreaditem( closing_delimiters ) );
                if pop11_try_nextreaditem( breakers ) ->> b then
                    [ ^b ] -> breakers;
                    if pop11_try_nextreaditem( closing_delimiters ) then
                        if b == comma then
                            mishap('Trailing comma found', [])
                        endif;
                        quitloop
                    endif;
                else
                    pop11_need_nextreaditem( closing_delimiters ) -> _;
                    quitloop
                endif;
            endrepeat;
        endif
	%];
    b, items;
enddefine;

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
        if reclassified_tokentype == "id" then
            reclassified_tokentype -> tokentype;
            unglue_option :: proglist -> proglist
        endif
    endif;
    if tokentype == "open" then
		lvars (sep, seq) = read_expr_seq_to( item.is_open_bracket, semi_comma, true );
        lvars dname = delimiter_name( item );
        consNode( "delimited", $(kind=dname, separator=sep), seq )
    elseif tokentype == "start" then
        read_form_expr( item )
    elseif tokentype == "id" then
        consNode( "identifier", $(name=item), [] )
    elseif tokentype == "force" then
        mishap( 'Misplaced macro indicator (' >< macro_mark >< ')', [] )
    elseif tokentype == "macro" then
        pop11_need_nextreaditem( macro_mark ) -> _;
        if item.isword then
            lvars e = read_opt_expr_prec(max_precedence, true);
            if e then
                consNode( "form", null_attrs, [% consNode( "part", $(keyword=item), [^e] ) %] )
            else
                consNode( "form", null_attrs, [% consNode( "part", $(keyword=item), [] ) %] )
            endif
        else
            mishap( 'Identifier required following `' >< macro_mark >< '`', [^item] )
        endif
    else
        mishap( 'Unexpected token at start of expression', [^item] )
    endif
enddefine;

define read_arguments( close_bracket );
    lvars (sep, args) = read_expr_seq_to( close_bracket, semi_comma, false);
    sep, consNode( "arguments", null_attrs, args )
enddefine;

define read_expr_prec( prec, accept_newline );
    lvars lhs = read_primary_expr();
    until null( proglist ) do
        lvars item1 = proglist.hd;
        quitif( accept_newline and newline_on_item( proglist ) );
        lvars p = precedence( item1 );
        if p and p <= prec then
            proglist.tl -> proglist;
            lvars close_bracket = false;
            if item1.is_open_bracket ->> close_bracket then
                lvars (sep, args) = read_arguments( close_bracket );
                lvars dname = delimiter_name( item1 );
                consNode( "apply", $(kind=dname, separator=sep), [^lhs ^args] ) -> lhs;
            elseif item1 == "." then
                lvars item2 = readitem();
                lvars tokentype2 = classify_item( item2, proglist );
                if tokentype2 == "id" or tokentype2 == "breaker" then
                    lvars item3 = not( proglist.null ) and proglist.hd;
                    if item3.is_open_bracket ->> close_bracket then
                        proglist.tl -> proglist;
                        lvars (sep, args) = read_arguments( close_bracket );
                        lvars dname = delimiter_name( item3 );
                        consNode( "invoke", $(kind=dname, separator=sep, name=item2), [^lhs ^args] ) -> lhs
                    else
                        consNode( "get", $(name=item2), [^lhs] ) -> lhs
                    endif
                else
                    mishap( 'Unexpected item after `.`', [^item2] )
                endif;
            else
                lvars rhs = read_expr_prec( p, false );
                consNode( "operator", $(name=item1), [^lhs ^rhs] ) -> lhs;
            endif
        else
            quitloop
        endif
    enduntil;
    return( lhs )
enddefine;

define read_opt_expr_prec( prec, accept_newline );
    returnif( proglist.null )( false );
    lvars item = proglist.hd;
    lvars tt = classify_item( item, proglist.tl );
    returnif( tt == "sep" or tt == "close" or tt == "end" or tt == "breaker" or tt == "label" or tt == "sign" )( false );
    returnif( accept_newline and newline_on_item( proglist ) )( false );
    read_expr_prec( prec, accept_newline )
enddefine;

define read_expr();
    read_expr_prec( max_precedence, false )
enddefine;

define read_expr_allow_newline();
    read_expr_prec( max_precedence, true )
enddefine;

vars procedure newline_on_item = newanyproperty(
    [], 12, 1, 8,
    false, false, "tmparg",
    false, false
);

vars procedure whitespace_after_item = newanyproperty(
    [], 12, 1, 8,
    false, false, "tmparg",
    false, false
);
vars procedure string_quote = newanyproperty(
    [], 12, 1, 8,
    false, false, "tmparg",
    false, false
);

;;; Identify forcing words by the macro_mark suffix.
define infer_forced_words( dlist );
    [%
        lvars p;
        for p on dlist do
            quitif( p.tl.null );
            lvars item = p.hd;
            if p.tl.hd == macro_mark and item.isword and not( item.is_sign ) then
                item
            endif
        endfor;
    %]
enddefine;

define infer_form_starts( dlist );
    [%
        lvars w;
        for w in dlist do
            if w.isword and is_form_end( w ) and w /== "end" then
                subword( 4, datalength( w ) - 3, w )
            endif
        endfor
    %]
enddefine;


define filter_and_annotate_proglist(itemiser) -> dlist;
    lvars dlist = pdtolist(itemiser);

    ;;; This is a sneaky hack for adding extra info to tokens - via the
    ;;; pairs of proglist! We also hack proglist around quite a bit.
    
    ;;; Here we capture the next character.
    lvars q;
    for q on dlist do
        lvars is_space = last_char_read.isinteger and locchar( last_char_read, 1, '\s\t\n\r' );
        is_space -> whitespace_after_item( q );
        if q.hd.isstring then
            consword(last_nonwhitespace_read, 1) -> string_quote( q )
        endif;
    endfor;
    
    ;;; In this loop we snip out any newlines but mark
    ;;; the subsequent pair.
    lvars p = dlist;
    until p.null or p.hd /== newline do
        p.tl ->> p -> dlist
    enduntil;
    
    until p.null or p.tl.null do
        if p.tl.hd == newline then
            p.tl.tl -> p.tl;
            true -> newline_on_item( p.tl );
        else
            p.tl -> p
        endif
    enduntil;
enddefine;

define wrap(procedure source);
    source() ->> last_char_read;
    if last_char_read.isinteger and not(locchar( last_char_read, 1, '\s\n\r\t' )) then
        last_char_read -> last_nonwhitespace_read
    endif
enddefine;

define :optargs monogram(procedure source -&- unglue="_");
    
    define lconstant monogram_parser();
        dlocal last_char_read;
        dlocal last_nonwhitespace_read;
        dlocal unglue_option = unglue;
        dlocal allow_newline_option = true; ;;; Fixing this to be true but leaving logic in in case I change course again.
        dlocal inferred_form_starts;
        dlocal inferred_forced_words;
        dlocal popnewline = true;
        dlocal inside_form = false;

        lvars procedure itemiser = incharitem(wrap(%source%));
        5 -> item_chartype( `;`, itemiser );
        7 -> item_chartype( `"`, itemiser );
        7 -> item_chartype( ```, itemiser );
        9 -> item_chartype( `#`, itemiser );

        dlocal proglist = filter_and_annotate_proglist(itemiser);
        infer_form_starts( proglist ) -> inferred_form_starts;
        infer_forced_words( proglist ) -> inferred_forced_words;
        until proglist.null do
            suspend( read_expr(), 1 )
        enduntil;
        termin
    enddefine;

    lvars proc = consproc( 0, monogram_parser );

    procedure() with_props monogram_expression_repeater;
        runproc( 0, proc )
    endprocedure
enddefine;
