class DiagnosticReport
  attr_reader :data
  def initialize(data)
    @data = data
  end

  def gamma_rate
    return 0 if counts.empty?
    counts.map { |count| count.common.to_s }.join.to_i(2)
  end
  
  def epsilon_rate
    return 0 if counts.empty?
    counts.map { |count| count.uncommon.to_s }.join.to_i(2)
  end

  private

  Count = Struct.new(:zeros, :ones, keyword_init: true) do
    def common
      zeros > ones ? 0 : 1
    end
    def uncommon
      zeros < ones ? 0 : 1
    end
  end

  def counts
    return [] if data.empty?

    @counts ||= determine_counts
  end

  def determine_counts
    # data
    # ["010", "111", ...]

    list_of_digit_arrays = data.map { |row| row.split(//).map(&:to_i) }
    # [ [0,1,0], [1,1,1], ... ]
    
    numbers_of_ones = list_of_digit_arrays.reduce { |sums, digits| sums.zip(digits).map(&:sum) }
    # [ 15, 26, 12, ... ]
    
    total_rows = data.length
    # 50

    numbers_of_zeros = numbers_of_ones.map { |ones| total_rows - ones }
    # [ 35, 24, 38, ... ]

    numbers_of_zeros.zip(numbers_of_ones).map { |zeros, ones| Count.new(zeros: zeros, ones: ones) }
    # [ Count(zeros: 35, ones: 15), Count(zeros: 24, ones: 26), ... ]
  end
end