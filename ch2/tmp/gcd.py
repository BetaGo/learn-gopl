def gdc(x, y):
    while y != 0:
        x, y = y, x % y
    return x

print(gdc(16,5))