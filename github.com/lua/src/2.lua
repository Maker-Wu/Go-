loadfile("a.lua")

print(tonumber("10e"))
print(tonumber(" -3 "))
print(tonumber("Z", 36))
print(tonumber("987", 8))

print((4 and 5)==5)
print((nil and 5) == nil)
print((nil or false) == nil)
print((0 or 5) == 0)