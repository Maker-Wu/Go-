local GAMES = {"PLAYER", "UNKOWN'S", " ", "BATTLE", "GROUNDS"}
GAMES[0] = "I Like "
GAMES[1] = "player "
GAMES.PLAYER = nil
local tempString = ""
GAMES[" "] = "_"

for he, v in ipairs(GAMES) do
    tempString = tempString..v
end

print(tempString)