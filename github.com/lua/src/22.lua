local AmazingTable = { Money = 1000}

function AmazingTable:new(o)
    o = o or {}
    self.__index = self
    setmetatable(o, self)
    return o
end

function AmazingTable:show()
    print(self.Money)
end

local BillGates = AmazingTable:new()
BillGates:show()

local Trump = AmazingTable:new({Money = 100})
Trump:show()