local a = {value = { x = 10 }}
setmetatable(a, {
    __index = function(t, k)
        return t.value[k]
    end
})

local b = a
a.value.x = 20
print(a.x)
print(a.value.x)
print(b.x)
print(b.value.x)