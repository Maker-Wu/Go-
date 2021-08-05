local function MakeUpper(inT)
    for key, v in pairs(inT) do
        inT[key] = string.upper(v)
    end
end

local MyGoodTable = {nihao = "Hello", chifan = "Eat"}

string.lower(MyGoodTable.chifan)

print(MyGoodTable.chifan)
print(MyGoodTable.nihao)

MakeUpper(MyGoodTable)
print(MyGoodTable.chifan)
print(MyGoodTable.nihao)