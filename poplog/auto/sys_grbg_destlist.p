compile_mode :pop11 +strict;

section;

define sys_grbg_destlist( L );
    until null( L ) do
        sys_grbg_destpair( L ) -> L
    enduntil
enddefine;

endsection;
