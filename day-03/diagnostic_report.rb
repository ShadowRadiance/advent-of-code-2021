class DiagnosticReport
  attr_reader :data
  def initialize(data)
    @data = data
    @word_length = data&.first&.split(//)&.length || 0
  end

  def gamma_rate
    rate_map(:most_common)
  end
  
  def epsilon_rate
    rate_map(:least_common)
  end

  def oxygen_generator_rating
    filter(:most_common, 1)
  end

  def co2_scrubber_rating
    filter(:least_common, 0)
  end

  private

  def rate_map(method)
    Array.new(@word_length,0)
      .map.with_index { |_, index| send(method, index) }
      .join
      .to_i(2)
  end

  def filter(method, default)
    working_set = data.dup

    @word_length.times do |index|
      target = (send(method, index, working_set) || default).to_s
      working_set = working_set.filter { |bin| bin[index] == target }
      break if working_set.size == 1 
    end
    working_set.first&.to_i(2) || 0
  end

  def most_common(index, array = nil)
    array ||= data
    ones = count_ones(index, array)
    zeros = array.length - ones

    return nil if ones == zeros
    ones > zeros ? 1 : 0
  end

  def least_common(index, array = nil)
    mc = most_common(index, array)
    return nil if mc.nil?
    1 - mc
  end

  def count_ones(index, array)
    array.count { |bin| bin[index]=="1" }
  end
end