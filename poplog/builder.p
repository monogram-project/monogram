uses options

define new_monogram_builder();

    lvars stack = [];
    lvars open_calls = [];

    define lconstant get_args( prev_stack );
        lvars b = [];
        until stack == prev_stack do
            conspair( stack.destpair -> stack, b ) -> b
        enduntil;
        b.sys_grbg_destlist
    enddefine;

    define lconstant push( data );
        conspair( conspair( stack, data ), open_calls ) -> open_calls;
    enddefine;

    define lconstant pop( keyword );
        if open_calls.null then
            mishap( 'Unexpected attempt to end', [^keyword] )
        endif;
        lvars ( prev_stack, call ) = destpair( open_calls.destpair -> open_calls );
        if call.front == keyword  then
            lvars t = [%
                call.hd.lowertoupper,
                call.tl.dl,
                get_args( prev_stack )
            %];
            conspair( t, stack ) -> stack
        else
            mishap( 'Found mismatch while trying to close ' >< keyword, [% call.front %] )
        endif
    enddefine;

    define lconstant add_number( n );
        conspair( [ NUMBER ^n ], stack ) -> stack;
    enddefine;

    define lconstant add_string( q, s );
        conspair( [ STRING ^q ^s ], stack ) -> stack;
    enddefine;

    define lconstant start_delimited( dname, sep );
        push( [delimited ^dname ^sep] )
    enddefine;

    define lconstant end_delimited();
        pop( "delimited" )
    enddefine;

    define lconstant start_apply( dname, sep );
        push( [apply ^dname ^sep] )
    enddefine;

    define lconstant end_apply();
        pop( "apply" )
    enddefine;

    define lconstant add_identifier( id );
        conspair( [ IDENTIFIER ^id ], stack ) -> stack;
    enddefine;

    define lconstant start_part( part_name );
        push( [part ^part_name] )
    enddefine;

    define lconstant end_part();
        pop( "part" )
    enddefine;

    define lconstant start_form();
        push( [form] )
    enddefine;

    define lconstant end_form();
        pop( "form" )
    enddefine;

    define lconstant start_invoke( dname, sep, name );
        push( [invoke ^dname ^sep ^name] )
    enddefine;

    define lconstant end_invoke();
        pop( "invoke" )
    enddefine;

    define lconstant start_get( name );
        push( [get ^name] )
    enddefine;

    define lconstant end_get();
        pop( "get" )
    enddefine;

    define lconstant start_operator( name );
        push( [operator ^name] )
    enddefine;

    define lconstant end_operator();
        pop( "operator" )
    enddefine;

    define lconstant start_arguments();
        push( [arguments] );
    enddefine;

    define lconstant end_arguments();
        pop( "arguments" )
    enddefine;

    define lconstant new();
        stack.dl
    enddefine;

    lvars methods =
        ${
            add_number = add_number,
            start_number = add_number,
            end_number = identfn,
            add_string = add_string,
            start_string = add_string,
            end_string = identfn,
            start_delimited = start_delimited,
            end_delimited = end_delimited,
            start_apply = start_apply,
            end_apply = end_apply,
            add_identifier = add_identifier,
            start_identifier = add_identifier,
            end_identifier = identfn,
            start_form = start_form,
            end_form = end_form,
            start_part = start_part,
            end_part = end_part,
            start_invoke = start_invoke,
            end_invoke = end_invoke,
            start_get = start_get,
            end_get = end_get,
            start_operator = start_operator,
            end_operator = end_operator,
            start_arguments = start_arguments,
            end_arguments = end_arguments,
            new = new
        };

    define lconstant monogram_builder( command );
        lvars p = methods( command );
        if p then
            p()
        else
            mishap( 'Unknown action', [^command] )
        endif
    enddefine;

    return( monogram_builder )
enddefine;


