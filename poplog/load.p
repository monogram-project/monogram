1 -> pop_file_versions;

extend_searchlist( 'auto', popautolist ) -> popautolist;
extend_searchlist( 'lib', popuseslist ) -> popuseslist;

load builder.p
load monogram.p

define convert( tree, builder );
    builder( tree.nodeAttrs.dl, "start_" <> tree.nodeName );
    lvars arg;
    for arg in tree.nodeChildren do
        convert( arg, builder )
    endfor;
    builder( "end_" <> tree.nodeName );
enddefine;

define show(source);
    lvars tree;
    for tree in monogram(source).pdtolist do
        { mg ^tree }==>
        lvars builder = new_monogram_builder();
        convert( tree, builder );
        builder( "new" ) ==>
    endfor
enddefine;

/*
lvars f;
for f in ['01' '02' '03' '04' '05' '06'] do
    npr( f );
    show(discin('../ex/' >< f >< '.mg'));
endfor;
*/
