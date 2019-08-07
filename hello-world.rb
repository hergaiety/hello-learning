# https://stackoverflow.com/questions/705729/how-do-i-create-a-ruby-hello-world#705754

class HelloWorld
  def initialize(name)
    @name = name.capitalize
  end
  def sayHi
    puts "Hello #{@name}!"
  end
end

hello = HelloWorld.new("World")
hello.sayHi
