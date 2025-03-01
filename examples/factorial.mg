# The factorial function in monogram.
def f(n):
    if n <= 1:
        1
    else:
        n * f(n - 1)
    endif
enddef