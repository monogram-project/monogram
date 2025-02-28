load mg_to_xml.p

sysgarbage();
sysgarbage();

define save();
    if syssave( 'monogram.psv' ) then
        lvars file;
        for file in poparglist do
            mg_to_xml(discin(file));
        endfor;
    endif;
enddefine;

