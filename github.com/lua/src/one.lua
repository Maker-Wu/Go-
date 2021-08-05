local function newCounter()
    local count = 0
    return function()
        count = count + 1
        return count
    end
end

LOL = newCounter()
print(LOL())
print(LOL())