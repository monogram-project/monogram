compile_mode :pop11 +strict;

uses pop11_comp_N;
uses pop11_named_arg_mark;

;;;section $-options => isoptions options_key newoptions appoptions subscr_options delete_options null_options;

;;; --- Optconfig, forward declaration -------------------------------------------

;;; Static info that is read-only.
defclass optconfig {
    optconfig_default,        ;;; default false
    optconfig_less_than,      ;;; default allphabefore
    optconfig_validate_name,  ;;; default isword
};


;;; --- Node, an internal class ------------------------------------------------

defclass node {
    node_left,
    node_name,
    node_value,
    node_right,
    node_height
};

define optdata_new_node( name, value, optconfig );
    if optconfig_validate_name( optconfig )( name ) then
        consnode( false, name, value, false, 1 )
    else
        mishap( 'Unexpected name/key', [ ^name ] )
    endif
enddefine;

define copy_node_tree( node ) -> node;
    if node then
        node.destnode.consnode -> node;
        copy_node_tree( node.node_left ) -> node.node_left;
        copy_node_tree( node.node_right ) -> node.node_right;
    endif
enddefine;

define get_value( root, key, optconfig );
    returnunless( root )( optconfig.optconfig_default );
    if key == root.node_name then
        root.node_value
    elseif optconfig_less_than( optconfig )( key, root.node_name ) then
        get_value( root.node_left, key, optconfig )
    else
        get_value( root.node_right, key, optconfig )
    endif
enddefine;

define getHeight( t );
    returnunless( t )( 0 );
    t.node_height
enddefine;

define getBalance( n );
    returnunless( n )( 0 );
    getHeight( n.node_left ) - getHeight( n.node_right )
enddefine;

define reviseHeight( n );
    returnunless( n );
    1 + max( getHeight( n.node_left ), getHeight( n.node_right ) ) -> n.node_height
enddefine;

define getMinValueNode( n );
    lvars root = n;
    while root and root.node_left do
        root.node_left -> root
    endwhile;
    root
enddefine;

define appnode(self, procedure p);
    returnunless( self );
    if self.node_left then
        appnode( self.node_left, p )
    endif;
    p( self.node_name, self.node_value );
    if self.node_right then
        appnode( self.node_right, p )
    endif
enddefine;

;;; Function to perform left rotation.
define leftRotate( self );
    lvars y = self.node_right;
    lvars T2 = y.node_left;
    self -> y.node_left;
    T2 -> self.node_right;
    reviseHeight( self );
    reviseHeight( y );
    y
enddefine;

;;; Function to perform right rotation.
define rightRotate( self );
    lvars y = self.node_left;
    lvars T3 = y.node_right;
    self -> y.node_right;
    T3 -> self.node_left;
    reviseHeight( self );
    reviseHeight( y );
    y
enddefine;

define update_or_insert_node( root, key, value, optconfig );
    ;;; Find the correct location and insert the node
    returnunless( root )( optdata_new_node( key, value, optconfig ) );
    returnif( key == root.node_name )( value -> root.node_value, root );

    lvars procedure less_than = optconfig_less_than( optconfig );

    if less_than( key, root.node_name ) then
        update_or_insert_node( root.node_left, key, value, optconfig ) -> root.node_left
    else
        update_or_insert_node( root.node_right, key, value, optconfig ) -> root.node_right
    endif;

    ;;; Update the balance factor.
    reviseHeight(root);
    lvars balanceFactor = getBalance( root );

    ;;; Rebalance the tree.
    if balanceFactor > 1 then
        ;;; The left side of the tree is heavier - and must be truthy.
        if less_than( key, root.node_left.node_name ) then
            rightRotate( root )
        else
            leftRotate( root.node_left ) -> root.node_left;
            rightRotate( root )
        endif
    elseif balanceFactor < -1 then
        ;;; The right side of the tree is heavier - and must be truthy.
        if less_than( root.node_right.node_name, key ) then
            leftRotate( root )
        else
            rightRotate( root.node_right ) -> root.node_right;
            leftRotate( root )
        endif
    else
        root
    endif
enddefine;

