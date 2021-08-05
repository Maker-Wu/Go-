local function f(...)
    print(...)
end

f "Hello World"
f {x = 10, y = 20}
f(type({}))
f 'a.lua'