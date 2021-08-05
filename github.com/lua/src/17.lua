mytable = setmetatable({key1 = "PENTA KILL"}, {__index = {key2 = "ABCD"}})

mytable.key1 = nil
print(mytable.key1, mytable.key2)

print("16a" < "15a")
print("16" < "17")
print("17" < "17")