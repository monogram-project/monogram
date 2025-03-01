1 -> pop_file_versions;

extend_searchlist( 'auto', popautolist ) -> popautolist;
extend_searchlist( 'lib', popuseslist ) -> popuseslist;

load monogram.p

define :optargs print_as_xml( tree -&- indent=0 );

    define lconstant pr_html_escape( txt );
        if txt.isword then
            fast_word_string( txt ) -> txt
        elseif txt.isnumber then
            txt >< '' -> txt
        endif;
        lvars ch;
        for ch in_string txt do
            if ch == `<` then
                pr( '&lt;' )
            elseif ch == `>` then
                pr( '&gt;' )
            elseif ch == `&` then
                pr( '&amp;' )
            elseif ch == `'` then
                pr( '&apos;' )
            elseif ch == `"` then
                pr( '&quot;' )
            else
                cucharout( ch )
            endif
        endfor; 
    enddefine;

    lvars name = tree.nodeName;
    lvars attrs = tree.nodeAttrs;
    sp( indent );
    cucharout( `<` );
    pr( name );
    appnamedtuple(
        attrs,
        procedure( k, v );
            sp( 1 );
            pr( k );
            pr( '="' );
            pr_html_escape( v ); 
            cucharout( '"' )
        endprocedure
    );
    if tree.nodeChildren.null then
        npr( '/>' )
    else
        npr( '>' );
        lvars child;
        for child in tree.nodeChildren do
            print_as_xml( child -&- indent=indent+4)
        endfor;
        sp( indent );
        pr( '</' );
        pr( name );
        npr( '>' );
    endif;
enddefine;

define mg_to_xml(source);
    lvars forest = monogram(source).pdtolist.expandlist;
    if forest.length == 1 then
        forest.hd.print_as_xml
    else
        npr('unit');
        applist( forest, print_as_xml(% -&- indent=4 %) );
        npr('</unit>');
    endif
enddefine;

