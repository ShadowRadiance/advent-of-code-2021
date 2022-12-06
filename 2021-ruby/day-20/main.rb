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

    image = InfiniteImage.new(input)
    50.times {
      puts image.count_lit
      image = enhancer.enhance(image) 
    }
    puts image
    puts image.count_lit
  end
end

App.new.run