;;; Function to delete a node
define delete_node( root, key, optconfig );
    returnunless( root )( root );
    if key == root.node_name then
        returnunless( root.node_left )( root.node_right );
        returnunless( root.node_right )( root.node_left );
        lvars temp = getMinValueNode( root.node_right );
        temp.node_name -> root.node_name;
        temp.node_value -> root.node_value;
        delete_node(root.node_right, temp.node_name, optconfig ) -> root.node_right;
    elseif optconfig_less_than( optconfig )( key, root.node_name ) then
        delete_node( root.node_left, key, optconfig ) -> root.node_left
    else
        delete_node( root.node_right, key, optconfig ) -> root.node_right
    endif;
    
    ;;; Update the balance factor of nodes
    reviseHeight( root );

    lvars balanceFactor = getBalance( root );

    ;;; Balance the tree
    if balanceFactor > 1 then
        if getBalance( root.node_left ) >= 0 then
            rightRotate( root )
        else
            leftRotate( root.node_left ) -> root.node_left;
            rightRotate( root )
        endif
    elseif balanceFactor < -1 then
        if getBalance( root.node_right ) <= 0 then
            leftRotate( root )
        else
            rightRotate( root.node_right ) -> root.node_right;
            leftRotate( root )
        endif
    else
        root
    endif;
enddefine;

define length_node( root );
    returnunless( root )( 0 );
    lvars L = length_node( root.node_left );
    lvars count = L fi_+ 1;
    while root.node_right ->> root do
        count fi_+ 1 -> count
    endwhile;
    count
enddefine;



;;; --- Optdata (Generic implementation) ---------------------------------------

;;; The configuration that is suitable for -options- specifically, which
;;; are defined to work exclusively on words and with a false default.
constant OPTION_OPTCONFIG = consoptconfig( false, alphabefore, isword );

defclass optdata {
    optdata_root,           ;;; pointer to a balanced binary tree implementation.
    optdata_loop_locks,     ;;; copy-on-write locks, allowing concurrent iteration and update.
    optdata_config          ;;; static, shared info.
};

define newoptdata( optconfig );
    consoptdata( false, false, optconfig )
enddefine;

define newoptdata_for_options();
    newoptdata( OPTION_OPTCONFIG )
enddefine;

define extend_optdata( marker, optdata ) -> optdata;
    if marker /== pop11_named_arg_mark then
        mishap( 'Missing optional args', [ ^optdata ] )
    endif;
    lvars N = fi_check( (), false, false );
    lvars optconfig = optdata.optdata_config;
    fast_repeat N times
        lvars ( value, name ) = ();
        if value == optconfig.optconfig_default then
            delete_node( optdata.optdata_root, name, optconfig ) -> optdata.optdata_root;
        else
            update_or_insert_node( optdata.optdata_root, name, value, optconfig ) -> optdata.optdata_root
        endif
    endrepeat;
enddefine;

define lconstant copy_when_locked( optdata );
    copy_node_tree( optdata.optdata_root ) -> optdata.optdata_root;
    false -> optdata.optdata_loop_locks;
enddefine;

define subscr_optdata( k, optdata );
    get_value( optdata.optdata_root, k, optdata.optdata_config )
enddefine;

define updaterof subscr_optdata( v, k, optdata );
    if optdata.optdata_loop_locks then
        copy_when_locked( optdata );
    endif;
    lvars optconfig = optdata.optdata_config;
    if v == optconfig.optconfig_default then
        delete_node( optdata.optdata_root, k, optconfig ) -> optdata.optdata_root;
    else
        update_or_insert_node( optdata.optdata_root, k, v, optconfig ) -> optdata.optdata_root
    endif
enddefine;

subscr_optdata -> class_apply( optdata_key );

define delete_optdata( k, optdata );
    if optdata.optdata_loop_locks then
        copy_when_locked( optdata );
    endif;
    delete_node( optdata.optdata_root, k, optdata.optdata_config ) -> optdata.optdata_root
enddefine;

define appoptdata( optdata, procedure p );
    ;;; Establish update detection for thiis instance of iteration.
    lvars lockpair = conspair(false, optdata.optdata_loop_locks);
    lockpair -> optdata.optdata_loop_locks;

    appnode( optdata.optdata_root, p );

    ;;; Remove update detection for this instance of iteration.
    if optdata.optdata_loop_locks == lockpair then
        ;;; Will almost always be this arm.
        lockpair.back -> optdata.optdata_loop_locks
    else
        ;;; When we have situations such as early exit from an inner iteration
        ;;; we will need this arm. It might not find any instance of the 
        ;;; lock-pair - which is fine. That would just mean there was an 
        ;;; intervening write that caused a copy-on-write event.
        lvars spine = optdata.optdata_loop_locks;
        while spine do
            if spine.back == lockpair then
                sys_grbg_destpair( lockpair ) -> spine.back -> _;
                return
            endif;
            spine.back -> spine
        endwhile;
    endif
