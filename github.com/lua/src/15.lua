a = 1
b = 2
local function f(a)
    return function(b)
        return a*b
    end
end

print(not not 1)
local b1 = f(3)
local b2 = f(4)
print(b1(2))
print(b2(3))