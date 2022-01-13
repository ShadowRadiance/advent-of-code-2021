# frozen_string_literal: true

# Simple Arithmetic Logic Unit
class ALU
  def initialize(input_stream: $stdin)
    @registers = {
      'w' => 0,
      'x' => 0,
      'y' => 0,
      'z' => 0
    }
    @input_stream = input_stream
  end

  OPERATIONS = %w[inp add mul div mod eql].freeze

  %i[w x y z].each do |register_name|
    define_method register_name do
      @registers[register_name.to_s]
    end
  end

  def execute(string)
    perform(*string.split(' '))
  end

  def perform(operation, register, register_or_value = nil)
    assert_operation(operation)
    assert_register(register)

    if register_or_value.nil?
      send(operation, register)
    else
      assert_register_or_value(register_or_value)

      send(operation, register, value_of(register_or_value))
    end
  end

  private

  def inp(register)
    assert_stream_not_eof
    registers[register] = Integer(input_stream.getc)
  end

  def add(register, value)
    registers[register] += value
  end

  def mul(register, value)
    registers[register] *= value
  end

  def div(register, value)
    assert_not_zero(value)
    registers[register] /= value
  end

  def mod(register, value)
    assert_not_negative(register)
    assert_positive(value)
    registers[register] -= div(register, value) * value
  end

  def eql(register, value)
    registers[register] = registers[register] == value ? 1 : 0
  end

  def value_of(value_or_register)
    if registers.key?(value_or_register)
      registers[value_or_register]
    else
      value_or_register.to_i
    end
  end

  def assert_operation(operation)
    raise ArgumentError, "#{operation} is not a valid operation" unless OPERATIONS.include?(operation)
  end

  def assert_register(register)
    raise ArgumentError, "#{register} is not a valid register" unless registers.key?(register)
  end

  def assert_register_or_value(register_or_value)
    return if registers.key?(register_or_value)

    return if register_or_value.to_i.to_s == register_or_value.to_s

    raise ArgumentError, "#{register_or_value} is neither a valid register nor an integer"
  end

  def assert_not_zero(value)
    raise ArgumentError, "#{value} must not be zero" if value.zero?
  end

  def assert_not_negative(register)
    raise ArgumentError, "#{register} (#{value_of(register)}) must not be negative" if value_of(register).negative?
  end

  def assert_positive(value)
    raise ArgumentError, "#{value} must be positive" unless value.positive?
  end

  def assert_stream_not_eof
    raise EOFError, 'No more data in input stream' if input_stream.eof?
  end

  attr_reader :input_stream, :registers
end
