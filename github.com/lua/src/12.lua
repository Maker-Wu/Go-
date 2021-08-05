local function func1()
    local t = {10, print, x = 12, k = "hi"}
    for k, v in pairs(t) do
        print(k, v)
    end
end

local function func2()
    local t = {10, print, 12, "hi"}
    for k, v in pairs(t) do
        print(k, v)
    end
end

local function func3()
    local t = {10, print, 12, "hi"}
    for k = 1, #t do
        print(k, t[k])
    end
end

--func1()
func3()