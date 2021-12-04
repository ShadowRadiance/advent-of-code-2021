class App
  def initialize
  end

  def run
  end

  def load_data_from_input
    File.readlines("./data/input.txt", chomp: true)
  end
end

App.new.run
