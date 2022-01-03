# frozen_string_literal: true

class ImageEnhancer
  attr_reader :algorithm
  def initialize(algorithm)
    # ..#.#..#####.#.#.etc
    @algorithm = algorithm

    # note:
    # if the algorithm maps 0 to # then the whole "off void" will turn on in a given step
    # if the algorithm maps 9 to . then the whole "on void" will turn off in a given step
  end

  def enhance(image)
    data = []
    ((image.min_y - 1)..(image.max_y + 1)).each do |y|
      ((image.min_x - 1)..(image.max_x + 1)).each do |x|
        nine_lights = image.nine_lights_around(x, y)
        data << @algorithm[binary_string(nine_lights).to_i(2)]
      end
    end

    new_image_background = image.background
    new_image_background = "#" if algorithm.chars.first == "#" && image.background == "."
    new_image_background = "." if algorithm.chars.last  == "." && image.background == "#"

    InfiniteImage.new(
      data.each_slice(image.width + 2).to_a.map(&:join),
      background: new_image_background
    )
  end

  private

  def binary_string(array_of_lights)
    array_of_lights.flatten.map { |light| binary(light) }.join
  end

  def binary(light)
    light == "#" ? 1 : 0
  end
end