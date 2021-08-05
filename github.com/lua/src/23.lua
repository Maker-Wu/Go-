local a = {}

for i = 0, 100 do
    if i%5 == 0 then
        a[i] = i
    end
end

print(#a)
print(a[0])
print(a[5])
print(a[100])