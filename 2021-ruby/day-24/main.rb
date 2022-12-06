# frozen_string_literal: true

require './program'
require './alu'

# find the largest valid fourteen-digit model number that contains no 0 digits
def find_largest_valid(monad)
  99_999_999_999_999.downto(11_111_111_111_111) do |model_number|
    next if model_number.to_s.include?('0')

    return model_number.to_s if monad.execute(model_number.to_s).z.zero?
  end
end

monad = Program.new(File.readlines('./data/input.txt', chomp: true), alu: ALU.new, debug: true)
puts find_largest_valid(monad)
