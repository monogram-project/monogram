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
;;;     8. Keywords - foo: foo?                 false               "keyword"
;;;     9. Label - : ?                          false               "label"
;;;    10. Signs - + / %                        true                "sign"
;;;    11. Identifiers                          false               "id"

vars unglue_option = false;
vars allow_newline_option = false;
vars inferred_form_starts = [];
vars allow_trailing_comma = false;

define peek_item();
    not( null( proglist ) ) and proglist.hd
enddefine;

define peek_nth_item( n );
    lvars PL = proglist;
    while n > 1 do
        returnif( null( PL ) )( false );
        PL.tl -> PL;
        n - 1 -> n;
    endwhile;
    not( null( PL ) ) and PL.hd
enddefine;

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
    return( word /== "!" and word /== ":" and word /== "?" );
enddefine;

define is_form_end( word );
    isstartstring( 'end', word )
enddefine;

define classify_item( item, next_item );
    returnunless( item.isword )( false );
    lvars L = datalength( item );
    returnif( L == 0 )( false );
    if L == 1 then
        lvars ch_first = subscrw( 1, item );
        returnif( ch_first == `!` )( "force" );
        returnif( ch_first == `:` or ch_first == `?` )( "label" );
        returnif( locchar( ch_first, 1, '({[' ) )( "open" );
        returnif( locchar( ch_first, 1, ']})' ) )( "close" );
        returnif( ch_first == `,` or ch_first == `;` )( "sep" );
    endif;
    returnif( item.is_form_end )( "end" );
    returnif( item.is_sign )( "sign" );
    if fast_lmember( item, inferred_form_starts ) then
        "start"
    elseif next_item == ":" or next_item == "?" then
        "keyword"
    else
        "id"
    endif
enddefine;

define is_id_form_opening( next_item, item_after );
    lvars tokentype = classify_item( next_item, item_after );
    not( tokentype ) or tokentype == "id"
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
    returnif( item == ":" or item == "?" )( false );
    lvars n = datalength( item );
    if n > 0 then
        lvars ch = subscrw( 1, item );
        lvars L = locchar( ch, 1, '.?({[*/%+-<&|=' );
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


vars procedure read_expr, read_expr_prec, read_expr_allow_newline, newline_on_item;

constant semi_comma = [ ; , ];

define read_form_expr(opening_word);
    lvars closing_keywords = [% "end", "end" <> opening_word %];
    lvars current_part = [];
    lvars current_keyword = opening_word;
    lvars procedure read = if allow_newline_option then read_expr_allow_newline else read_expr endif;
    lvars contents = [%
        lvars first_expr = true;
        until pop11_try_nextreaditem( closing_keywords ) do
            if proglist.null then
                mishap( 'Unexpected end of file while reading form', [^opening_word])
            else
                lvars item1 = proglist.hd;
                lvars tokentype1 = classify_item( item1, peek_nth_item(2) );
                if tokentype1 == "label" and unglue_option then
                    unglue_option :: proglist -> proglist;
                    unglue_option -> item1;
                    "keyword" -> tokentype1;
                endif;
                if tokentype1 == "sep" or tokentype1 == "sign" or tokentype1 == "close" or tokentype1 == "label" then
                    mishap( 'Unexpected item at start of expression (in ' >< opening_word >< ')', [^item1] )
                elseif tokentype1 == "end" then
                    mishap( 'Mismatched closing keyword', [^item1] )
                elseif tokentype1 == "keyword" then
                    [part ^current_keyword ^^current_part];
                    [] -> current_part;
                    item1 -> current_keyword;
                    ;;; Skip the `:` or `?`.
                    proglist.tl.tl -> proglist;
                    true -> first_expr;
                else
                    if not( first_expr ) then
                        lvars msg = if allow_newline_option then 'Semi-colon or line-break expected' else 'Semi-colon expected' endif;
                        mishap( msg, [^item1] )
                    endif;
                    current_part <> [% read() %] -> current_part;
                    pop11_try_nextreaditem( ";" ) -> first_expr;
                    first_expr or (allow_newline_option and proglist.newline_on_item) -> first_expr;
                endif
            endif
        enduntil;
        [part ^current_keyword ^^current_part];
    %];
    [form ^^contents]
