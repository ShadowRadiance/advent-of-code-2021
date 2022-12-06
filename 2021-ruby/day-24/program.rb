# frozen_string_literal: true

require 'stringio'

# Program
class Program
  def initialize(instructions, alu:, debug: false)
    @instructions = instructions
    @alu = alu
    @debug = debug
  end

  def execute(input_string)
    alu.input_stream = StringIO.new(input_string)
    puts input_string if debug
    instructions.each do |instruction|
      alu.execute(instruction)
    end
    alu
  end

  private

  attr_reader :instructions, :alu, :debug
end
