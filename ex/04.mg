[serializable]

def check():
    if a && b || c && d:
        x = x + 1
        z = 1
    else-if a.is_defined:
        return! a           #[macro return ^a]
    else:
        pass!               #[macro pass]
    endif
enddef