enddefine;

define is_null_optdata( optdata );
    optdata.optdata_root == false
enddefine;

define length_optdata( optdata );
    length_node( optdata.optdata_root )
enddefine;

define copy_optdata( optdata );
    lvars t = copy_node_tree( optdata.optdata_root );
    
    ;;; Create shallow copy.
    optdata.destoptdata.consoptdata -> optdata;

    ;;; Overwrite the mutable fields.
    t -> optdata.optdata_root;
    false -> optdata.optdata_loop_locks;

    optdata
enddefine;


;;; --- Options ----------------------------------------------------------------
;;; This is a thin outer wrapper around optdata that allows us to encapsulate
;;; the fact that although optdata can work with any datatype T where
;;; equality on T is == and there is an ordering T.less_than.
;;; our options datatype is deliberately limited to words.


defclass options {
    options_data
};

define newoptions();
    lvars optdata = newoptdata_for_options();
    consoptions(
        if stacklength() fi_> 0 and dup() == pop11_named_arg_mark then
            extend_optdata( optdata )
        else
            optdata
        endif
    )
enddefine;

define subscr_options( k, opts );
    subscr_optdata( k, opts.options_data )
enddefine;

define updaterof subscr_options( v, k, opts );
    v -> subscr_optdata( k, opts.options_data )
enddefine;

subscr_options -> class_apply( options_key );

define delete_options( k, opts );
    lvars optdata = opts.options_data;
    delete_optdata( k, optdata )
enddefine;

define appoptions( opts, procedure p );
    appoptdata( opts.options_data, p )
enddefine;

define is_null_options( opts );
    opts.options_data.is_null_optdata
enddefine;

define copy_options( opts );
    lvars optdata = opts.options_data;
    consoptions( copy_optdata( optdata ) )
enddefine;

define print_options( opts );
    dlvars sep = '';
    pr( '${' );
    appoptions(
        opts,
        procedure( name, value );
            pr( sep );
            pr( name );
            pr( '=' );
            pr( value );
            ', ' -> sep;
        endprocedure
    );
    pr( '}' )
enddefine;

print_options -> class_print( options_key );


;;; --- Syntax -----------------------------------------------------------------

define lconstant procedure check_duplicates( key_index_list );
    lvars tail;
    for tail on key_index_list do
        if fast_back( tail ).ispair then
            if fast_front( fast_front( tail ) ) == fast_front( fast_front( fast_back( tail ) ) ) then
                mishap( 'Trying to construct dict with non-unique key', [% front(front(tail)) %] )
            endif
        endif
    endfor;
enddefine;

define compile_newoptions_to( closing_keyword ) -> actual_closer;
    dlocal pop_new_lvar_list;
    lvars keys = [];
    lvars n = 0;
    lvars tmpvars = {%
        until pop11_try_nextreaditem( closing_keyword ) ->> actual_closer do
            while pop11_try_nextreaditem( "," ) do endwhile;
            n + 1 -> n;
            lvars k = readitem();                    ;;; TODO: must be a word
            unless k.isword do
                mishap( 'Expected word as dict key', [^k] )
            endunless;
            conspair( k, n ) :: keys -> keys;
            pop11_need_nextreaditem( "=" ) -> _;
            dlvars tmpvar = sysNEW_LVAR();
            pop11_comp_N( procedure(); pop11_comp_expr(); sysPOP(tmpvar) endprocedure, 0 );
            tmpvar
        enduntil;
    %};
    nc_listsort(
        keys,
        procedure( x, y ); alphabefore( x.front, y.front ) endprocedure
    ) -> keys;
    check_duplicates( keys );
    lvars name_index_pair;
    for name_index_pair in keys do
        sysPUSH( subscrv( name_index_pair.back, tmpvars ) );
        sysPUSHQ( name_index_pair.front );
    endfor;
    sysPUSHQ( tmpvars.datalength );
    sysPUSHQ( pop11_named_arg_mark );
    sysCALL( "newoptions" );
enddefine;

;;;endsection;
