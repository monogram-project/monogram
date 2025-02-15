1 -> pop_file_versions;

extend_searchlist( 'auto', popautolist ) -> popautolist;
extend_searchlist( 'lib', popuseslist ) -> popuseslist;

load monogram.p

define show(source);
    applist( monogram(source).pdtolist, pretty)
enddefine;

/*
lvars f;
for f in ['01' '02' '03' '04' '05' '06'] do
    npr( f );
    show(discin('../ex/' >< f >< '.mg'));
endfor;
*/
