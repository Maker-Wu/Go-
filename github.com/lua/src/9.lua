local a = {}
a[1] = 1
a[1000] = 1
a[1000] = nil
print(#a)