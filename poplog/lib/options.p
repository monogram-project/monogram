
compile_mode :pop11 +strict;

;;;section $-options => isoptions options_key newoptions appoptions subscr_options delete_options null_options;


defclass node {
    node_left,
    node_name,
    node_value,
    node_right,
    node_height
};

define newnode( key, value );
    consnode( false, key, value, false, 1 )
enddefine;

define get_value( root, key, default );
    returnunless( root )( default );
    if key == root.node_name then
        root.node_value
    elseif alphabefore( key, root.node_name ) then
        get_value( root.node_left, key, default )
    else
        get_value( root.node_right, key, default )
    endif
enddefine;

define getHeight( t );
    returnunless( t )( 0 );
    t.node_height
enddefine;

define getBalance( self );
    returnunless( self )( 0 );
    getHeight( self.node_left ) - getHeight( self.node_right )
enddefine;

define reviseHeight( self );
    returnunless( self );
    1 + max( getHeight( self.node_left ), getHeight( self.node_right ) ) -> self.node_height
enddefine;

define getMinValueNode( self );
    lvars root = self;
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

define update_or_insert_node( root, key, value );
    ;;; Find the correct location and insert the node
    returnunless( root )( newnode( key, value) );
    returnif( key == root.key )( value -> root.node_value, root );

    if alphabefore( key, root.node_name ) then
        update_or_insert_node(root.node_left, key, value) -> root.node_left
    else
        update_or_insert_node(root.node_right, key, value) -> root.node_right
    endif;

    ;;; Update the balance factor.
    reviseHeight(root);
    lvars balanceFactor = getBalance( root );

    ;;; Rebalance the tree.
    if balanceFactor > 1 then
        ;;; The left side of the tree is heavier - and must be truthy.
        if key < root.node_left.node_name then
            rightRotate( root )
        else
            leftRotate( root.node_left ) -> root.node_left;
            rightRotate( root )
        endif
    elseif balanceFactor < -1 then
        ;;; The right side of the tree is heavier - and must be truthy.
        if key > root.node_right.node_name then
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
define delete_node(root, key);
    ;;; Find the node to be deleted and remove it
    returnunless( root )( root );
    if key == root.node_name then
        returnunless( root.node_left )( root.node_right );
        returnunless( root.node_right )( root.node_left );
        lvars temp = getMinValueNode( root.node_right );
        temp.node_name -> root.node_name;
        delete_node(root.node_right, temp.node_name) -> root.node_right;
    elseif key < root.node_name then
        delete_node( root.node_left, key ) -> root.node_left
    else
        delete_node( root.node_right, key ) -> root.node_right
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
            root.node_right = rightRotate( root.node_right );
            leftRotate( root )
        endif
    else
        root
    endif
enddefine;

defclass options {
    options_root,
    options_default
};

define newoptions( def );
    consoptions( false, def )
enddefine;

define subscr_options( k, opts );
    get_value( opts.options_root, k, opts.options_default )
enddefine;

define updaterof subscr_options( v, k, opts );
    if k == opts.options_default then
        delete_node( opts.options_root, k ) -> opts.options_root
    else
        update_or_insert_node( opts.options_root, k, v ) -> opts.options_root
    endif
enddefine;

subscr_options -> class_apply( options_key );

define delete_options( k, opts );
    delete_node( opts.options_root, k ) -> opts.options_root
enddefine;

define appoptions( opts, procedure p );
    appnode( opts.options_root, p )
enddefine;

define null_options( opts );
    opts.options_root /== false
enddefine;


;;;endsection;
