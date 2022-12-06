class OrigamiParser

  Location = Struct.new(:x, :y)

  Instruction = Struct.new(:axis, :index)

  attr_reader :dots, :instructions

  def initialize(raw_data)
    @dots = []
    @instructions = []
    parse(raw_data)
  end

  def parse(input)
    ### series of lines describing dot locations as "{x},{y}"
    ###   where x increases to the right and y increases down
    ### empty line
    ### series of lines describing instructions as "fold along {axis}={index}"
    ###   where axis is x or y

    parsing_dots = true
    input.lines(chomp: true).each do |line|
      if parsing_dots
        parsing_dots = parse_dot(line)
      else
        return unless parse_instruction(line)
      end
    end
  end

  private

  def parse_dot(line)
    return false if line.empty?

    @dots << Location.new(*line.split(",").map(&:to_i))
    true
  end

  def parse_instruction(line)
    return false if line.empty?

    @instructions << line.match(/fold along (x|y)=(\d+)/) { |match| Instruction.new(match[1], match[2].to_i) }
    true
  end
end