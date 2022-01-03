# frozen_string_literal: true

require "./image_enhancer"
require "./infinite_image"

class App
  def run
    input = File.readlines("./data/input.txt", chomp: true)
    algorithm = input.shift
    input.shift
    enhancer = ImageEnhancer.new(algorithm)

    ## GOTCHA! 
    ## The Test / Example data makes 000000000 => 0
    ## The Read data has 000000000 => 1 (infinite void turns on!)
    ##               and 111111111 => 0 (infinite void turns off but will have affected "center")

    image_0 = InfiniteImage.new(input)
    puts image_0
    image_1 = enhancer.enhance(image_0)
    puts image_1
    image_2 = enhancer.enhance(image_1)
    puts image_2
    puts image_2.count_lit
  end
end

App.new.run
