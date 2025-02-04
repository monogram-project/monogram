compile_mode :pop11 +strict;

;;;section $-options => isoptions options_key newoptions appoptions subscr_options delete_options null_options;

;;; --- Options, forward declaration -------------------------------------------

defclass options {
    options_root,
    options_default,
    options_less_than,
    options_validate_name,
    options_loop_locks
};


;;; --- Node, an internal class ------------------------------------------------

defclass node {
    node_left,
    node_name,
    node_value,
    node_right,
    node_height
};

define options_new_node( name, value, opts );
    if options_validate_name( opts )( name ) then
        consnode( false, name, value, false, 1 )
    else
        mishap( 'Unexpected name/key for options', [ ^name ] )
    endif
enddefine;

define copy_node_tree( node ) -> node;
    if node then
        node.destnode.consnode -> node;
        copy_node_tree( node.node_left ) -> node.node_left;
        copy_node_tree( node.node_right ) -> node.node_right;
    endif
enddefine;

define get_value( root, key, opts );
    returnunless( root )( opts.options_default );
    if key == root.node_name then
        root.node_value
    elseif options_less_than( opts )( key, root.node_name ) then
        get_value( root.node_left, key, opts )
    else
        get_value( root.node_right, key, opts )
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

define update_or_insert_node( root, key, value, opts );
    ;;; Find the correct location and insert the node
    returnunless( root )( options_new_node( key, value, opts ) );
    returnif( key == root.node_name )( value -> root.node_value, root );

    lvars procedure less_than = options_less_than( opts );

    if less_than( key, root.node_name ) then
        update_or_insert_node( root.node_left, key, value, opts ) -> root.node_left
    else
        update_or_insert_node( root.node_right, key, value, opts ) -> root.node_right
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
define delete_node(root, key, opts);
    returnunless( root )( root );
    if key == root.node_name then
        returnunless( root.node_left )( root.node_right );
        returnunless( root.node_right )( root.node_left );
        lvars temp = getMinValueNode( root.node_right );
        temp.node_name -> root.node_name;
        temp.node_value -> root.node_value;
        delete_node(root.node_right, temp.node_name, opts ) -> root.node_right;
    elseif options_less_than( opts )( key, root.node_name ) then
        delete_node( root.node_left, key, opts ) -> root.node_left
    else
        delete_node( root.node_right, key, opts ) -> root.node_right
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


;;; --- Options ----------------------------------------------------------------

define :optargs newoptions(-&- def=false, less_than=alphabefore, validate_name=isword );
    consoptions( false, def, less_than, validate_name, false )
enddefine;

define lconstant copy_when_locked( opts );
    copy_node_tree( opts.options_root ) -> opts.options_root;
    false -> opts.options_loop_locks;
enddefine;

define subscr_options( k, opts );
    get_value( opts.options_root, k, opts )
enddefine;

define updaterof subscr_options( v, k, opts );
    if opts.options_loop_locks then
        copy_when_locked( opts );
    endif;
    if v == opts.options_default then
        delete_node( opts.options_root, k, opts ) -> opts.options_root
    else
        update_or_insert_node( opts.options_root, k, v, opts ) -> opts.options_root
    endif
enddefine;

subscr_options -> class_apply( options_key );

define delete_options( k, opts );
    if opts.options_loop_locks then
        copy_when_locked( opts );
    endif;
    delete_node( opts.options_root, k, opts ) -> opts.options_root
enddefine;

define appoptions( opts, procedure p );
    lvars lockpair = conspair(false, opts.options_loop_locks);
    lockpair -> opts.options_loop_locks;
    appnode( opts.options_root, p );
    if opts.options_loop_locks == lockpair then
        lockpair.back -> opts.options_loop_locks
    else
        lvars spine = opts.options_loop_locks;
        while spine do
            if spine.back == lockpair then
                sys_grbg_destpair( lockpair ) -> spine.back -> _;
                return
            endif;
            spine.back -> spine
        endwhile;
    endif
enddefine;

define null_options( opts );
    opts.options_root == false
enddefine;

define copy_options( opts );
    copy_node_tree( opts.options_root ) -> opts.options_root
enddefine;

define print_options( opts );
    define lconstant prnode( n );
        if n then
            pr('(');
            prnode(n.node_left);
            pr('<');
            pr(n.node_name);
            pr('=');
            pr(n.node_value);
            pr('>');
            prnode(n.node_right);
            pr(')');
        else
            pr('.')
        endif
    enddefine;
    prnode(opts.options_root);
    pr(newline);
enddefine;

;;;endsection;
