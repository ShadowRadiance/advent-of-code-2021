# frozen_string_literal: true

class Stack
  Underflow = Class.new(RuntimeError)
  Overflow = Class.new(RuntimeError)

  attr_reader :capacity

  def initialize(capacity: 16)
    @rep = []
    @capacity = capacity
  end

  def to_a
    @rep.dup
  end

  def empty?
    @rep.size.zero?
  end

  def full?
    @rep.size == capacity
  end

  def push(thing)
    raise Overflow if full?
    @rep << thing
  end

  def pop
    raise Underflow if empty?
    @rep.pop
  end

  def top
    raise Underflow if empty?
    @rep.last
  end

  def size
    @rep.size
  end
end