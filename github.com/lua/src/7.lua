local a = {}

foo = function()
    print("foo1")
end

a.foo = foo
foo = function()
    print("foo2")
end
a.foo()
foo()