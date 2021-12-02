class Submarine
  attr_reader :x_pos, :depth

  def initialize
    @x_pos = 0
    @depth = 0
  end

  def follow_instructions(data)
    data.each { |instruction| follow_instruction(instruction) }
    self
  end

  def follow_instruction(inst)
    command, distance = inst.split(" ")
    distance = distance.to_i
    raise "invalid distance" if distance==0

    case command
    when 'forward'
      @x_pos += distance
    when 'down'
      @depth += distance
    when 'up'
      @depth -= distance
      @depth = 0 if @depth < 0
    else
      raise "invalid command"
    end
  end
end