enddefine;

define read_expr_seq_to( closing_delimiters, breakers, allow_newline );
	lvars items = [%
        if pop11_try_nextreaditem( closing_delimiters ) then
            ;;; Done.
        else
            repeat
                read_expr();
                quitif( pop11_try_nextreaditem( closing_delimiters ) );
                lvars b;
                if pop11_try_nextreaditem( breakers ) ->> b then
                    if pop11_try_nextreaditem( closing_delimiters ) then
                        if b == "," then
                            if allow_trailing_comma then
                                [trailing_comma]
                            else
                                mishap('Trailing comma found', [])
                            endif
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
    items;
enddefine;

define read_primary_expr();
    lvars item = readitem();
    lvars tokentype = classify_item( item, peek_item() );
    returnunless( tokentype )( [constant ^item] );
    if tokentype == "keyword" and unglue_option then
        lvars reclassified_tokentype = classify_item( item, unglue_option );
        if reclassified_tokentype == "id" then
            reclassified_tokentype -> tokentype;
            unglue_option :: proglist -> proglist
        endif
    endif;
    if tokentype == "open" then
		lvars seq = read_expr_seq_to( item.is_open_bracket, semi_comma, true );
        lvars dname = delimiter_name( item );
        [delimited ^dname ^^seq]
    elseif tokentype == "start" then
        read_form_expr( item )
    elseif tokentype == "id" then
        ;;; The interpretation depends on what comes next.
        if null(proglist) then
            [identifier ^item]
        else
            lvars item1 = proglist.hd;
            if is_id_form_opening( item1, peek_nth_item(2) ) then
                read_form_expr( item )
            else
                [identifier ^item]
            endif
        endif
    elseif tokentype == "force" then
        lvars item1 = readitem();
        if item1.isword then
            [form [part ^item1]]
        else
            mishap( 'Identifier required following `!`', [^item] )
        endif
    else
        mishap( 'Unexpected token at start of expression', [^item] )
    endif
enddefine;

define read_arguments( close_bracket );
    [arguments ^^(read_expr_seq_to( close_bracket, semi_comma, false))]
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
                lvars args = read_arguments( close_bracket );
                lvars dname = delimiter_name( item1 );
                [apply ^dname ^lhs ^args] -> lhs;
            elseif item1 == "." then
                lvars item2 = readitem();
                lvars tokentype2 = classify_item( item2, peek_item() );
                if tokentype2 == "id" then
                    lvars item3 = not( proglist.null ) and proglist.hd;
                    if item3.is_open_bracket ->> close_bracket then
                        proglist.tl -> proglist;
                        lvars args = read_arguments( close_bracket );
                        lvars dname = delimiter_name( item3 );
                        [invoke ^dname ^item2 ^lhs ^args] -> lhs
                    else
                        [get ^item2 ^lhs] -> lhs
                    endif
                else
                    mishap( 'Unexpected item after `.`', [^item2] )
                endif;
            else
                lvars rhs = read_expr_prec( p, false );
                [operator ^item1 ^lhs ^rhs] -> lhs;
            endif
        else
            quitloop
        endif
    enduntil;
    return( lhs )
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

define filter_and_annotate_proglist();
    ;;; This is a sneaky hack for adding extra info to tokens - via the
    ;;; pairs of proglist! In this loop we snip out any newlines but mark
    ;;; the subsequent pair.
    lvars p = proglist;
    until p.null or p.hd /== newline do
        p.tl -> proglist
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

define :optargs monogram(procedure source -&- unglue=false, opt_seps=false, opt_trailing/allow_trailing_comma=false);
    dlocal unglue_option = unglue;
    dlocal allow_newline_option = opt_seps;
    dlocal inferred_form_starts;
    dlocal popnewline = true;
    dlocal allow_trailing_comma = opt_trailing;

    lvars procedure itemiser = incharitem(source);
    5 -> item_chartype( `;`, itemiser );
    9 -> item_chartype( `#`, itemiser );

    dlocal proglist = pdtolist(itemiser);
    filter_and_annotate_proglist();

    infer_form_starts( proglist ) -> inferred_form_starts;

    read_expr()
enddefine;
